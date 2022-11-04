package g

import "github.com/eoe2005/g/gworker"

func RunDaemon(ws ...gworker.Worker) {
	initConfig()
	gworker.Run(ws)
}
