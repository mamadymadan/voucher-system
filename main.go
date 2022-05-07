package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mamadymadan/voucher-system/models"
	"github.com/mamadymadan/voucher-system/pkg/logging"
	"github.com/mamadymadan/voucher-system/pkg/setting"
	"github.com/mamadymadan/voucher-system/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	routersInit := routers.InitRouter()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
