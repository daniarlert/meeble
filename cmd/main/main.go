package main

import (
	"fmt"
	"meeble/pkg/cli"
	"os"
)

func main() {
	app, err := cli.NewApp("meeble", "Keep track of your mood", "v0.0.1")
	if err != nil {
		fmt.Printf("there was an unexpected error while setting the app: %v\n", err)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("there was an unexpected error: %v\n", err)
	}
}
