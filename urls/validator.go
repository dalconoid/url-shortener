package urls

import (
	"fmt"
	"regexp"
)

const urlRegex = `^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`
const slugRegex = `^[a-zA-Z0-9\-]{3,20}$`
var bannedSlugs =[]string{"alive", "register-url"}

// ValidateURL checks if url is correct
func ValidateURL(url string) bool {
	matched, _ := regexp.MatchString(urlRegex, url)
	return matched
}

// ValidateSlug checks if slug is correct
func ValidateSlug(slug string) bool {
	for _, s := range bannedSlugs {
		if slug == s {
			fmt.Println("banned slug " + s)
			return false
		}
	}
	matched, _ := regexp.MatchString(slugRegex, slug)
	return matched
}