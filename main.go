package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shuntagami/tenji-maker-go/tenjibuilder"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	// ReadString will block until the delimiter is entered
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("failed")
		return
	}

	// remove the delimeter from the string
	text = strings.TrimSuffix(text, "\n")

	fmt.Println(tenjibuilder.Build(text))
}
