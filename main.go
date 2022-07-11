package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"meeble/models"
	"os"
)

func main() {
	var mood string
	var reason string

	app := &cli.App{
		Name:                 "meeble",
		Usage:                "keep track of your mood",
		Version:              "v0.0.1",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mood",
				Aliases:     []string{"m"},
				Value:       "",
				Usage:       "your current mood",
				Destination: &mood,
			},
			&cli.StringFlag{
				Name:        "reason",
				Aliases:     []string{"r"},
				Value:       "",
				Usage:       "the reason of your mood",
				Destination: &reason,
			},
		},
		Action: func(ctx *cli.Context) error {
			if mood == "" {
				return errors.New("you have to specify what is your mood")
			}

			record := models.NewRecord(mood, reason)
			fmt.Printf("record created successfully: %s: %s", record.Mood, record.Reason)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
