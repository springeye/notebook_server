package main

import (
	"fmt"
	conf "notebook/config"
	"notebook/router"
)

func main() {

	router.SetupServer().Run(fmt.Sprintf("0.0.0.0:%d", conf.Conf.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
