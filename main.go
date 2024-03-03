package main

import (
	"context"
	"fmt"

	"github.com/gracefulm/go-pkl-sample/config"
)

func main() {
	cfg, err := config.LoadFromPath(context.Background(), "config/pkl/dev/config.pkl")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Host=%s\n", cfg.AppConfig.Host)
	fmt.Printf("Port=%d\n", cfg.AppConfig.Port)
	fmt.Printf("LogLevel=%s\n", cfg.AppConfig.LogConfig.LogLevel)
}
