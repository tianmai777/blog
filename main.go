package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tianmai777/blog/pkg/logger"
	"github.com/tianmai777/blog/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/gin-gonic/gin"
	"github.com/tianmai777/blog/global"
	"github.com/tianmai777/blog/internal/model"
	"github.com/tianmai777/blog/internal/routers"
	"github.com/tianmai777/blog/pkg"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("setup setting failed: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("setup db engine failed: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("setup log failed: %v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init setupTracer err: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := pkg.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	date := time.Now().Format("20060102")
	global.Log = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + date + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog", "127.0.0.1:6831")
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}
