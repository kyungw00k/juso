package cli

import (
	"github.com/spf13/cobra"
)

var Version = "dev"

var (
	flagOutput string
	flagLang   string
	flagJibun  bool
)

var rootCmd = &cobra.Command{
	Use:     "kozip <keyword>",
	Short:   "Korean postal code lookup CLI",
	Version: Version,
	Args:    cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		// search logic will be added later
		return nil
	},
}

func init() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	pf := rootCmd.PersistentFlags()
	pf.StringVarP(&flagOutput, "output", "o", "auto", "Output format: auto, table, json, jsonl, csv")
	pf.StringVar(&flagLang, "lang", "ko", "Address language: ko, en, all")
	pf.BoolVar(&flagJibun, "jibun", false, "Show jibun address instead of road address")

	rootCmd.SuggestionsMinimumDistance = 2
}

func Execute() error {
	return rootCmd.Execute()
}
