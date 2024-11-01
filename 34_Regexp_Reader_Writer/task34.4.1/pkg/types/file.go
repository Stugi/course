package types

import "fmt"

type File struct {
	Name  string
	IsDir bool
	Size  int64
}

func (f File) String() string {
	nameDir := "FILE"
	if f.IsDir {
		nameDir = "DIRECTORY"
	}
	return fmt.Sprintf("%s (%d bytes) [%s]", f.Name, f.Size, nameDir)
}
