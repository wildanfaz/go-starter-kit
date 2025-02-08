package main

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wildanfaz/go-starter-kit/cmd"
)

func main() {
	cmd.InitCmd(context.Background())
}
