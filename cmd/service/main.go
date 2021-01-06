package main

import (
	"fmt"
	"github.com/google/martian/log"
	"github.com/unit-test-example/internal/config"
	"github.com/unit-test-example/internal/router"
)

func main() {

	// initialize application log service name and message field name
	var configs, err = config.Init()
	if err != nil {
		log.Errorf(config.ErrMsgUnableToInitConfig, err)
	}

	// initialize router
	ginRouter := router.Init(configs)
	if err := ginRouter.Run(fmt.Sprintf(":%d", configs.Server.GetPort())); err != nil {
		log.Errorf("Unable to start application", err)
	}
}
