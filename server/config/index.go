package config

import (
	"log"

	"github.com/spf13/viper"
)

type environment struct {
	Instance *viper.Viper
}

func NewEnvironment(absPath, filename string) *environment {
	v := viper.New()
	v.AddConfigPath(absPath)
	v.SetConfigName(filename)

	log.Println("NewEnvironment.go: ", absPath)

	if rErr := v.ReadInConfig(); rErr != nil {
		log.Fatalf("NewEnvironment.go: failed to init viper: %s", rErr.Error())
	}

	return &environment{
		Instance: v,
	}
}
