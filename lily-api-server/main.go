package main

import (
	"fmt"

	"github.com/yuripedia/lily-api-server/lily-api-server/config"
)

func main() {
	appConfig := config.GetAppConfig()
	fmt.Println(*appConfig)
	rdbConfig := config.GetRdbConfig()
	fmt.Println(*rdbConfig)
}
