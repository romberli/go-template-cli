package message

import (
	"github.com/romberli/go-util/config"
)

const (
	ErrPrintHelpInfo              = 400001
	ErrEmptyLogFileName           = 400002
	ErrNotValidLogFileName        = 400003
	ErrNotValidLogLevel           = 400004
	ErrNotValidLogFormat          = 400005
	ErrNotValidLogMaxSize         = 400006
	ErrNotValidLogMaxDays         = 400007
	ErrNotValidLogMaxBackups      = 400008
	ErrNotValidServerPort         = 400009
	ErrNotValidServerAddr         = 400010
	ErrNotValidPidFile            = 400011
	ErrNotValidServerReadTimeout  = 400012
	ErrNotValidServerWriteTimeout = 400013
	ErrValidateConfig             = 400014
	ErrInitDefaultConfig          = 400015
	ErrReadConfigFile             = 400016
	ErrOverrideCommandLineArgs    = 400017
	ErrAbsoluteLogFilePath        = 400018
	ErrInitLogger                 = 400019
	ErrBaseDir                    = 400020
	ErrInitConfig                 = 400021
	ErrCheckServerPid             = 400022
	ErrCheckServerRunningStatus   = 400023
	ErrServerIsRunning            = 400024
	ErrStartAsForeground          = 400025
	ErrSavePidToFile              = 400026
	ErrKillServerWithPid          = 400027
	ErrKillServerWithPidFile      = 400028
	ErrGetPidFromPidFile          = 400029
	ErrSetSid                     = 400030
	ErrRemovePidFile              = 400031
)

func initErrorMessage() {
	Messages[ErrPrintHelpInfo] = config.NewErrMessage(DefaultMessageHeader, ErrPrintHelpInfo, "got message when printing help information")
	Messages[ErrEmptyLogFileName] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyLogFileName, "log file name could not be an empty string")
	Messages[ErrNotValidLogFileName] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFileName, "log file name must be either unix or windows path format, %s is not valid")
	Messages[ErrNotValidLogLevel] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogLevel, "log level must be one of [debug, info, warn, message, fatal], %s is not valid")
	Messages[ErrNotValidLogFormat] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFormat, "log level must be either text or json, %s is not valid")
	Messages[ErrNotValidLogMaxSize] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxSize, "log max size must be between %d and %d, %d is not valid")
	Messages[ErrNotValidLogMaxDays] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxDays, "log max days must be between %d and %d, %d is not valid")
	Messages[ErrNotValidLogMaxBackups] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxBackups, "log max backups must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerPort] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerPort, "Server port must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerAddr] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerAddr, "server addr must be formatted as host:port, %s is not valid")
	Messages[ErrNotValidPidFile] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidPidFile, "pid file name must be either unix or windows path format, %s is not valid")
	Messages[ErrNotValidServerReadTimeout] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerReadTimeout, "server read timeout must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerWriteTimeout] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerWriteTimeout, "server write timeout must be between %d and %d, %d is not valid")
	Messages[ErrValidateConfig] = config.NewErrMessage(DefaultMessageHeader, ErrValidateConfig, "validate config failed")
	Messages[ErrInitDefaultConfig] = config.NewErrMessage(DefaultMessageHeader, ErrInitDefaultConfig, "init default configuration failed")
	Messages[ErrReadConfigFile] = config.NewErrMessage(DefaultMessageHeader, ErrReadConfigFile, "read config file failed")
	Messages[ErrOverrideCommandLineArgs] = config.NewErrMessage(DefaultMessageHeader, ErrOverrideCommandLineArgs, "override command line arguments failed")
	Messages[ErrAbsoluteLogFilePath] = config.NewErrMessage(DefaultMessageHeader, ErrAbsoluteLogFilePath, "get absolute path of log file failed. log file: %s")
	Messages[ErrInitLogger] = config.NewErrMessage(DefaultMessageHeader, ErrInitLogger, "initialize logger failed")
	Messages[ErrBaseDir] = config.NewErrMessage(DefaultMessageHeader, ErrBaseDir, "get base dir of %s failed")
	Messages[ErrInitConfig] = config.NewErrMessage(DefaultMessageHeader, ErrInitConfig, "init config failed")
	Messages[ErrCheckServerPid] = config.NewErrMessage(DefaultMessageHeader, ErrCheckServerPid, "check server pid file failed")
	Messages[ErrCheckServerRunningStatus] = config.NewErrMessage(DefaultMessageHeader, ErrCheckServerRunningStatus, "check server running status failed")
	Messages[ErrServerIsRunning] = config.NewErrMessage(DefaultMessageHeader, ErrServerIsRunning, "pid file exists or server is still running. pid file: %s")
	Messages[ErrStartAsForeground] = config.NewErrMessage(DefaultMessageHeader, ErrStartAsForeground, "got message when starting das as foreground")
	Messages[ErrSavePidToFile] = config.NewErrMessage(DefaultMessageHeader, ErrSavePidToFile, "got message when writing pid to file")
	Messages[ErrKillServerWithPid] = config.NewErrMessage(DefaultMessageHeader, ErrKillServerWithPid, "kill server failed. pid: %d")
	Messages[ErrKillServerWithPidFile] = config.NewErrMessage(DefaultMessageHeader, ErrKillServerWithPidFile, "kill server with pid file failed. pid: %d, pid file: %s")
	Messages[ErrGetPidFromPidFile] = config.NewErrMessage(DefaultMessageHeader, ErrGetPidFromPidFile, "get pid from pid file failed. pid file: %s")
	Messages[ErrSetSid] = config.NewErrMessage(DefaultMessageHeader, ErrSetSid, "set sid failed when daemonizing server")
	Messages[ErrRemovePidFile] = config.NewErrMessage(DefaultMessageHeader, ErrRemovePidFile, "remove pid file failed. pid file: %s")
}
