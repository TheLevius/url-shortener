package main

import (
	"fmt"
	"url-shortener/cmd/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// TODO: init logger: slog
	// TODO: init storage: sqLite
	// TODO: init router: chi
	// TODO: run server: 
}
