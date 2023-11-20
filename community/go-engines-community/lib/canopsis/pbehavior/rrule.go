package pbehavior

import (
	"time"

	libtime "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/time"
	librrule "github.com/teambition/rrule-go"
)

const rruleEndMaxYears = 10

func GetRruleEnd(
	start libtime.CpsTime,
	rrule string,
	loc *time.Location,
) (*libtime.CpsTime, error) {
	if rrule == "" {
		return nil, nil
	}

	rOption, err := librrule.StrToROption(rrule)
	if err != nil {
		return nil, err
	}

	if rOption.Until.IsZero() && rOption.Count == 0 {
		return nil, nil
	}

	rOption.Dtstart = start.Time.In(loc)
	r, err := librrule.NewRRule(*rOption)
	if err != nil {
		return nil, err
	}

	before := time.Now().In(loc).AddDate(rruleEndMaxYears, 0, 0)
	t := r.Before(before, true)
	if t.IsZero() {
		return nil, nil
	}

	return &libtime.CpsTime{Time: t}, nil
}
