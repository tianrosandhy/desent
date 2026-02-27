package utils

import (
	"regexp"
	"strings"
)

func Like(keyword string) string {
	likeKeyword := strings.ReplaceAll(keyword, " ", "%")
	likeKeyword = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(likeKeyword, "%")
	likeKeyword = "%" + likeKeyword + "%"

	return likeKeyword
}
