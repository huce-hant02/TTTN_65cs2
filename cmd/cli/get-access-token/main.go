package main

import (
	"flag"
	"mono-base/internal/usecases/user"
	"mono-base/pkg/config"
	"mono-base/pkg/constants"
	"mono-base/pkg/logger"
)

type App struct {
	Name           string
	Version        string
	ConfigFilePath string
	ConfigFile     string
	loginUseCase   user.LoginUseCase
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

func inject(
	app *App,
	loginUseCase user.LoginUseCase,
) error {
	app.loginUseCase = loginUseCase
	return nil
}

func (a *App) Run() error {
	ctx := logger.NewBackgroundContextWithTraceID("get-access-token")
	ctxLogger := logger.NewLogger(ctx)
	result, err := a.loginUseCase.Execute(ctx, user.LoginInput{
		Username: "test",
		Password: "test",
	})
	if err != nil {
		return err
	}
	ctxLogger.Infof("User Token: %s", result.AccessToken)

	return nil
}

func main() {
	app := &App{}
	app.initFlag()
	app.initConfig()
	err := wireApp(app)
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
