package cli

import (
	"github.com/urfave/cli/v2"
	"meeble/pkg/db"
)

func NewApp(name, usage, version string) (*cli.App, error) {
	localDB, err := db.NewTempDB()
	if err != nil {
		return nil, err
	}

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "mood",
			Aliases: []string{"m"},
			Value:   "",
			Usage:   "your current mood",
		},
		&cli.StringFlag{
			Name:    "reason",
			Aliases: []string{"r"},
			Value:   "",
			Usage:   "the reason of your mood",
		},
	}

	app := &cli.App{
		Name:                 name,
		Usage:                usage,
		Version:              version,
		EnableBashCompletion: true,
		Suggest:              true,
		Flags:                flags,
		Action:               addUsingFlags(localDB.DB),
	}

	return app, nil
}
