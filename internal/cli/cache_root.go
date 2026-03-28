package cli

import (
	"github.com/kyungw00k/juso/internal/i18n"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:     "cache",
	Short:   i18n.T(i18n.MsgCacheShort),
	GroupID: "cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(cacheCmd)
}
