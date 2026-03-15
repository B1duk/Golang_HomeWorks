package main

import (
	"fmt"
	"regexp"
)

func main() {

	var str string
	fmt.Scan(&str)

	reg := `(\d{1,2}\.\d{1,2}\.\d{4})([А-Яа-я0-23-]+)([А-Яа-я0-23-]+)([А-Яа-я]+)\.(\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,6})`
	re := regexp.MustCompile(reg)
	match := re.FindStringSubmatch(str)

	for i := 1; i < len(match); i++ {
		fmt.Println(match[i])
	}

}
