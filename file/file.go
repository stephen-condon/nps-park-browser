package file

import "os"

var filename string

func Read() (string, error) {
	if filename == "" {
		setFileName("./api.txt")
	}
	data, err := os.ReadFile(filename)

	return string(data), err
}

func setFileName(theFile string) {
	filename = theFile
}
