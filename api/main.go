package main

import "github.com/mopeneko/vshuki/api/router"

func main() {
	e := router.Init()

	e.Logger.Fatal(e.Start(":4000"))
}
