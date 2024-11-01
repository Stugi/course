package reader

import (
	"os"

	"task34.4.1/pkg/types"
)

type ReaderFile interface {
	GetFilesFromPath(path string) ([]types.File, error)
}

func GetFilesFromPath(path string) ([]types.File, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	filesList := make([]types.File, 0)

	for _, f := range files {
		fInfo, _ := f.Info()
		filesList = append(filesList, types.File{
			Name:  f.Name(),
			IsDir: f.IsDir(),
			Size:  fInfo.Size(),
		})
	}

	return filesList, nil
}
