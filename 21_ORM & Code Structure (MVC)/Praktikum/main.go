package main

import config "main.go/Config"

func main() {
	config.Init()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
