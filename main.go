package main

import (
	sungoqAPI "github.com/hadihammurabi/sungoq/api"
	sungoqService "github.com/hadihammurabi/sungoq/service"
)

func main() {

	svc, err := sungoqService.New()
	if err != nil {
		panic(err)
	}

	api, err := sungoqAPI.New(
		sungoqAPI.WithService(svc),
	)
	if err != nil {
		panic(err)
	}

	api.Start()

}
