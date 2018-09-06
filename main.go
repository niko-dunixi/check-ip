package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	if body, err := http.Get(`http://checkip.amazonaws.com/`); err == nil {
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(body.Body)
		fmt.Println(strings.TrimSpace(buffer.String()))
	} else {
		errorMessage := fmt.Sprintf("+%v", err)
		fmt.Println(errorMessage)
		os.Exit(1)
	}
}
