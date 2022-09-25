package main

import (
	"log"

	sungoqAPI "github.com/hadihammurabi/sungoq/api"
	sungoqService "github.com/hadihammurabi/sungoq/service"
)

func main() {

	_, err := sungoqService.New()
	if err != nil {
		log.Fatal(err)
	}

	api := sungoqAPI.New()
	api.Start()

}
