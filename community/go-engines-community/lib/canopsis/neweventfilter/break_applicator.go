package neweventfilter

import (
	"context"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/eventfilter/pattern"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
)

type breakApplicator struct {}

func NewBreakApplicator() RuleApplicator {
	return &breakApplicator{}
}

func (a *breakApplicator) Apply(_ context.Context, _ Rule, event types.Event, _ pattern.EventRegexMatches) (string, types.Event, error) {
	return OutcomeBreak, event, nil
}
