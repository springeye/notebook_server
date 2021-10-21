package main

import (
	"fmt"
	conf "notebook/config"
	"notebook/router"
)

/**
 * install swage tools
 * go get -v -u github.com/swaggo/swag/cmd/swag
 */
//go:generate swag init -g server.go
// @securityDefinitions.apikey user_token
// @in header
// @name Authorization
func main() {

	router.SetupServer().Run(fmt.Sprintf("0.0.0.0:%d", conf.Conf.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
