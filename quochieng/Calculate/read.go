package Calculate

import (
	"fmt"
	"os"
)

type Status struct {
	Name    string
	Account string
	Balance int
}

func ReadFile() string {
	var final string
	file, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	}
	final = string(file)

	return final
}
