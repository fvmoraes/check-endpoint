package cli

import (
	"check-endpoint/controllers"
	"check-endpoint/logs"
	"os"

	"github.com/urfave/cli/v2"
)

func CliApp() {
	app := &cli.App{
		Name:  "check-endpoint",
		Usage: "./check-endpoint(.exe)",
		Commands: []*cli.Command{
			{
				Name: "run",
				Usage: `
				--file (-f) <File containing list of URLs> 
				--repeat (-r) <Number of repetitions> 
				--interval (-i) <Time interval between requests (in milliseconds)>
				`,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Required: true,
						Usage:    "File containing list of URLs. Keep the list containing valid URLs and no additional line breaks",
					},
					&cli.IntFlag{
						Name:     "repeat",
						Aliases:  []string{"r"},
						Required: true,
						Usage:    "Number of repetitions. Use integer values",
					},
					&cli.IntFlag{
						Name:     "interval",
						Aliases:  []string{"i"},
						Required: true,
						Usage:    "Time interval between requests (in milliseconds). Use values greater than 200ms for better data collection",
					},
				},
				Action: SendData,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logs.LogApplicationError(err)
	}
}

func SendData(c *cli.Context) error {
	command := os.Args[1]
	file := c.String("file")
	repeat := c.Int("repeat")
	interval := c.Int("interval")
	logs.LogApplicationStartInformation(command, file, repeat, interval)
	controllers.GetInputArgs(file, repeat, interval)
	return nil
}
