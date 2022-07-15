package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"meeble/pkg/db"
	"meeble/pkg/models"
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

			localDB, err := db.NewLocalDB("moods.db")
			if err != nil {
				return err
			}

			m := &models.Mood{Mood: mood, Reason: reason}
			if err := models.CreateMood(localDB.DB, m); err != nil {
				return err
			}

			m, err = models.GetFirstMood(localDB.DB)
			if err != nil {
				return err
			}

			fmt.Println(m)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
