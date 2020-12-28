package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"time"
)

type fileType string

const FileTypeToml fileType = "toml"
const FileTypeYaml fileType = "yaml"
const FileTypeJson fileType = "json"

type AppConfig interface {
	GetPort() int
	Tracing() bool
	SetPort(port string)
	GetName() string
	SetLogLevel(logLevel string)
	GetLogLevel() string
}

type APP struct {
	PORT     string
	NAME     string
	VERSION  string
	TRACING  string
	LOGLEVEL string
}

func (a *APP) GetPort() int {
	var port, err = strconv.Atoi(a.PORT)
	if err != nil {
		return 80
	}
	return port
}

func (a *APP) Tracing() bool {
	var result, _ = strconv.ParseBool(a.TRACING)
	return result
}

func (a *APP) SetPort(port string) {
	a.PORT = port
}

func (a *APP) GetName() string {
	return a.NAME
}

func (a *APP) SetLogLevel(logLevel string) {
	a.LOGLEVEL = logLevel
}

func (a *APP) GetLogLevel() string {
	return a.LOGLEVEL
}

type DB struct {
	DIALECT          string
	HOST             string
	PORT             string
	USERNAME         string
	PASSWORD         string
	DATABASE         string
	TIMEOUT          string
	PARSETIME        string
	MAXIDLE_CONN     int
	MAXOPEN_CONN     int
	MAXCONN_LIFETIME int // minute
}

type AWS struct {
	REGION string
	ID     string
	KEY    string
	TOKEN  string
}

type SECRET struct {
	TYPE    string
	NAME    string
	VERSION string
}

type REDIS struct {
	ADDR           string
	PASSWORD       string
	DB             string
	POOLSIZE       string
	DURATION       string
	MARSHALINDENT  string
	MODE           string
	ROUTEBYLATENCY string
	ROUTERANDOMLY  string
	DIALTIMEOUT    string
}

func (r REDIS) GetDuration() (*time.Duration, error) {
	duration, err := time.ParseDuration(r.DURATION)
	if err != nil {
		return nil, err
	}
	return &duration, nil
}

func (r REDIS) IsMarshalIndent() bool {
	if val, err := strconv.ParseBool(r.MARSHALINDENT); err == nil {
		return val
	} else {
		return false
	}
}

func (r REDIS) RouteByLatency() bool {
	if val, err := strconv.ParseBool(r.ROUTEBYLATENCY); err == nil {
		return val
	} else {
		return false
	}
}

func (r REDIS) RouteRandomly() bool {
	if val, err := strconv.ParseBool(r.ROUTERANDOMLY); err == nil {
		return val
	} else {
		return false
	}
}

func (r REDIS) GetPoolSize() int {
	if i, err := strconv.Atoi(r.POOLSIZE); err != nil {
		return 10
	} else {
		return i
	}
}

func (r REDIS) GetDialTimeOut() time.Duration {
	duration, err := time.ParseDuration(r.DIALTIMEOUT)
	if err != nil {
		return 0
	}
	return duration
}

func (r REDIS) GetDB() int {
	value, err := strconv.Atoi(r.DB)
	if err != nil {
		return 0
	}
	return value
}

type PARAMETER struct {
	NAME     string
	SECURED  string
	REGION   string
	CREDTYPE string
}

func (p PARAMETER) Secured() bool {
	value, err := strconv.ParseBool(p.SECURED)
	if err != nil {
		return false
	}
	return value
}

func SetEnvironmentServerPort(config interface{}) {
	port := os.Getenv("SERVER_PORT")
	if appConfig, ok := config.(APP); ok && port != "" {
		appConfig.SetPort(port)
		config = appConfig
	}
}

// Init config from specified path
func Init(path string, config interface{}) error {
	return InitWithType(path, config, FileTypeToml)
}

// InitWithType initialize configuration from file with specified FileType
func InitWithType(configPath string, config interface{}, fileType fileType) error {
	env := os.Getenv("ENV")
	if env != "" {
		if env == "prod" {
			gin.SetMode("release")
		}
		env = fmt.Sprintf("%s%s", "-", env)
	}

	viper.SetConfigType(string(fileType))
	viper.SetConfigName(fmt.Sprintf("application%s", env))
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(config); err != nil {
		return err
	}

	SetEnvironmentServerPort(config)
	return nil
}
