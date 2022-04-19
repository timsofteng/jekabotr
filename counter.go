package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func ReadExactLine(lineNumber int) string {
	number := fmt.Sprint(lineNumber, "p")
	cmd := exec.Command("sed", "-n", number, "text.txt")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}

func countNumbers() int {
	cmd := exec.Command("sed", "-n", "$=", "text.txt")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	outStr := out.String()

	outStr = strings.TrimSuffix(outStr, "\n")

	// string to int
	i, err := strconv.Atoi(outStr)

	if err != nil {
		log.Fatal(err)
	}

	return i
}
