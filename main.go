package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"omo-bla-acm/config"
	"omo-bla-acm/handler"
	"os"
	"path/filepath"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
)

func main() {
	config.Setup()

	// New Service
	service := micro.NewService(
		micro.Name(config.Schema.Service.Name),
		micro.Version(BuildVersion),
		micro.RegisterTTL(time.Second*time.Duration(config.Schema.Service.TTL)),
		micro.RegisterInterval(time.Second*time.Duration(config.Schema.Service.Interval)),
		micro.Address(config.Schema.Service.Address),
	)

	// Initialise service
	service.Init()

	// Initialise service
	cli := service.Client()
	cli.Init(
		client.Retries(int(config.Schema.Client.Retry)),
		client.RequestTimeout(time.Second*time.Duration(config.Schema.Client.Timeout)),
		client.Retry(clientRetry),
	)

	accountHandler := new(handler.Account)
	accountHandler.AuthService = proto.NewAuthService(config.Schema.MSA.Account, cli)
	accountHandler.ProfileService = proto.NewProfileService(config.Schema.MSA.Account, cli)

	// Register Handler
	proto.RegisterAuthHandler(service.Server(), accountHandler)
	proto.RegisterProfileHandler(service.Server(), accountHandler)

	app, _ := filepath.Abs(os.Args[0])

	logger.Info("-------------------------------------------------------------")
	logger.Info("- Business Logic Agent -> Run")
	logger.Info("-------------------------------------------------------------")
	logger.Infof("- version      : %s", BuildVersion)
	logger.Infof("- application  : %s", app)
	logger.Infof("- md5          : %s", md5hex(app))
	logger.Infof("- build        : %s", BuildTime)
	logger.Infof("- commit       : %s", CommitID)
	logger.Info("-------------------------------------------------------------")
	// Run service
	if err := service.Run(); err != nil {
		logger.Error(err)
	}
}

func md5hex(_file string) string {
	h := md5.New()

	f, err := os.Open(_file)
	if err != nil {
		return ""
	}
	defer f.Close()

	io.Copy(h, f)

	return hex.EncodeToString(h.Sum(nil))
}

func clientRetry(_ctx context.Context, _req client.Request, _retryCount int, _err error) (bool, error) {
	if nil != _err {
		logger.Warnf("retry %d, reason is %v\n\r", _retryCount, _err)
		return true, nil
	}
	return false, nil
}
