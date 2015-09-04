package commands

import (
	"github.com/spf13/cobra"

	"github.com/authit/crypto/gen"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "gen",
		Short: "Generate a key pair",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			gen.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
