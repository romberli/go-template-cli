/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pingcap/errors"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/romberli/go-template-cli/config"
	"github.com/romberli/go-template-cli/pkg/message"
)

const defaultConfigFileType = "yaml"

var (
	// config
	baseDir string
	cfgFile string
	// daemon
	daemon    bool
	daemonStr string
	// log
	logFileName           string
	logLevel              string
	logFormat             string
	logMaxSize            int
	logMaxDays            int
	logMaxBackups         int
	logRotateOnStartupStr string
	logStdoutStr          string
	// server
	serverAddr         string
	serverPid          int
	serverPidFile      string
	serverReadTimeout  int
	serverWriteTimeout int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-template-cli",
	Short: "go-template-cli",
	Long:  `go-template-cli is a template of golang command line program`,
	Run: func(cmd *cobra.Command, args []string) {
		// if no subcommand is set, it will print help information.
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				fmt.Println(fmt.Sprintf(constant.LogWithStackString, message.NewMessage(message.ErrPrintHelpInfo, errors.Trace(err))))
				os.Exit(constant.DefaultAbnormalExitCode)
			}

			os.Exit(constant.DefaultNormalExitCode)
		}

		// init config
		err := initConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf(constant.LogWithStackString, message.NewMessage(message.ErrInitConfig, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(fmt.Sprintf(constant.LogWithStackString, errors.Trace(err)))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
}

func init() {
	// set usage template
	rootCmd.SetUsageTemplate(UsageTemplateWithoutDefault())

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// config
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", constant.DefaultRandomString, "config file path")
	// daemon
	rootCmd.PersistentFlags().StringVar(&daemonStr, "daemon", constant.DefaultRandomString, fmt.Sprintf("whether run in background as a daemon(default: %s)", constant.FalseString))
	// log
	rootCmd.PersistentFlags().StringVar(&logFileName, "log-file", constant.DefaultRandomString, fmt.Sprintf("specify the log file name(default: %s)", filepath.Join(config.DefaultLogDir, log.DefaultLogFileName)))
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", constant.DefaultRandomString, fmt.Sprintf("specify the log level(default: %s)", log.DefaultLogLevel))
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", constant.DefaultRandomString, fmt.Sprintf("specify the log format(default: %s)", log.DefaultLogFormat))
	rootCmd.PersistentFlags().IntVar(&logMaxSize, "log-max-size", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max size(default: %d)", log.DefaultLogMaxSize))
	rootCmd.PersistentFlags().IntVar(&logMaxDays, "log-max-days", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max days(default: %d)", log.DefaultLogMaxDays))
	rootCmd.PersistentFlags().IntVar(&logMaxBackups, "log-max-backups", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max backups(default: %d)", log.DefaultLogMaxBackups))
	rootCmd.PersistentFlags().StringVar(&logRotateOnStartupStr, "log-rotate-on-startup", constant.DefaultRandomString, fmt.Sprintf("specify if rotating the log file on startup(default: %s)", constant.FalseString))
	rootCmd.PersistentFlags().StringVar(&logStdoutStr, "log-stdout", constant.DefaultRandomString, fmt.Sprintf("specify the log the message to stdout as well(default: %s)", constant.FalseString))
	// server
	rootCmd.PersistentFlags().StringVar(&serverAddr, "server-addr", constant.DefaultRandomString, fmt.Sprintf("specify the server addr(default: %s)", config.DefaultServerAddr))
	rootCmd.PersistentFlags().StringVar(&serverPidFile, "server-pid-file", constant.DefaultRandomString, fmt.Sprintf("specify the server pid file path(default: %s)", filepath.Join(config.DefaultBaseDir, fmt.Sprintf("%s.pid", config.DefaultCommandName))))
	rootCmd.PersistentFlags().IntVar(&serverReadTimeout, "server-read-timeout", constant.DefaultRandomInt, fmt.Sprintf("specify the read timeout in seconds of http request(default: %d)", config.DefaultServerReadTimeout))
	rootCmd.PersistentFlags().IntVar(&serverWriteTimeout, "server-write-timeout", constant.DefaultRandomInt, fmt.Sprintf("specify the write timeout in seconds of http request(default: %d)", config.DefaultServerWriteTimeout))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() error {
	var err error

	// init default config
	err = initDefaultConfig()
	if err != nil {
		return message.NewMessage(message.ErrInitDefaultConfig, err.Error())
	}

	// read config with config file
	err = ReadConfigFile()
	if err != nil {
		return message.NewMessage(message.ErrInitDefaultConfig, err)
	}

	// override config with command line arguments
	err = OverrideConfigByCLI()
	if err != nil {
		return message.NewMessage(message.ErrOverrideCommandLineArgs, err)
	}

	// init log
	fileName := viper.GetString(config.LogFileNameKey)
	level := viper.GetString(config.LogLevelKey)
	format := viper.GetString(config.LogFormatKey)
	maxSize := viper.GetInt(config.LogMaxSizeKey)
	maxDays := viper.GetInt(config.LogMaxDaysKey)
	maxBackups := viper.GetInt(config.LogMaxBackupsKey)

	fileNameAbs := fileName
	isAbs := filepath.IsAbs(fileName)
	if !isAbs {
		fileNameAbs, err = filepath.Abs(fileName)
		if err != nil {
			return message.NewMessage(message.ErrAbsoluteLogFilePath, errors.Trace(err), fileName)
		}
	}
	_, _, err = log.InitFileLogger(fileNameAbs, level, format, maxSize, maxDays, maxBackups)
	if err != nil {
		return message.NewMessage(message.ErrInitLogger, err)
	}

	log.SetDisableDoubleQuotes(true)
	log.SetDisableEscape(true)

	if viper.GetBool(config.LogStdoutKey) {
		log.AddWriteSyncer(log.NewStdoutWriteSyncer())
	}

	if viper.GetBool(config.LogRotateOnStartupKey) {
		err = log.Rotate()
		if err != nil {
			return message.NewMessage(message.ErrRotateLogFile, err)
		}
	}
	return nil
}

// initDefaultConfig initiate default configuration
func initDefaultConfig() (err error) {
	// get base dir
	baseDir, err = filepath.Abs(config.DefaultBaseDir)
	if err != nil {
		return message.NewMessage(message.ErrBaseDir, errors.Trace(err), config.DefaultCommandName)
	}
	// set default config value
	config.SetDefaultConfig(baseDir)
	err = config.ValidateConfig()
	if err != nil {
		return err
	}

	return nil
}

// ReadConfigFile read configuration from config file, it will override the init configuration
func ReadConfigFile() (err error) {
	if cfgFile != constant.EmptyString && cfgFile != constant.DefaultRandomString {
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType(defaultConfigFileType)
		err = viper.ReadInConfig()
		if err != nil {
			return errors.Trace(err)
		}
		err = config.ValidateConfig()
		if err != nil {
			return message.NewMessage(message.ErrValidateConfig, err)
		}
	}

	return nil
}

// UsageTemplateWithoutDefault returns a usage template which does not contain default part
func UsageTemplateWithoutDefault() string {
	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsagesWithoutDefault | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsagesWithoutDefault | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}
