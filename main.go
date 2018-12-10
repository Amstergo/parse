package main

import (
	"app/news"
	"app/router"
)

func main() {
	r := router.New()
	a := news.New()

	go a.Serve()
	r.Run()
}
