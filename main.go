package main

import (
	"fmt"

	sungoqService "github.com/hadihammurabi/sungoq/service"
)

func main() {
	service, err := sungoqService.New()
	fmt.Println(service.Topic, err)
}
