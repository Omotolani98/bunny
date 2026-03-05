package main

import (
	"context"
	"os"

	"github.com/Omotolani98/bunny/cmd"
	"github.com/charmbracelet/fang"
)

func main() {
	if err := fang.Execute(context.Background(), cmd.Init()); err != nil {
		os.Exit(1)
	}
}
