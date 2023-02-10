package main

import (
	"server/initializers"
	"server/routes"
)

func main() {

	initializers.ConnectToBD()
	initializers.Migration()

	e := routes.Init()


	e.Logger.Fatal(e.Start(":8050"))


}
