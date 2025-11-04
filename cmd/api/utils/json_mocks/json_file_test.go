package json_mocks_test

import (
	"taska-core-me-go/cmd/api/utils/json_mocks"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetJSONFile_Success(t *testing.T) {
	_, currentFile, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(currentFile)

	tmpFolder := filepath.Join(testDir, "testdata")
	_ = os.MkdirAll(tmpFolder, 0755)

	filePath := filepath.Join(tmpFolder, "sample.json")
	content := `{"hello":"taska"}`
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create test json file: %v", err)
	}

	data := json_mocks.GetJSONFile("testdata", "sample.json")

	if string(data) != content {
		t.Errorf("expected %q, got %q", content, string(data))
	}
}

func TestGetJSONFile_FileNotFound(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for missing file, but did not panic")
		}
	}()

	_ = json_mocks.GetJSONFile("testdata", "no_such_file.json")
}
