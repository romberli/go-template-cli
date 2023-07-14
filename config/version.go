package config

import (
	"fmt"

	"github.com/romberli/go-util/constant"
)

var (
	AppName    string
	Version    string
	BuildTime  string
	FullCommit string
	Branch     string
)

func ShortInfo() string {
	if AppName == constant.EmptyString {
		return constant.EmptyString
	}

	return fmt.Sprintf("%s-%s", AppName, Version)
}

func FullInfo() string {
	buildInfo := fmt.Sprintf("AppName: %s\nVersion: %s\nBuildTime: %s\nFullCommit: %s", AppName, Version, BuildTime, FullCommit)

	if Branch != constant.EmptyString {
		buildInfo += fmt.Sprintf("\nBranch: %s", Branch)
	}

	return buildInfo
}
