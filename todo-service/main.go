package main

import (
	"fmt"
	"os"

	"todo-service/config"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("app cannot start, err:%v",err)
		os.Exit(0)
	}
}

func run() error {
	_, err := config.NewPropertiesConfigProvider()
	if err != nil {
		return err
	}
	return nil
}