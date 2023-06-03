package main

import (
	"check-endpoint/cli"
	"check-endpoint/controllers"
)

func main() {
	//	application := cli.CliApp()
	//	if err := application.Run(os.Args); err != nil {
	//		logs.LogApplicationError(err)
	//	} else {
	//		logs.LogApplicationStartInformation()
	//	}
	cli.CliApp()
	controllers.CalculateStatistics()

}
