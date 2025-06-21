package epub

import (
	"os"
	"testing"
)

func TestCreateEPUB(t *testing.T) {
	src := "../../tmp/test_unpack"
	out := "../../output/test_output.epub"

	if _, err := os.Stat(src); os.IsNotExist(err) {
		t.Skip("test_unpack not found — run TestUnzipEPUB first")
	}

	err := CreateEPUB(src, out)
	if err != nil {
		t.Fatalf("CreateEPUB failed: %v", err)
	}

	// Проверим, что файл создан
	if _, err := os.Stat(out); err != nil {
		t.Errorf("EPUB file not created: %v", err)
	}
}
