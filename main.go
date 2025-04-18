package main

import (
	"fmt"
	"github.com/benjaminafoster/gator/internal/config"
)

func main() {
	cfg, _ := config.Read()
	cfg.SetUser("benjamin")
	cfg, _ = config.Read()

	fmt.Println(cfg)
}