package main

import "strings"

func cleanInput(text string) []string {
	sanitzedText := strings.TrimSpace(strings.ToLower(text))
	return strings.Split(sanitzedText, " ")
}