package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dating-app-service/internal/premium/port"
	"github.com/dating-app-service/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository port.IPremiumRepo
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

func (s *Suite) Test_repository_UpdateUserPremium() {

	var (
		query = `UPDATE "users" SET "is_premium"=$1 WHERE email = $2`
	)
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		prepare func(arg args)
	}{
		{
			name: "success update user premium",
			args: args{
				email: "user@email.com",
			},
			wantErr: nil,
			prepare: func(arg args) {
				s.mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "Error update user premium",
			args: args{
				email: "user@email.com",
			},
			wantErr: errors.New("internal server error"),
			prepare: func(arg args) {
				s.mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("internal server error"))
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.prepare(tt.args)
			err := s.repository.UpdateUserPremium(context.Background(), tt.args.email)
			if tt.wantErr == nil {
				assert.Nil(s.T(), err, "should be not err")
			} else {
				assert.Equal(s.T(), err.Error(), err.Error())
				return
			}
		})
	}
}
