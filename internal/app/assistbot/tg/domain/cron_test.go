package domain

import (
	"errors"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/utils"
	"testing"
)

type cronProbe struct {
	in          string
	expected    Cron
	expectedErr error
}

var probes = []cronProbe{
	{in: "1 1 1 1 1", expected: Cron{Minute: "1", Hour: "1", DayMonth: "1", Month: "1", DayWeek: "1"}},
	{in: "1 * 1 * 0", expected: Cron{Minute: "1", Hour: "*", DayMonth: "1", Month: "*", DayWeek: "0"}},
	{in: "a b c d e", expectedErr: utils.ErrorWrongCronRegex},
}

func TestCronFromString(t *testing.T) {
	for _, probe := range probes {
		actual, actualErr := CronFromString(probe.in)
		if !errors.Is(actualErr, probe.expectedErr) {
			t.Errorf("actual error: %v \n expected: %v \n", actualErr, probe.expectedErr)
			continue
		}

		if actual != probe.expected {
			t.Errorf("actual: %v \n expected: %v\n", actual, probe.expected)
		}
	}
}
