package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode"

	kozip "github.com/kyungw00k/kozip"
	"github.com/kyungw00k/kozip/api"
	"github.com/kyungw00k/kozip/cache"
	"github.com/kyungw00k/kozip/internal/i18n"
	"github.com/kyungw00k/kozip/internal/output"
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
	Short:   i18n.T(i18n.MsgRootShort),
	Long:    i18n.T(i18n.MsgRootLong),
	Version: Version,
	Args:    cobra.ArbitraryArgs,
	Example: `  kozip 강남역
  kozip gangnam --lang en
  kozip 역삼동 --jibun
  kozip 강남역 --lang all
  kozip 강남역 -o json
  kozip 강남역 -o csv > results.csv`,
	RunE: runSearch,
}

func init() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	rootCmd.AddGroup(
		&cobra.Group{ID: "cache", Title: i18n.T(i18n.GroupCache)},
		&cobra.Group{ID: "util", Title: i18n.T(i18n.GroupUtil)},
	)

	pf := rootCmd.PersistentFlags()
	pf.StringVarP(&flagOutput, "output", "o", "auto", i18n.T(i18n.FlagOutputUsage))
	pf.StringVar(&flagLang, "lang", "ko", i18n.T(i18n.FlagLangUsage))
	pf.BoolVar(&flagJibun, "jibun", false, i18n.T(i18n.FlagJibunUsage))

	rootCmd.SuggestionsMinimumDistance = 2
}

func Execute() error {
	return rootCmd.Execute()
}

func runSearch(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Help()
	}

	keyword := strings.Join(args, " ")

	// --lang 미지정 시 키워드 언어 자동 감지
	if !cmd.Flags().Changed("lang") && isASCII(keyword) {
		flagLang = "en"
	}

	c, err := cache.Open()
	if err != nil {
		return err
	}
	defer c.Close()

	var results []kozip.AddressResult

	if data, ok := c.Get(keyword); ok {
		if err := json.Unmarshal(data, &results); err != nil {
			results = nil
		}
	}

	if results == nil {
		client := api.NewClient()
		results, err = client.Search(cmd.Context(), keyword)
		if err != nil {
			return err
		}

		if err := c.Set(keyword, results); err != nil {
			fmt.Fprintln(os.Stderr, "warning: failed to cache results:", err)
		}
	}

	if len(results) == 0 {
		fmt.Fprintln(os.Stderr, i18n.T(i18n.ErrNoResults))
		return nil
	}

	columns := getColumns()
	formatter := output.NewFormatter(flagOutput, columns)
	return formatter.Format(prepareOutput(results))
}

func getColumns() []output.TableColumn {
	switch flagLang {
	case "en":
		return []output.TableColumn{
			{Header: i18n.T(i18n.HdrPostcode), Key: "postcode5"},
			{Header: i18n.T(i18n.HdrAddress), Key: "en_address"},
		}
	case "all":
		return []output.TableColumn{
			{Header: i18n.T(i18n.HdrPostcode), Key: "postcode5"},
			{Header: i18n.T(i18n.HdrKoAddress), Key: "ko_address"},
			{Header: i18n.T(i18n.HdrEnAddress), Key: "en_address"},
		}
	default:
		return []output.TableColumn{
			{Header: i18n.T(i18n.HdrPostcode), Key: "postcode5"},
			{Header: i18n.T(i18n.HdrAddress), Key: "ko_address"},
		}
	}
}

func isASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func prepareOutput(results []kozip.AddressResult) []kozip.AddressResult {
	if !flagJibun {
		return results
	}
	out := make([]kozip.AddressResult, len(results))
	for i, r := range results {
		out[i] = r
		out[i].KoAddress = r.KoJibun
		out[i].EnAddress = r.EnJibun
	}
	return out
}
