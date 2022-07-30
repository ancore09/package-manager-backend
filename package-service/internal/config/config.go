package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Build information -ldflags .
var (
	version    string = "dev"
	commitHash string = "-"
)

var cfg *Config

func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

// Grpc - contains parameter address grpc.
type Grpc struct {
	Port              int    `yaml:"port"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	Timeout           int64  `yaml:"timeout"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Host              string `yaml:"host"`
}

// Gateway - contains parameters for grpc-gateway port
type Gateway struct {
	Port               int      `yaml:"port"`
	Host               string   `yaml:"host"`
	AllowedCORSOrigins []string `yaml:"allowedCorsOrigins"`
}

// Swagger - contains parameters for swagger port
type Swagger struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Filepath string `yaml:"filepath"`
}

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Version     string
	CommitHash  string
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project             Project `yaml:"project"`
	Grpc                Grpc    `yaml:"grpc"`
	Gateway             Gateway `yaml:"gateway"`
	Swagger             Swagger `yaml:"swagger"`
	CategoryServiceAddr string  `yaml:"categoryServiceAddr"`
	DB                  DB      `yaml:"db"`
}

func (c Config) GetDSN() string {
	return c.DB.DSN
}

func (c Config) GetMaxConns() int {
	return c.DB.MaxOpenConns
}

func (c Config) GetMaxIdleConns() int {
	return c.DB.MaxIdleConns
}

func (c Config) GetConnMaxIdleTime() time.Duration {
	return c.DB.ConnMaxIdleTime
}

func (c Config) GetConnMaxLifeTime() time.Duration {
	return c.DB.ConnMaxLifeTime
}

type DB struct {
	DSN             string        `yaml:"DSN"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"MaxIdleConns"`
	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime"`
	ConnMaxLifeTime time.Duration `yaml:"connMaxLifeTime"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(configYML string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(configYML)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return nil
}
