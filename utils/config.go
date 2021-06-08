package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"strings"
)

type Config struct {
	ServerAddress      string
	DBHost string
	DBUser string
	DBPassword string
	DBName string
	DBPort string
	DBSSL string
	DBConnectionString string
}

//LoadConfig loads config from path=p
func LoadConfig(p string) (*Config, error) {
	dir := path.Dir(p)
	_, file := path.Split(p)
	parts := strings.Split(file, ".")
	filename := parts[0]
	extension := parts[1]

	viper.AddConfigPath(dir)
	viper.SetConfigName(filename)
	viper.SetConfigType(extension)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := Config{}

	viper.SetDefault("DB.HOST", "localhost")
	viper.SetDefault("DB.USER", "postgres")
	viper.SetDefault("DB.PASSWORD", "password")
	viper.SetDefault("DB.NAME", "balance")
	viper.SetDefault("DB.PORT", "5432")
	viper.SetDefault("DB.SSL", "disable")
	config.DBHost = fmt.Sprintf("%v", viper.Get("DB.HOST"))
	config.DBUser = fmt.Sprintf("%v", viper.Get("DB.USER"))
	config.DBPassword = fmt.Sprintf("%v", viper.Get("DB.PASSWORD"))
	config.DBName = fmt.Sprintf("%v", viper.Get("DB.NAME"))
	config.DBPort = fmt.Sprintf("%v", viper.Get("DB.PORT"))
	config.DBSSL = fmt.Sprintf("%v", viper.Get("DB.SSL"))

	viper.SetDefault("SERVER.PORT", "8080")
	srvPort := viper.Get("SERVER.PORT")
	config.ServerAddress = fmt.Sprintf(":%v", srvPort)

	return &config, nil
}

func CreatePGSQLConnString (config *Config) {
	config.DBConnectionString = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSL)
}