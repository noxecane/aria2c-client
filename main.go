package main

import (
	"fmt"
	"os"

	"github.com/noxecane/aria2c-client/aria2c"
)

func main() {
	url := os.Args[1]
	r := make(chan aria2c.RPCResult)

	// ask for download
	go aria2c.AddDownload(r, url)

	// handle results
	result := <-r
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("GID", result.Result)
}
