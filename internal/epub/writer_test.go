package epub

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateEPUB(t *testing.T) {
	src := t.TempDir()

	// Создаём структуру для EPUB в src
	mimetypePath := filepath.Join(src, "mimetype")
	if err := os.WriteFile(mimetypePath, []byte("application/epub+zip"), 0o644); err != nil {
		t.Fatalf("Failed to write mimetype: %v", err)
	}

	metaInf := filepath.Join(src, "META-INF")
	if err := os.MkdirAll(metaInf, 0o755); err != nil {
		t.Fatalf("Failed to create META-INF: %v", err)
	}

	containerXML := `<?xml version="1.0" encoding="UTF-8"?>
<container version="1.0" xmlns="urn:oasis:names:tc:opendocument:xmlns:container">
<rootfiles><rootfile full-path="content.opf" media-type="application/oebps-package+xml"/></rootfiles>
</container>`
	containerPath := filepath.Join(metaInf, "container.xml")
	if err := os.WriteFile(containerPath, []byte(containerXML), 0o644); err != nil {
		t.Fatalf("Failed to write container.xml: %v", err)
	}

	// Выходной файл
	out := filepath.Join(t.TempDir(), "test.epub")

	if err := CreateEPUB(src, out); err != nil {
		t.Fatalf("CreateEPUB failed: %v", err)
	}

	if _, err := os.Stat(out); err != nil {
		t.Errorf("EPUB not created: %v", err)
	}
}
