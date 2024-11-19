package Calculate

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Status struct {
	Name      string
	Account   string
	Remaining int
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

func (val *Status) Update() {
	data := ReadFile()
	values := strings.Fields(data)

	balance, err := strconv.Atoi(values[2])
	if err != nil {
		fmt.Printf("%s", err)
	}
	val.Remaining = balance
}
