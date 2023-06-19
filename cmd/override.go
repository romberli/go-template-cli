package cmd

import (
	"strings"

	"github.com/pingcap/errors"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/romberli/go-template-cli/config"
	"github.com/romberli/go-template-cli/pkg/message"
)

// OverrideConfigByCLI read configuration from command line interface, it will override the config file configuration
func OverrideConfigByCLI() error {
	// override config
	overrideConfigByCLI()
	// override daemon
	err := overrideDaemonByCLI()
	if err != nil {
		return err
	}
	// override log
	err = overrideLogByCLI()
	if err != nil {
		return err
	}
	// override server
	err = overrideServerByCLI()
	if err != nil {
		return err
	}
	// validate configuration
	err = config.ValidateConfig()
	if err != nil {
		return message.NewMessage(message.ErrValidateConfig, err)
	}

	return nil
}

// overrideConfigByCLI overrides the config section by command line interface
func overrideConfigByCLI() {
	if cfgFile != constant.EmptyString && cfgFile != constant.DefaultRandomString {
		viper.Set(config.ConfKey, cfgFile)
	}
}

// overrideDaemonByCLI overrides the daemon section by command line interface
func overrideDaemonByCLI() error {
	if daemonStr != constant.DefaultRandomString {
		daemon, err := cast.ToBoolE(daemonStr)
		if err != nil {
			return errors.Trace(err)
		}

		viper.Set(config.DaemonKey, daemon)
	}

	return nil
}

// overrideLogByCLI overrides the log section by command line interface
func overrideLogByCLI() error {
	if logFileName != constant.DefaultRandomString {
		viper.Set(config.LogFileNameKey, logFileName)
	}
	if logLevel != constant.DefaultRandomString {
		logLevel = strings.ToLower(logLevel)
		viper.Set(config.LogLevelKey, logLevel)
	}
	if logFormat != constant.DefaultRandomString {
		logLevel = strings.ToLower(logFormat)
		viper.Set(config.LogFormatKey, logFormat)
	}
	if logMaxSize != constant.DefaultRandomInt {
		viper.Set(config.LogMaxSizeKey, logMaxSize)
	}
	if logMaxDays != constant.DefaultRandomInt {
		viper.Set(config.LogMaxDaysKey, logMaxDays)
	}
	if logMaxBackups != constant.DefaultRandomInt {
		viper.Set(config.LogMaxBackupsKey, logMaxBackups)
	}
	if logRotateOnStartupStr != constant.DefaultRandomString {
		rotateOnStartup, err := cast.ToBoolE(logRotateOnStartupStr)
		if err != nil {
			return errors.Trace(err)
		}

		viper.Set(config.LogRotateOnStartupKey, rotateOnStartup)
	}

	return nil
}

// overrideServerByCLI overrides the server section by command line interface
func overrideServerByCLI() error {
	if serverAddr != constant.DefaultRandomString {
		viper.Set(config.ServerAddrKey, serverAddr)
	}
	if serverPidFile != constant.DefaultRandomString {
		viper.Set(config.ServerPidFileKey, serverPidFile)
	}
	if serverReadTimeout != constant.DefaultRandomInt {
		viper.Set(config.ServerReadTimeoutKey, serverReadTimeout)
	}
	if serverWriteTimeout != constant.DefaultRandomInt {
		viper.Set(config.ServerWriteTimeoutKey, serverWriteTimeout)
	}

	return nil
}
