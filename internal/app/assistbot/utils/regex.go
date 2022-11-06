package utils

import "regexp"

var CronRegex = regexp.MustCompile("[0-9*] [0-9*] [0-9*] [0-9*] [0-9*]")
var MultiSpaceRegex = regexp.MustCompile(" {2,}")
