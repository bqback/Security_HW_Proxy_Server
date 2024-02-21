package config

import (
	"os"
	"proxy_server/internal/apperrors"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

// ServerConfig
// структура для хранения параметров сервера
type Config struct {
	// Session  *SessionConfig  `yaml:"-"`
	Server *ServerConfig `yaml:"server"`
	// CORS     *CORSConfig     `yaml:"cors"`
	Database *DatabaseConfig `yaml:"db"`
	Logging  *LoggingConfig  `yaml:"logging"`
}

type ServerConfig struct {
	BackendPort uint   `yaml:"backend_port"`
	Host        string `yaml:"host"`
}

type DatabaseConfig struct {
	User              string `yaml:"user"`
	Password          string `yaml:"-"`
	Host              string `yaml:"-"`
	Port              uint64 `yaml:"port"`
	DBName            string `yaml:"db_name"`
	AppName           string `yaml:"app_name"`
	Schema            string `yaml:"schema"`
	ConnectionTimeout uint64 `yaml:"connection_timeout"`
}

type LoggingConfig struct {
	Level                  string `yaml:"level"`
	DisableTimestamp       bool   `yaml:"disable_timestamp"`
	FullTimestamp          bool   `yaml:"full_timestamp"`
	DisableLevelTruncation bool   `yaml:"disable_level_truncation"`
	LevelBasedReport       bool   `yaml:"level_based_report"`
	ReportCaller           bool   `yaml:"report_caller"`
}

// LoadConfig
// создаёт конфиг из .env файла, находящегося по полученному пути
func LoadConfig(envPath string, configPath string) (*Config, error) {
	var (
		config Config
		err    error
	)

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	if envPath == "" {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(envPath)
	}

	if err != nil {
		return nil, apperrors.ErrEnvNotFound
	}

	// config.Session, err = NewSessionConfig()
	// if err != nil {
	// 	return nil, err
	// }

	config.Database.Password, err = GetDBPassword()
	if err != nil {
		return nil, err
	}

	config.Database.Host = GetDBConnectionHost()

	return &config, nil
}

// GetDBConnectionHost
// возвращает имя хоста из env для соединения с БД (по умолчанию localhost)
func GetDBConnectionHost() string {
	host, hOk := os.LookupEnv("POSTGRES_HOST")
	if !hOk {
		return "localhost"
	}
	return host
}

// getDBConnectionHost
// возвращает пароль из env для соединения с БД
func GetDBPassword() (string, error) {
	pwd, pOk := os.LookupEnv("POSTGRES_PASSWORD")
	if !pOk {
		return "", apperrors.ErrDatabasePWMissing
	}
	return pwd, nil
}
