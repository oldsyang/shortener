package initialize

import "go.uber.org/zap"

func InitLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"./user-web-logger.log",
		"stderr",
		"stdout",
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}
