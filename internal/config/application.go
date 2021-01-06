package config

import (
	"github.com/unit-test-example/pkg/config"
)

var Path = "./configs"

const ErrMsgUnableToInitConfig = "Unable to init config"

type Application struct {
	Server   config.APP
	Database config.DB
}

func Init() (*Application, error) {
	var configs = new(Application)
	if err := config.Init(Path, configs); err != nil {
		return nil, err
	}
	//log.Setup(ServiceName, ServiceVersion, configs.Server.GetLogLevel(), MessageFieldName)
	return configs, nil
}
