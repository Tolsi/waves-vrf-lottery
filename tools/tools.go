package tools

import "os"

func PrintErrorAndExit(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
