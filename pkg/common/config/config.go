package config

import "github.com/spf13/viper"

type config struct {
	Port  string `mapstructure:"PORT"`
	DBurl string `mapstructure:"DB_URL"`
}

func loadConfig(c config, err error) {
	viper.AddConfigPath("./pkg/common/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadRemoteConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	return
}
