package main

import (
	"fmt"
	"task1/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
