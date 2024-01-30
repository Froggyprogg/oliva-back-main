package main

import (
	"fmt"
	"oliva-back-main/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	//TODO: init storage: postgres

	//TODO: init router

	//TODO: run server
}
