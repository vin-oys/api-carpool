package util

import "strings"

func GetContactNumberFromUsername(username string) string {
	trimmed := strings.Replace(username, "(", "", -1)
	return strings.Replace(trimmed, ")", "", -1)
}
