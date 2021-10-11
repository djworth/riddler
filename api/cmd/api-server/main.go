package main

import (
	"log"

	"github.com/djworth/riddler/pkg/web"
)

func main() {
	log.Fatalln(web.Serve(":3000"))
}
