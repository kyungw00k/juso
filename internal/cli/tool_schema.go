package cli

import (
	"encoding/json"
	"fmt"

	"github.com/kyungw00k/kozip/internal/i18n"
	"github.com/spf13/cobra"
)

type toolParam struct {
	Type        string   `json:"type"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Default     any      `json:"default,omitempty"`
}

type toolSchema struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Parameters  *toolSchemaParameters `json:"parameters,omitempty"`
}

type toolSchemaParameters struct {
	Type       string               `json:"type"`
	Properties map[string]toolParam `json:"properties"`
	Required   []string             `json:"required,omitempty"`
}

var schemas = []toolSchema{
	{
		Name:        "kozip_search",
		Description: "Search Korean postal codes and addresses by keyword. Returns postcode, Korean/English road and jibun addresses, building name, and map URLs.",
		Parameters: &toolSchemaParameters{
			Type: "object",
			Properties: map[string]toolParam{
				"keyword": {Type: "string", Description: "Search keyword (Korean or English address, building name, postal code)"},
				"lang":    {Type: "string", Description: "Address language for display", Enum: []string{"ko", "en", "all"}, Default: "ko"},
				"jibun":   {Type: "boolean", Description: "Show jibun (lot-number) address instead of road address", Default: false},
				"output":  {Type: "string", Description: "Output format", Enum: []string{"json", "jsonl", "csv", "table"}, Default: "json"},
			},
			Required: []string{"keyword"},
		},
	},
	{
		Name:        "kozip_cache_clear",
		Description: "Clear the local address search cache.",
	},
	{
		Name:        "kozip_cache_stats",
		Description: "Show cache statistics: number of cached entries and database file size.",
	},
}

var toolSchemaCmd = &cobra.Command{
	Use:     "tool-schema [command]",
	Short:   i18n.T(i18n.MsgToolSchemaShort),
	GroupID: "util",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var out interface{}

		if len(args) == 1 {
			name := "kozip_" + args[0]
			for _, s := range schemas {
				if s.Name == name {
					out = s
					break
				}
			}
			if out == nil {
				return fmt.Errorf("unknown command: %s", args[0])
			}
		} else {
			out = schemas
		}

		b, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(toolSchemaCmd)
}
