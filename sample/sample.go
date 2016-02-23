package main

import (
	"log"

	"github.com/spiegel-im-spiegel/icat4json"
)

func main() {
	json, err := icat4json.Get(icat4json.ToolICATW)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(json)
}
