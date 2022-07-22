package configs

import(
	"github.com/spf13/viper"
)

var cfg *config

type config struct{
	API APIConfig
	DB DBConfig
}

// Porta da api
type APIConfig struct{
	Port string
}

// Configurações do banco de dados
type DBConfig struct{
	Host string
	Port string
	User string
	Password string
	Database string
}

// incialização das configurações da api
func init(){
	viper.SetDefault("api.port", ":5000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

// Load carrega as configurações da api a partir do arquivo toml
func Load() error{
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil{
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok{
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		User: viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

// GetDB retorna as configurações do banco de dados
func GetDB() DBConfig{
	return cfg.DB
}

// GetPortaServidor retorna a porta do servidor
func GetPortaServidor() string{
	return cfg.API.Port
}