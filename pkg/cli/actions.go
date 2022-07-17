package cli

import (
	"errors"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"meeble/pkg/models"
)

func addUsingFlags(db *gorm.DB) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if ctx.String("mood") == "" {
			return errors.New("you have to specify a mood")
		}

		m := &models.Mood{Mood: ctx.String("mood"), Reason: ctx.String("reason")}
		if err := models.CreateMood(db, m); err != nil {
			return err
		}

		return nil
	}
}
