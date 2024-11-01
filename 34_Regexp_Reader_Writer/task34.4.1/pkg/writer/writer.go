package writer

import (
	"os"

	"task34.4.1/pkg/types"
)

const fileName = "LS.txt"

func WriteInfoAboutFiles(files []types.File) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	for _, file := range files {
		_, _ = f.Write([]byte(file.String()))
		_, _ = f.Write([]byte("\n"))
	}

	return nil
}
