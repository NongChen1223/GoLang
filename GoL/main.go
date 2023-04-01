package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	// Run("里面不指定端口号默认为8080")
	r := router.Router()
	r.Run(":8081") // listen and serve on
}
