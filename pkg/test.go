package pkg

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dating-app-service/config/db"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

type DatabaseMock struct {
	GormDB  *db.GormDB
	SQLMock sqlmock.Sqlmock
}

func ConnectDB() (*DatabaseMock, error) {
	var (
		sqlDB *sql.DB
		err   error
	)
	dbMock := &DatabaseMock{}

	sqlDB, dbMock.SQLMock, err = sqlmock.New()
	if err != nil {
		return dbMock, errors.New("failed to open mock sql db")
	}

	if sqlDB == nil {
		return dbMock, errors.New("mock db is null")
	}

	if dbMock.SQLMock == nil {
		return dbMock, errors.New("sqlmock is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return dbMock, errors.New("failed to open gorm v2 db")
	}

	dbMock.GormDB = &db.GormDB{DB: gormDB}

	if dbMock.GormDB == nil {
		return dbMock, errors.New("gorm db is null")
	}

	return dbMock, nil
}
