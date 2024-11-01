package reader

import (
	"io"
	"os"
	"regexp"
)

const regex = `\d+[+-]\d+=\?`

func GetStringsFromFile(path string) (ret []string, err error) {
	var str string
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1)
	re := regexp.MustCompile(regex)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}

		str = str + string(buf[:n])
		if re.FindAllStringIndex(str, -1) != nil {
			indexes := re.FindAllStringIndex(str, -1)
			ret = append(ret, str[indexes[len(indexes)-1][0]:indexes[len(indexes)-1][1]])
			str = str[indexes[len(indexes)-1][1]:]
		}
	}

	return ret, err
}
