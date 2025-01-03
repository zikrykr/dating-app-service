package cmd

import (
	"github.com/dating-app-service/config"
	"github.com/dating-app-service/config/db"
	authHandler "github.com/dating-app-service/internal/auth/handler"
	authPorts "github.com/dating-app-service/internal/auth/port"
	authrepo "github.com/dating-app-service/internal/auth/repository"
	authService "github.com/dating-app-service/internal/auth/service"
	premiumHandler "github.com/dating-app-service/internal/premium/handler"
	premiumPorts "github.com/dating-app-service/internal/premium/port"
	premiumRepo "github.com/dating-app-service/internal/premium/repository"
	premiumService "github.com/dating-app-service/internal/premium/service"
	recommendationHandler "github.com/dating-app-service/internal/recommendations/handler"
	recommendationPorts "github.com/dating-app-service/internal/recommendations/port"
	recommendationRepo "github.com/dating-app-service/internal/recommendations/repository"
	recommendationService "github.com/dating-app-service/internal/recommendations/service"
	swipeHandler "github.com/dating-app-service/internal/swipe/handler"
	swipePorts "github.com/dating-app-service/internal/swipe/port"
	swipeRepo "github.com/dating-app-service/internal/swipe/repository"
	swipeService "github.com/dating-app-service/internal/swipe/service"
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
	dbInstance         *gorm.DB
	AuthRepo           authPorts.IAuthRepo
	RecommendationRepo recommendationPorts.IRecommendationRepo
	SwipeRepo          swipePorts.ISwipeRepository
	PremiumRepo        premiumPorts.IPremiumRepo
}

// Services
type initServicesApp struct {
	SignUpService         authPorts.ISignUpService
	LoginService          authPorts.ILoginService
	ProfileService        authPorts.IProfileService
	RecommendationService recommendationPorts.IRecommendationService
	SwipeService          swipePorts.ISwipeService
	PremiumService        premiumPorts.IPremiumService
}

// Handler
type InitHandlerApp struct {
	SignUpHandler         authPorts.ISignUpHandler
	LoginHandler          authPorts.ILoginHandler
	ProfileHandler        authPorts.IProfileHandler
	RecommendationHandler recommendationPorts.IRecommendationHandler
	SwipeHandler          swipePorts.ISwipeHandler
	PremiumHandler        premiumPorts.IPremiumHandler
}

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

	initializeApp.Repositories.AuthRepo = authrepo.NewRepository(gormDB)
	initializeApp.Repositories.RecommendationRepo = recommendationRepo.NewRepository(gormDB)
	initializeApp.Repositories.SwipeRepo = swipeRepo.NewRepository(gormDB)
	initializeApp.Repositories.PremiumRepo = premiumRepo.NewRepository(gormDB)
}

func initAppService(initializeApp *InternalAppStruct) {
	initializeApp.Services.SignUpService = authService.NewSignUpService(initializeApp.Repositories.AuthRepo)
	initializeApp.Services.LoginService = authService.NewLoginService(initializeApp.Repositories.AuthRepo)
	initializeApp.Services.ProfileService = authService.NewProfileService(initializeApp.Repositories.AuthRepo)
	initializeApp.Services.RecommendationService = recommendationService.NewRecommendationService(initializeApp.Repositories.RecommendationRepo, initializeApp.Repositories.AuthRepo, initializeApp.Repositories.SwipeRepo)
	initializeApp.Services.SwipeService = swipeService.NewSwipeService(initializeApp.Repositories.SwipeRepo, initializeApp.Repositories.AuthRepo)
	initializeApp.Services.PremiumService = premiumService.NewPremiumService(initializeApp.Repositories.PremiumRepo, initializeApp.Repositories.AuthRepo)
}

func initAppHandler(initializeApp *InternalAppStruct) {
	initializeApp.Handler.SignUpHandler = authHandler.NewSignUpHandler(initializeApp.Services.SignUpService)
	initializeApp.Handler.LoginHandler = authHandler.NewLoginHandler(initializeApp.Services.LoginService)
	initializeApp.Handler.ProfileHandler = authHandler.NewProfileHandler(initializeApp.Services.ProfileService)
	initializeApp.Handler.RecommendationHandler = recommendationHandler.NewRecommendationHandler(initializeApp.Services.RecommendationService)
	initializeApp.Handler.SwipeHandler = swipeHandler.NewSwipeHandler(initializeApp.Services.SwipeService)
	initializeApp.Handler.PremiumHandler = premiumHandler.NewPremiumHandler(initializeApp.Services.PremiumService)
}
