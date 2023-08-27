package main

import (
	"fmt"

	"github.com/caiostarke/hermes/text-ingestor/configs"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.DBDriver)
}
