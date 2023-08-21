package functions

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"unsafe"
)

func Watch() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// fileDescriptor := int(file.Fd())

	// Create an inotify instance
	inotifyFd, err := syscall.InotifyInit()
	if err != nil {
		fmt.Println("Error initializing inotify:", err)
		return
	}
	defer syscall.Close(inotifyFd)

	// Add the file to the watch list
	_, err = syscall.InotifyAddWatch(inotifyFd, filePath, syscall.IN_MODIFY)
	if err != nil {
		fmt.Println("Error adding watch:", err)
		return
	}

	fmt.Printf("Listening for changes on %s...\n", filePath)

	// Continuously monitor for events
	buf := make([]byte, syscall.SizeofInotifyEvent*10)
	for {
		n, err := syscall.Read(inotifyFd, buf)
		if err != nil {
			fmt.Println("Error reading events:", err)
			return
		}

		var offset uint32
		for offset < uint32(n) {
			event := (*syscall.InotifyEvent)(unsafe.Pointer(&buf[offset]))
			if event.Wd == 1 && event.Mask&syscall.IN_MODIFY > 0 {
				content, err := readContent(filePath)
				if err != nil {
					//handle
				}
				fmt.Println("File modified:", string(content))
			}
			offset += syscall.SizeofInotifyEvent + uint32(event.Len)
		}
	}
}

func readContent(filePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// func difference(prevFileContent []byte, fileContent []byte) map[int]byte {
// 	differenceMap := make(map[int]byte)
// 	a := len(prevFileContent)
// 	b := len(fileContent)
// 	if a-b > 0 {
// 		i, j := 0, 0
// 		for i < b {
// 			if prevFileContent[i] == fileContent[i] {
// 				i++
// 				continue
// 			}
// 			differenceMap[i] = fileContent[i]
// 		}
// 	}
// 	for fileContent
// 	return differenceMap
// }

// func deletion(b []byte) map(int, byte) {

// }

// func addition(b []byte) map(int, byte) {

// }
