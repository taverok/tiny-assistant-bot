package utils

import "regexp"

var CronRegex = regexp.MustCompile("(([*]|[0-9,]+) ){4}([*]|[0-9,]+)")
