package main

import (
	"log"
	"os"
	"strings"

	"github.com/dunstack/dorm/mig"
	"github.com/urfave/cli/v2"
)

func main() {
	migrator := &mig.Migrator{}
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "sql",
						Usage: "generate sql migration",
					},
				},
				Usage: "generate migration",
				Action: func(ctx *cli.Context) error {
					name := strings.Join(ctx.Args().Slice(), "_")
					if ctx.Bool("sql") {
						if err := migrator.GenerateSQLMigration(name); err != nil {
							return err
						}
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
