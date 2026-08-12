package commands

import (
	"github.com/CloudPloyHQ/cli/src/common"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ploy cli",
	Run: func(cmd *cobra.Command, args []string) {
		println(common.CurrentCliVersion)
	},
}
