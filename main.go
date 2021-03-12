package main

import (
	"code-analyser/versions"
	"log"
)

func main() {
	log.Println(versions.ParserVersion("versions/go.yaml"))
}
