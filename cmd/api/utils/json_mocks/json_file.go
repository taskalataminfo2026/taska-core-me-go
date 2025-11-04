package json_mocks

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetJSONFile(folder, fileName string) []byte {
	_, currentFile, _, _ := runtime.Caller(0)
	testUtilsDir := filepath.Dir(currentFile)
	json, err := os.ReadFile(filepath.Join(testUtilsDir, folder, fileName))
	if err != nil {
		panic(err)
	}
	return json
}
