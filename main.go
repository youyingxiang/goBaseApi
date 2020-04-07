package main

import (
	Mysql "rbac/databases"
	"rbac/router"
)

func main()  {
	defer Mysql.DB.Close()
	router.InitRouter()
}
