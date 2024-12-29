package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	appSetup "github.com/dating-app-service/cmd/setup"
	"github.com/dating-app-service/config"
	"github.com/dating-app-service/constants"
	authRoutes "github.com/dating-app-service/internal/auth/routes"
	"github.com/dating-app-service/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartServer(setupData appSetup.SetupData) {
	conf := config.GetConfig()
	// appName := conf.App.Name
	if conf.App.Env == constants.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	// GIN Init
	router := gin.Default()
	router.UseRawPath = true

	// swag.Register(swag.Name, &doc.SwagDoc{})
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/app", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Dating App Service is running")
	})
	// router.GET("/health", setupData.InternalApp.Handler.HealthCheckHandler.HealthCheck)

	router.Use(middleware.CORSMiddleware())

	// init public route
	initPublicRoute(router, setupData.InternalApp)

	// router.Use(middleware.JWTAuthMiddleware())
	// router.Use(middleware.AddAppGroupIDs())
	// router.Use(middleware.SetPointName(setupData.ExternalApp.CustomerService))

	// //Init Main APP and Route
	// initRoute(router, setupData.InternalApp)
	// initInternalRoute(router, setupData.InternalApp)

	port := config.GetConfig().Http.Port
	httpServer := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := httpServer.ListenAndServe(); err != nil {
			logrus.Error(fmt.Printf("listen: %s\n", err))
		}
	}()
	logrus.Info("webserver started")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit

	logrus.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logrus.Panic("Server Shutdown:", err)
	}

	_ = appSetup.CloseDB()

	logrus.Info("Server exiting")
}

func initPublicRoute(router *gin.Engine, internalAppStruct appSetup.InternalAppStruct) {
	apiRouter := router.Group("/api/v1/publics")
	authRoutes.PublicRoutes.NewPublicRoutes(apiRouter.Group("/auth"), internalAppStruct.Handler.SignUpHandler)
}
