package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/authit/crypto/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "crypto"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
