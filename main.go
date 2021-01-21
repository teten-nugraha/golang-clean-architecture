package main

import (
	routes2 "github.com/teten-nugraha/golang-crud/routes"
)

func main() {

	routes := routes2.Init()

	routes.Logger.Fatal(routes.Start(":9999"))
}
