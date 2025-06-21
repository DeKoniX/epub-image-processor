package epub

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CreateEPUB создает EPUB из распакованной директории
func CreateEPUB(sourceDir, outPath string) error {
	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("cannot create epub file: %w", err)
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	// Проходим по всем файлам
	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relPath := strings.TrimPrefix(path, sourceDir)
		relPath = strings.TrimPrefix(relPath, string(os.PathSeparator)) // на случай ведущего `/`

		// mimetype должен быть без сжатия и первым
		var header *zip.FileHeader
		if relPath == "mimetype" {
			header, err = zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			header.Method = zip.Store // без сжатия
		} else {
			header, err = zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			header.Method = zip.Deflate
		}
		header.Name = relPath

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		inFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer inFile.Close()

		_, err = io.Copy(writer, inFile)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
