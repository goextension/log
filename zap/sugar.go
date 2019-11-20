package zap

import (
	"github.com/goextension/log"
	"go.uber.org/zap"
)

var DefaultZapFilePath = "zap.log"

// InitGlobalZapSugar ...
func InitZapSugar() {
	logger, e := zap.NewProduction(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if e != nil {
		panic(e)
	}
	log.Register(logger.Sugar())
}

func InitZapFileSugar() {
	cfg := zap.NewProductionConfig()
	//p, _ := os.Getwd()
	//p = filepath.Join(p, "zap.log")
	//os.OpenFile(p,os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC,os.ModePerm)

	cfg.OutputPaths = []string{
		DefaultZapFilePath,
	}

	logger, e := cfg.Build()
	if e != nil {
		panic(e)
	}
	log.Register(logger.Sugar())
}

// NewZapFile ...
func NewZapFile() *zap.Logger {
	cfg := zap.NewProductionConfig()
	//p, _ := os.Getwd()
	//p = filepath.Join(p, "zap.log")
	//os.OpenFile(p,os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC,os.ModePerm)

	cfg.OutputPaths = []string{
		DefaultZapFilePath,
	}

	logger, e := cfg.Build()
	if e != nil {
		panic(e)
	}
	return logger
}

//NewZapFileSugar ...
func NewZapFileSugar() *zap.SugaredLogger {
	return NewZapFile().Sugar()
}
