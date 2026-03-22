package render

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
)

// ToPNG mengambil screenshot dari file HTML yang sudah di-generate
func ToPNG(htmlPath string) (string, error) {
	// Buat path absolut untuk file HTML agar bisa dibuka browser
	absPath, err := filepath.Abs(htmlPath)
	if err != nil {
		return "", fmt.Errorf("gagal mendapatkan path absolut: %w", err)
	}
	fileURL := "file://" + absPath
	pngPath := htmlPath[:len(htmlPath)-len(filepath.Ext(htmlPath))] + ".png"

	// Setup context chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Timeout 30 detik untuk rendering
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var buf []byte
	err = chromedp.Run(ctx,
		chromedp.Navigate(fileURL),
		// Tunggu font dan animasi awal selesai (0.5s)
		chromedp.Sleep(500*time.Millisecond),
		// Ambil screenshot pada elemen .card
		chromedp.Screenshot(".card", &buf, chromedp.NodeVisible),
	)
	if err != nil {
		return "", fmt.Errorf("gagal merender PNG: %w (Pastikan Chrome/Chromium terinstal)", err)
	}

	if err := os.WriteFile(pngPath, buf, 0644); err != nil {
		return "", fmt.Errorf("gagal menulis file PNG: %w", err)
	}

	return pngPath, nil
}
