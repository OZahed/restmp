package cmd

import (
	"context"

	"github.com/OZahed/restmp/internal/configs"
	"github.com/OZahed/restmp/internal/log"
)

func ExecuteRoot(ctx context.Context) {
	appConf := &configs.AppConfig{}
	if err := appConf.Initialize(context.Background()); err != nil {
		log.Errorf("could not make config %v", err)
	}

	
}
