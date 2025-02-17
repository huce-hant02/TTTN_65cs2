package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	httpConfig "mono-base/cmd/http/config"
	edit_history "mono-base/cmd/http/controllers/edit-history"
	"mono-base/cmd/http/controllers/user"
	_ "mono-base/cmd/http/docs"
	"mono-base/cmd/http/middleware"
	"mono-base/pkg/config"
	"mono-base/pkg/constants"
	"net/http"
)

type App struct {
	Name                    string
	Version                 string
	ConfigFilePath          string
	ConfigFile              string
	router                  *gin.Engine
	restConfig              httpConfig.RestServer
	userControllerV1        user.Controller
	userControllerV2        user.Controller
	editHistoryControllerV1 edit_history.Controller
}

func (a *App) initFlag() {
	flag.StringVar(&a.Name, "name", "service-name", "")
	flag.StringVar(&a.Version, "version", "1.0.0", "")
	flag.StringVar(&a.ConfigFilePath, "config-file-path", "./configs", "Config file path: path to config dir")
	flag.StringVar(&a.ConfigFile, "config-file", "config", "Config file path: path to config dir")
	flag.Parse()
}

func (a *App) initConfig() {
	configSource := &config.Viper{
		ConfigType: constants.ConfigTypeFile,
		FilePath:   a.ConfigFilePath,
		ConfigFile: a.ConfigFile,
	}
	err := configSource.InitConfig()
	if err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	a.registerRoute()
	err := a.router.Run(fmt.Sprintf("%s:%s", a.restConfig.Path, a.restConfig.Port))
	if err != nil {
		return err
	}
	return nil
}

func (a *App) registerRoute() {
	// Routes
	a.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Swagger route
	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user.RegisterRoutesV1(a.router, a.userControllerV1)
	edit_history.RegisterRoutesV1(a.router, a.editHistoryControllerV1)
	user.RegisterRoutesV2(a.router, a.userControllerV2)
}

func inject(
	app *App,
	userControllerV1 *user.ControllerV1,
	editHistoryControllerV1 *edit_history.Controller,
	userControllerV2 *user.ControllerV2,
) error {
	app.userControllerV1 = userControllerV1
	editHistoryControllerV1 = editHistoryControllerV1
	app.userControllerV2 = userControllerV2
	return nil
}

// @title Devices manager API
// @version 1.0
// @description Restfull API Application for web devices management
// @scheme https
// @host devices-stg.phx-smartuni.com
// @BasePath /api
func main() {
	app := &App{}
	app.initFlag()
	app.initConfig()
	restConfig := httpConfig.RestServer{}
	err := viper.UnmarshalKey("http", &restConfig)
	if err != nil {
		panic(err)
	}
	app.restConfig = restConfig

	// Init gin router
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(middleware.JsonLogMiddleware), gin.Recovery(), middleware.CorsMiddleware())
	router.HandleMethodNotAllowed = true
	router.NoMethod(
		func(context *gin.Context) {
			context.AbortWithStatus(http.StatusMethodNotAllowed)
		},
	)
	app.router = router

	err = wireApp(app)
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}

}
