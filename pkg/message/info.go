package message

import (
	"github.com/romberli/go-util/config"
)

const (
	// server
	InfoServerStart      = 200001
	InfoServerStop       = 200002
	InfoServerIsRunning  = 200003
	InfoServerNotRunning = 200004
)

func initInfoMessage() {
	// server
	Messages[InfoServerStart] = config.NewErrMessage(DefaultMessageHeader, InfoServerStart, "go-template-web started successfully. pid: %d, pid file: %s")
	Messages[InfoServerStop] = config.NewErrMessage(DefaultMessageHeader, InfoServerStop, "go-template-web stopped successfully. pid: %d, pid file: %s")
	Messages[InfoServerIsRunning] = config.NewErrMessage(DefaultMessageHeader, InfoServerIsRunning, "go-template-web is running. pid: %d")
	Messages[InfoServerNotRunning] = config.NewErrMessage(DefaultMessageHeader, InfoServerNotRunning, "go-template-web is not running. pid: %d")
}
