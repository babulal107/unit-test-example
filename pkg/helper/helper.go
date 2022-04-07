package helper

import (
	"fmt"
	"os"
)

// RemoveFile ... this function is used to delete file from local system
func RemoveFile(filePath string) error {
	fmt.Println(fmt.Sprintf("requested file path : %s", filePath))

	err := os.Remove(filePath)
	if err != nil {
		fmt.Println(fmt.Sprintf("error unable to delete file path : %s", filePath))
		return err
	}
	fmt.Println("File successfully deleted")
	return nil
}