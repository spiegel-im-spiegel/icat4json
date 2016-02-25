package main

import (
	"fmt"
	"log"

	"github.com/spiegel-im-spiegel/icat4json"
)

func main() {
	json, err := icat4json.Get(icat4json.ToolICATW)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Decode()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Title: %v\n", data.Title)
	fmt.Printf("  URL: %v\n", data.Link)
	fmt.Printf(" Date: %v\n", data.Date)
	fmt.Print("Items:\n")
	for _, item := range data.Itemdata {
		fmt.Printf("\t%v: %v (%v)\n", item.Date, item.Title, item.Link)
	}
}
