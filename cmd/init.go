package cmd

import (
	"context"
	"fmt"

	"github.com/OZahed/restmp/internal/configs"
	"github.com/OZahed/restmp/internal/http"
	"github.com/OZahed/restmp/internal/log"
)

func ExecuteRoot(ctx context.Context) {

	if err := log.Initialize(ctx, log.LOG_LEVEL(log.L_DEBUG)); err != nil {
		fmt.Println(err.Error())
		return
	}
	appConf := &configs.AppConfig{}
	if err := appConf.Initialize(context.Background()); err != nil {
		log.Logger.Error(fmt.Sprintf("could not make config %v", err))
		return
	}

	S := http.NewServer()
	S.ListenAndServe(context.Background(), ":9000")
	log.Logger.Error("Done")
}
