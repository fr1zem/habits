package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Database PostgresConfig `mapstructure:"postgres"`
	Server   ServerConfig   `mapstructure:"server"`
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string `mapstructure:"SSLMode"`
}

type ServerConfig struct {
	Host string
	Port int
}

func Init() (*Config, error) {
	setDefaults()

	err := parseConfig("/home/user_ams/Projects/GoProjects/CLIappHabits/configs", "main", "yaml")
	if err != nil {
		return nil, fmt.Errorf("parseConfig: %w", err)
	}

	var cfg Config
	err = unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	setFromEnv(&cfg)

	return &cfg, err

}

func setDefaults() {
	viper.SetDefault("postgres.port", 18080)
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.SSLMode", "disable")
	viper.SetDefault("server.host", "172.24.96.1")
	viper.SetDefault("server.port", 18080)
}

func parseConfig(folder, file, format string) error {

	viper.AddConfigPath(folder)
	viper.SetConfigName(file)
	viper.SetConfigType(format)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("postgres.host", &cfg.Database.Host); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres.port", &cfg.Database.Port); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres.SSLMode", &cfg.Database.SSLMode); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("server.host", &cfg.Server.Host); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("server.port", &cfg.Server.Port); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.Database.DBName = os.Getenv("POSTGRES_DBNAME")
	cfg.Database.User = os.Getenv("POSTGRES_USER")
	cfg.Database.Password = os.Getenv("POSTGRES_PASSWORD")

}
