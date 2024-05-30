package option

import (
	"os"

	"github.com/aide-cloud/moon/pkg/types"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		log.Warn("Not Command")
	},
}

func Execute() {
	rootCmd.AddCommand(serverCmd, versionCmd, genCmd)
	if err := rootCmd.Execute(); !types.IsNil(err) {
		log.Error(err)
		os.Exit(1)
	}
}
