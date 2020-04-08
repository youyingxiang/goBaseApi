package main

import (
	"rbac/databases"
	"rbac/router"
)

func main()  {
	defer Mysql.DB.Close()
	router.InitRouter()
}
