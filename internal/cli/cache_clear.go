package cli

import (
	"fmt"

	"github.com/kyungw00k/juso/cache"
	"github.com/kyungw00k/juso/internal/i18n"
	"github.com/spf13/cobra"
)

var cacheClearCmd = &cobra.Command{
	Use:   "clear",
	Short: i18n.T(i18n.MsgCacheClear),
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cache.Open()
		if err != nil {
			return err
		}
		defer c.Close()

		if err := c.Clear(); err != nil {
			return err
		}

		fmt.Println(i18n.T(i18n.MsgCacheCleared))
		return nil
	},
}

func init() {
	cacheCmd.AddCommand(cacheClearCmd)
}
