package cmd

import (
	"fmt"
	"strings"
	"repo-card/internal/github"
	"repo-card/internal/render"
	"github.com/spf13/cobra"
)

var outputFlag string
var pngFlag bool

func init() {
	generateCmd.Flags().StringVarP(&outputFlag, "output", "o", "", "Nama file output HTML")
	generateCmd.Flags().BoolVarP(&pngFlag, "png", "p", false, "Simpan kartu sebagai gambar PNG")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "gen [owner/repo]",
	Short: "Generate kartu HTML untuk sebuah repositori",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parts := strings.Split(args[0], "/")
		if len(parts) != 2 {
			fmt.Println("Gunakan format: owner/repo")
			return
		}

		fmt.Println("🚀 Memproses data GitHub...")
		repo, err := github.FetchRepo(parts[0], parts[1])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		readme := github.FetchReadme(parts[0], parts[1], repo.DefaultBranch)
		
		fmt.Println("🎨 Merender kartu...")
		path, err := render.ToFile(*repo, readme, outputFlag)
		if err != nil {
			fmt.Printf("Gagal menyimpan file HTML: %v\n", err)
			return
		}
		fmt.Printf("✨ HTML Berhasil: %s\n", path)

		if pngFlag {
			fmt.Println("📸 Mengonversi ke PNG (membutuhkan Chrome/Chromium)...")
			pngPath, err := render.ToPNG(path)
			if err != nil {
				fmt.Printf("❌ Gagal generate PNG: %v\n", err)
				return
			}
			fmt.Printf("🖼️ PNG Berhasil: %s\n", pngPath)
		}
	},
}
