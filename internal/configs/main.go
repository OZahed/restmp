package configs

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct{}

func (AppConfig) Initialize(_ context.Context) error {
	viper.AddConfigPath(fmt.Sprintf("/tmp/%s", AppName()))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.config/%s", AppName()))
	viper.AddConfigPath(".")
	viper.SetEnvPrefix(AppName())
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	if viper.GetString("tz") != "" {
		z, err := time.LoadLocation(viper.GetString("tz"))
		if err != nil {
			return err
		}
		time.Local = z

	}
	return nil
}

func (AppConfig) Finalize() error {
	return nil
}
