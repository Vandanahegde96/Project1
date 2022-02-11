package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"encoding/json"

    "project1/service"
)

func main() {
	var inputText string
	fmt.Print("Enter Text: ")

	reader := bufio.NewReader(os.Stdin)
    inputText, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}


	output := service.CountStringsInText(inputText)

    // Marshal output to json format
	jsonout, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonout))
}
