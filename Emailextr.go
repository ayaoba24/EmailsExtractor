package main

import (
	"strings"
)


func extractEmails(text string) []string {
	var emails []string
	words := strings.Fields(text)

	for _, word := range words {
		word = strings.Trim(word, ".,!?;:<>\"'()")

		if strings.Contains(word, "@") {
			parts := strings.Split(word, "@")

			if len(parts) == 2 && parts[0] != "" && strings.Contains(parts[1], ".") {
				emails = append(emails, word) 
			}
		}
	}
	return emails
}