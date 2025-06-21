// Package imageproc provides image processing functionality
// with options for resizing and converting images to grayscale
package imageproc

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/disintegration/imaging"
)

type Options struct {
	ResizePercent int
	Grayscale     bool
	Workers       int
}

var allowedName = regexp.MustCompile(`^v\d+_((\d{3})|(\d+\.\d+))_.+\.(jpg|jpeg|png|webp)$`)

func ProcessImages(rootDir string, opts Options) error {
	supported := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	var files []string

	// Сканируем подходящие файлы
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		base := filepath.Base(path)

		if supported[ext] && allowedName.MatchString(base) {
			files = append(files, path)
		}
		return nil
	})

	total := len(files)
	if total == 0 {
		fmt.Println("📭 No matching images found.")
		return nil
	}

	fmt.Printf("🧵 Starting %d workers to process %d images\n", opts.Workers, total)

	var wg sync.WaitGroup
	fileChan := make(chan string)
	counter := make(chan string)

	// Прогрессбар — БЕЗ go!
	progressDone := make(chan struct{})
	go func() {
		processed := 0
		for filename := range counter {
			processed++
			fmt.Printf("🖼  [%3d%%] %s\n", processed*100/total, filepath.Base(filename))
		}
		close(progressDone)
	}()

	// Workers
	for i := 0; i < opts.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range fileChan {
				img, err := imaging.Open(path)
				if err != nil {
					fmt.Printf("⚠️  Skipping broken image %s: %v\n", path, err)
					continue
				}
				if opts.ResizePercent > 0 && opts.ResizePercent < 100 {
					w := img.Bounds().Dx() * opts.ResizePercent / 100
					h := img.Bounds().Dy() * opts.ResizePercent / 100
					img = imaging.Resize(img, w, h, imaging.Lanczos)
				}
				if opts.Grayscale {
					img = imaging.Grayscale(img)
				}
				if err := imaging.Save(img, path); err != nil {
					fmt.Printf("⚠️  Failed to save image %s: %v\n", path, err)
				}
				counter <- path
			}
		}()
	}

	// Отправка заданий
	go func() {
		for _, path := range files {
			fileChan <- path
		}
		close(fileChan)
	}()

	wg.Wait()
	close(counter)
	<-progressDone // ждём пока выведутся все строки

	fmt.Println("✅ Image processing complete.")
	return nil
}
