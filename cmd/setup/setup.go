package cmd

import (
	"github.com/dating-app-service/config"
	"github.com/dating-app-service/config/db"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SetupData struct {
	ConfigData  config.Config
	InternalApp InternalAppStruct
}

type InternalAppStruct struct {
	Repositories initRepositoriesApp
	Services     initServicesApp
	Handler      InitHandlerApp
}

// Repositories
type initRepositoriesApp struct {
	dbInstance *gorm.DB
}

// Services
type initServicesApp struct {
}

// Handler
type InitHandlerApp struct {
}

// BaseURL base url of api
const BaseURL = "/v1/api"

// CloseDB close connection to db
var CloseDB func() error

func InitSetup() SetupData {
	configData := config.GetConfig()

	//DB INIT
	dbConn, err := db.Init()
	if err != nil {
		logrus.Fatal("database error", err)
	}

	CloseDB = func() error {
		if err := dbConn.CloseConnection(); err != nil {
			return err
		}

		return nil
	}

	internalAppVar := initInternalApp(dbConn.GormDB)

	return SetupData{
		ConfigData:  configData,
		InternalApp: internalAppVar,
	}
}

func initInternalApp(gormDB *db.GormDB) InternalAppStruct {
	var internalAppVar InternalAppStruct

	initAppRepo(gormDB, &internalAppVar)
	initAppService(&internalAppVar)
	initAppHandler(&internalAppVar)

	return internalAppVar
}

func initAppRepo(gormDB *db.GormDB, initializeApp *InternalAppStruct) {
	// Get Gorm instance
	initializeApp.Repositories.dbInstance = gormDB.DB
}

func initAppService(initializeApp *InternalAppStruct) {
	// initializeApp.Services.PromoService = promoService.NewService(initializeApp.Repositories.PromoRepo, initializeApp.Repositories.TrxHandler, externalApp.CRMService)
}

func initAppHandler(initializeApp *InternalAppStruct) {
	// initializeApp.Handler.PromoHandler = promoHandler.NewHandler(promoHandler.PromoHandlerOptions{
	// 	PromoService:           initializeApp.Services.PromoService,
	// 	PromoServiceV2:         initializeApp.Services.PromoV2Service,
	// 	UserService:            initializeApp.Services.UserService,
	// 	RuleService:            initializeApp.Services.RuleService,
	// 	UsageService:           initializeApp.Services.UsageService,
	// 	TrxHandler:             initializeApp.Repositories.TrxHandler,
	// 	AuditlogService:        externalApp.AuditLog,
	// 	ScaleConversionService: externalApp.ScaleConversionService,
	// 	BudgetService:          initializeApp.Services.BudgetService,
	// 	CRMService:             externalApp.CRMService,
	// })

}
