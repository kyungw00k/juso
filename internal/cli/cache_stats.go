package cli

import (
	"fmt"

	"github.com/kyungw00k/juso/cache"
	"github.com/kyungw00k/juso/internal/i18n"
	"github.com/spf13/cobra"
)

var cacheStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: i18n.T(i18n.MsgCacheStats),
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cache.Open()
		if err != nil {
			return err
		}
		defer c.Close()

		stats, err := c.Stats()
		if err != nil {
			return err
		}

		fmt.Println(i18n.Tf(i18n.MsgCacheEntries, stats.Entries))
		fmt.Println(i18n.Tf(i18n.MsgCacheSize, cache.FormatSize(stats.Size)))
		return nil
	},
}

func init() {
	cacheCmd.AddCommand(cacheStatsCmd)
}
