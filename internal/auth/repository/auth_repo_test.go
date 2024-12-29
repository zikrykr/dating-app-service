package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dating-app-service/internal/auth/model"
	"github.com/dating-app-service/internal/auth/port"
	"github.com/dating-app-service/pkg"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository port.IAuthRepo
}

func (s *Suite) SetupSuite() {
	dbMock, err := pkg.ConnectDB()
	require.NoError(s.T(), err)
	s.mock = dbMock.SQLMock
	s.repository = NewRepository(dbMock.GormDB)
}

func (s *Suite) AfterTest(_, _ string) {
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func Test_runner(t *testing.T) {
	suite.Run(t, &Suite{})
}

func (s *Suite) Test_repository_GetUserByEmail() {

	var (
		query = `SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`
	)
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantRes model.User
		wantErr error
		prepare func(arg args)
	}{
		{
			name: "success get user by email",
			args: args{
				email: "user@email.com",
			},
			wantErr: nil,
			prepare: func(arg args) {
				s.mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(arg.email, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("user_id"))
			},
			wantRes: model.User{
				ID: "user_id",
			},
		},
		{
			name: "Error get user by email",
			args: args{
				email: "user@email.com",
			},
			wantErr: errors.New("internal server error"),
			prepare: func(arg args) {
				s.mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(arg.email, 1).
					WillReturnError(errors.New("internal server error"))
			},
			wantRes: model.User{},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.prepare(tt.args)
			result, err := s.repository.GetUserByEmail(context.Background(), tt.args.email)
			if tt.wantErr == nil {
				assert.Nil(s.T(), err, "should be not err")
			} else {
				assert.Equal(s.T(), err.Error(), err.Error())
				return
			}

			if diff := deep.Equal(tt.wantRes, result); diff != nil {
				s.T().Errorf("got unexpeted result.\nexpect: %v\nactual: %v\ndiff: %v", tt.wantRes, result, diff)
			}
		})
	}
}
