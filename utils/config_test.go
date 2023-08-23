package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// var mutex sync.Mutex

func createFileSync(filepath string, content []byte) {
	// mutex.Lock()
	// defer mutex.Unlock()
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	file.Write(content)
	defer file.Close()
}
func Test_readDirModOrder(t *testing.T) {

	path, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Error(err)
	}
	firstElement, err := GenerateRandomID(12)
	createFileSync(filepath.Join(path, firstElement), []byte(""))
	time.Sleep(100 * time.Millisecond)

	// var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// wg.Add(1)
		// go func() {
		// defer wg.Done()
		uuid, err := GenerateRandomID(12)
		if err != nil {
			t.Error(err)
		}
		createFileSync(filepath.Join(path, uuid), []byte(""))

		// }()
		time.Sleep(100 * time.Millisecond)

	}
	// wg.Wait()

	fileElements, err := GetSortedFileInfos(path)
	if fileElements[0].Name() != firstElement {
		t.Error("\nExpected: ", firstElement, "Got: ", fileElements[0].Name(), "\n", fileElements[len(fileElements)-1].Name())
	}
}
