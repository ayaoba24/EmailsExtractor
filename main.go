package main

import (
	"fmt"
)

func main() {
	text := "Contact us at admin@test.com or support@gmail.com."
	emails := extractEmails(text)
	for _, e := range emails {
		fmt.Println(e)
	}
}