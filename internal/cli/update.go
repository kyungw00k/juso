package cli

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/kyungw00k/kozip/internal/i18n"
	"github.com/spf13/cobra"
)

var flagCheckOnly bool

type ghRelease struct {
	TagName string    `json:"tag_name"`
	Assets  []ghAsset `json:"assets"`
}

type ghAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   i18n.T(i18n.MsgUpdateShort),
	GroupID: "util",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
		defer cancel()

		rel, err := fetchLatestRelease(ctx)
		if err != nil {
			return err
		}

		latest := strings.TrimPrefix(rel.TagName, "v")
		current := strings.TrimPrefix(Version, "v")

		if latest == current {
			fmt.Printf("Already up to date (v%s)\n", current)
			return nil
		}

		fmt.Printf("New version available: v%s → v%s\n", current, latest)

		if flagCheckOnly {
			return nil
		}

		assetName := fmt.Sprintf("kozip_%s_%s", runtime.GOOS, runtime.GOARCH)
		var downloadURL string
		for _, a := range rel.Assets {
			if strings.Contains(a.Name, assetName) && strings.HasSuffix(a.Name, ".tar.gz") {
				downloadURL = a.BrowserDownloadURL
				break
			}
		}
		if downloadURL == "" {
			return fmt.Errorf("no asset found for %s/%s", runtime.GOOS, runtime.GOARCH)
		}

		exe, err := os.Executable()
		if err != nil {
			return err
		}
		exe, err = filepath.EvalSymlinks(exe)
		if err != nil {
			return err
		}

		tmpDir, err := os.MkdirTemp("", "kozip-update-*")
		if err != nil {
			return err
		}
		defer os.RemoveAll(tmpDir)

		fmt.Printf("Downloading %s...\n", filepath.Base(downloadURL))
		if err := downloadAndExtract(ctx, downloadURL, tmpDir); err != nil {
			return err
		}

		newBin := filepath.Join(tmpDir, "kozip")
		backup := exe + ".bak"

		os.Rename(exe, backup)
		if err := copyFile(newBin, exe); err != nil {
			os.Rename(backup, exe)
			return err
		}
		os.Remove(backup)

		fmt.Printf("Updated to v%s\n", latest)
		return nil
	},
}

func init() {
	updateCmd.Flags().BoolVar(&flagCheckOnly, "check", false, "Check for updates without installing")
	rootCmd.AddCommand(updateCmd)
}

func fetchLatestRelease(ctx context.Context) (*ghRelease, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://api.github.com/repos/kyungw00k/kozip/releases/latest", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rel ghRelease
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return nil, err
	}
	return &rel, nil
}

func downloadAndExtract(ctx context.Context, url, dir string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	limited := io.LimitReader(resp.Body, 100*1024*1024)

	gz, err := gzip.NewReader(limited)
	if err != nil {
		return err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if hdr.Typeflag != tar.TypeReg {
			continue
		}

		target := filepath.Join(dir, filepath.Base(hdr.Name))
		f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		io.Copy(f, tr)
		f.Close()
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o755)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
