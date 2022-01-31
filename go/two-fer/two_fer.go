package twofer

import "fmt"

// ShareWith determines who you are sharing with
func ShareWith(line string) string {
	if line == "" {
		line = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", line)
}
