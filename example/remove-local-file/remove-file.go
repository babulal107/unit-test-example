package main

import (
	"fmt"
	"github.com/unit-test-example/pkg/helper"
	"os"
)

func main() {

	currentDir, _ := os.Getwd()
	fileName := "Indaram_Sales Invoice_1.pdf"
	tmpFileDir := currentDir + "/assets/idstats/" + fileName
	err := helper.RemoveFile(tmpFileDir)
	if err != nil {
		fmt.Println("error while delete file : ", err.Error())
	}

}
