package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/dekonix/epub-image-processor/internal/epub"
	imageproc "github.com/dekonix/epub-image-processor/internal/image"
)

func main() {
	defaultWorkers := runtime.NumCPU()
	// CLI флаги
	inputEPUB := flag.String("in", "", "Path to input EPUB file")
	outputEPUB := flag.String("out", "", "Optional: custom output path (default: same filename in ./output/)")
	resize := flag.Int("resize", 100, "Resize percentage for images (e.g. 50 means 50%)")
	grayscale := flag.Bool("grayscale", false, "Convert images to grayscale")
	workers := flag.Int("workers", defaultWorkers, "Number of parallel workers to process images")

	flag.Parse()

	if *inputEPUB == "" {
		log.Fatal("You must specify input EPUB with -in")
	}

	outputDir := "tmp/unpacked"

	// Автоматически определить имя выходного файла, если не указано
	outputPath := *outputEPUB
	if outputPath == "" {
		inputBase := filepath.Base(*inputEPUB)          // example: test.epub
		outputPath = filepath.Join("output", inputBase) // example: output/test.epub
	}

	_ = os.RemoveAll(outputDir)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("output", 0755); err != nil {
		log.Fatal(err)
	}

	fmt.Println("📖 Unzipping EPUB...")
	if err := epub.UnzipEPUB(*inputEPUB, outputDir); err != nil {
		log.Fatalf("Unzip failed: %v", err)
	}

	fmt.Println("🖼  Processing images...")
	if err := imageproc.ProcessImages(outputDir, imageproc.Options{
		ResizePercent: *resize,
		Grayscale:     *grayscale,
		Workers:       *workers,
	}); err != nil {
		log.Fatalf("Image processing failed: %v", err)
	}

	fmt.Println("📦 Repacking EPUB...")
	if err := epub.CreateEPUB(outputDir, outputPath); err != nil {
		log.Fatalf("EPUB creation failed: %v", err)
	}

	fmt.Println("✅ Done. Output EPUB saved to:", outputPath)
}
