package main

import (
	"fmt"
	"github.com/tdytaylor/jvm.go/cmd"
)

func main() {
	parseCmd := cmd.ParseCmd()
	if parseCmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if parseCmd.HelpFlag || parseCmd.Class == "" {
		cmd.PrintUsage()
	} else {
		startJVM(parseCmd)
	}
}

func startJVM(cmd *cmd.Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.CpOption, cmd.Class, cmd.Args)
}
