package Calculate

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile() []string {
	var final []string
	data, err := os.Open("text.txt")
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	defer data.Close()
	scanner:=bufio.NewScanner(data)
	for scanner.Scan(){
		input:=scanner.Text()
		final = append(final,input)
	}
	return final

}
