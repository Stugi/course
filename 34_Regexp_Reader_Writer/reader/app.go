package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Пример чтения данных из файла:

func main() {

	body := strings.NewReader("hello world")

	req, _ := http.NewRequest("POST", "https://google.com", body)
	client := http.Client{}
	response, _ := client.Do(req)
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
