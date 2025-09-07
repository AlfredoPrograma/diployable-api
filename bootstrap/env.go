package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

// Env holds configuration values for the application's environment,
// The struct fields are mapped from environment variables using mapstructure tags.
type Env struct {
	ApiPort            int    `mapstructure:"API_PORT"`
	DbConnectionString string `mapstructure:"DB_CONNECTION_STRING"`
}

// LoadEnv loads environment variables from a .env file using Viper,
// unmarshals them into an Env struct, and returns the struct.
// If the .env file cannot be read or unmarshaled, the function logs the error and terminates the application.
func LoadEnv() Env {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("cannot read environment variables from .env file")
		log.Fatalln(err)
	}

	var env Env

	if err := viper.Unmarshal(&env); err != nil {
		log.Println("cannot unmarshal environments variavle from .env file into Env struct")
		log.Fatalln(err)
	}

	return env
}
