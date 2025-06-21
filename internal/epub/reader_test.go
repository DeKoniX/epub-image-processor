package epub

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUnzipEPUB(t *testing.T) {
	src := filepath.Join("testdata", "test.epub")
	dest := t.TempDir()

	if _, err := os.Stat(src); os.IsNotExist(err) {
		t.Skip("No test EPUB found at testdata/test.epub — skipping test")
	}

	if err := UnzipEPUB(src, dest); err != nil {
		t.Fatalf("Unzip failed: %v", err)
	}

	// Проверим, что распаковался хотя бы mimetype
	mimetypePath := filepath.Join(dest, "mimetype")
	if _, err := os.Stat(mimetypePath); err != nil {
		t.Errorf("mimetype not found: %v", err)
	}
}
