package main

import (
	"fmt"

	"github.com/Twsouza/code-delivery-simulator/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringsJson, _ := route.ExportJsonPositions()
	fmt.Println(stringsJson[0])
}
