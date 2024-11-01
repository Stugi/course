package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	rf "task34.6.1/pkg/reader"
	mathexample "task34.6.1/pkg/types"
	wf "task34.6.1/pkg/writer"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите полный путь файла: ")
	inPath, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	inPath = strings.TrimSpace(inPath)
	fmt.Println("Ваш файл: ", inPath)

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Введите полный путь для записи: ")
	outPath, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	outPath = strings.TrimSpace(outPath)
	fmt.Println("Ваш путь для записи: ", outPath)

	// outPath := "./out.txt"
	// inPath := "/Users/arinastugireva/Documents/projects/learning/course/34_Regexp_Reader_Writer/task34.6.1/example.txt"

	examples, err := rf.GetStringsFromFile(inPath)

	if err != nil {
		fmt.Errorf("GetFilesFromPath() = %v", err)
	}

	for i, example := range examples {
		examples[i] = mathexample.New(example).EndString

	}
	err = wf.WriteInfoAboutFiles(outPath, examples)

	if err != nil {
		fmt.Errorf("WriteInfoAboutFiles() = %v", err)
	}
}
