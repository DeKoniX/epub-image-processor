package epub

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUnzipEPUB(t *testing.T) {
	src := filepath.Join("../../assets/test.epub")
	dest := "../../tmp/test_unpack"

	err := os.RemoveAll(dest)
	if err != nil {
		t.Fatalf("cleanup failed: %v", err)
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
