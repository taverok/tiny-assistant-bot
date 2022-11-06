package utils

import "fmt"

var (
	ErrorWrongCronRegex = fmt.Errorf("cron must match regex %s", CronRegex)
)
