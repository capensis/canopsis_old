package correlation

import (
	"math"
	"sort"
	"time"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/datetime"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
)

const (
	Opened = iota
	Ready
	Created
)

type EventExtraInfosMeta struct {
	Rule      Rule
	Count     int64
	LastChild types.AlarmWithEntity
	// deprecated
	Children types.AlarmWithEntity
}

type MetaAlarmState struct {
	ID string `bson:"_id"`

	// MetaAlarmName is meta alarm entity name
	MetaAlarmName          string `bson:"meta_alarm_name"`
	MetaAlarmComponentName string `bson:"meta_alarm_component_name"`

	Version int64 `bson:"version"`
	State   int   `bson:"state"`

	ChildrenEntityIDs  []string `bson:"children_entity_ids"`
	ChildrenTimestamps []int64  `bson:"children_timestamps"`

	// only for corel rule type
	ParentsEntityIDs  []string `bson:"parents_entity_ids"`
	ParentsTimestamps []int64  `bson:"parents_timestamps"`

	// ExpiredAt is used to remove expired state.
	// Type should be time.Time for ttl index.
	ExpiredAt time.Time `bson:"expired_at,omitempty"`

	ChildInactiveExpireAt *datetime.CpsTime `bson:"child_inactive_expire_at,omitempty"`
}

func (s *MetaAlarmState) Reset(id string) {
	*s = MetaAlarmState{
		ID:                     id,
		MetaAlarmComponentName: "",
		MetaAlarmName:          "",
		State:                  Opened,
		// 1 cap for a possible new alarm in the group
		ChildrenEntityIDs:  make([]string, 0, 1),
		ChildrenTimestamps: make([]int64, 0, 1),
		ParentsEntityIDs:   make([]string, 0, 1),
		ParentsTimestamps:  make([]int64, 0, 1),
	}
}

func (s *MetaAlarmState) ResetExpireAt(timeIntervalInSeconds int64) {
	s.ExpiredAt = time.Now().Add(time.Duration(timeIntervalInSeconds) * time.Second)
}

func (s *MetaAlarmState) ResetChildInactiveExpireAt(childInactiveDelay *datetime.DurationWithUnit) {
	if childInactiveDelay != nil && childInactiveDelay.Value > 0 {
		childInactiveExpireAt := childInactiveDelay.AddTo(datetime.NewCpsTime())
		s.ChildInactiveExpireAt = &childInactiveExpireAt
	}
}

func (s *MetaAlarmState) IsOutdated(alarmLastUpdate, timeInterval int64) bool {
	openTime := s.GetParentOpenTime()
	childOpenTime := s.GetChildrenOpenTime()
	if childOpenTime < openTime {
		openTime = childOpenTime
	}

	return s.ID == "" || alarmLastUpdate > openTime+timeInterval && s.State != Opened
}

func (s *MetaAlarmState) PushChild(entityID string, timestamp int64, ruleTimeInterval int64) []string {
	for idx, v := range s.ChildrenEntityIDs {
		if v == entityID {
			s.ChildrenTimestamps = append(s.ChildrenTimestamps[:idx], s.ChildrenTimestamps[idx+1:]...)
			s.ChildrenEntityIDs = append(s.ChildrenEntityIDs[:idx], s.ChildrenEntityIDs[idx+1:]...)

			break
		}
	}

	// if length = 0 => init times and ids with single values
	if len(s.ChildrenTimestamps) == 0 {
		s.ChildrenTimestamps = []int64{timestamp}
		s.ChildrenEntityIDs = []string{entityID}

		return nil
	}

	// times is always sorted, so the 0 element is always open timestamp
	openTimestamp := s.ChildrenTimestamps[0]

	//if alarm is late
	if timestamp < openTimestamp {
		//check if interval can be shifted, if any alarm in the Group will be lost => then we cannot shift time
		if s.ChildrenTimestamps[len(s.ChildrenTimestamps)-1] > timestamp+ruleTimeInterval {
			return nil
		}

		// Push to front, because it's new minimal value
		s.ChildrenTimestamps = append([]int64{timestamp}, s.ChildrenTimestamps...)
		s.ChildrenEntityIDs = append([]string{entityID}, s.ChildrenEntityIDs...)

		return nil
	}

	newAlarmRuleStartTime := timestamp - ruleTimeInterval
	if newAlarmRuleStartTime > openTimestamp {
		idx := sort.Search(len(s.ChildrenTimestamps), func(i int) bool {
			return s.ChildrenTimestamps[i] >= newAlarmRuleStartTime
		})

		removedEntityIDs := make([]string, idx)
		copy(removedEntityIDs, s.ChildrenEntityIDs[:idx])

		//remove outdated from front, push to the end, because it's the oldest value
		s.ChildrenTimestamps = append(s.ChildrenTimestamps[idx:], timestamp)
		s.ChildrenEntityIDs = append(s.ChildrenEntityIDs[idx:], entityID)

		return removedEntityIDs
	}

	//insert to insertIdx to keep sort
	insertIdx := sort.Search(len(s.ChildrenTimestamps), func(i int) bool {
		return s.ChildrenTimestamps[i] >= timestamp
	})

	s.ChildrenTimestamps = append(s.ChildrenTimestamps[:insertIdx], append([]int64{timestamp}, s.ChildrenTimestamps[insertIdx:]...)...)
	s.ChildrenEntityIDs = append(s.ChildrenEntityIDs[:insertIdx], append([]string{entityID}, s.ChildrenEntityIDs[insertIdx:]...)...)

	return nil
}

func (s *MetaAlarmState) PushParent(entityID string, timestamp int64, ruleTimeInterval int64) {
	for idx, v := range s.ParentsEntityIDs {
		if v == entityID {
			s.ParentsTimestamps = append(s.ParentsTimestamps[:idx], s.ParentsTimestamps[idx+1:]...)
			s.ParentsEntityIDs = append(s.ParentsEntityIDs[:idx], s.ParentsEntityIDs[idx+1:]...)

			break
		}
	}

	// if length = 0 => init times and ids with single values
	if len(s.ParentsTimestamps) == 0 {
		s.ParentsTimestamps = []int64{timestamp}
		s.ParentsEntityIDs = []string{entityID}

		return
	}

	// times is always sorted, so the 0 element is always open timestamp
	openTimestamp := s.ParentsTimestamps[0]

	//if alarm is late
	if timestamp < openTimestamp {
		//check if interval can be shifted, if any alarm in the Group will be lost => then we cannot shift time
		if s.ParentsTimestamps[len(s.ParentsTimestamps)-1] > timestamp+ruleTimeInterval {
			return
		}

		// Push to front, because it's new minimal value
		s.ParentsTimestamps = append([]int64{timestamp}, s.ParentsTimestamps...)
		s.ParentsEntityIDs = append([]string{entityID}, s.ParentsEntityIDs...)

		return
	}

	newAlarmRuleStartTime := timestamp - ruleTimeInterval
	if newAlarmRuleStartTime > openTimestamp {
		idx := sort.Search(len(s.ParentsTimestamps), func(i int) bool {
			return s.ParentsTimestamps[i] >= newAlarmRuleStartTime
		})

		//remove outdated from front, push to the end, because it's the oldest value
		s.ParentsTimestamps = append(s.ParentsTimestamps[idx:], timestamp)
		s.ParentsEntityIDs = append(s.ParentsEntityIDs[idx:], entityID)

		return
	}

	//insert to insertIdx to keep sort
	insertIdx := sort.Search(len(s.ParentsTimestamps), func(i int) bool {
		return s.ParentsTimestamps[i] >= timestamp
	})

	s.ParentsTimestamps = append(s.ParentsTimestamps[:insertIdx], append([]int64{timestamp}, s.ParentsTimestamps[insertIdx:]...)...)
	s.ParentsEntityIDs = append(s.ParentsEntityIDs[:insertIdx], append([]string{entityID}, s.ParentsEntityIDs[insertIdx:]...)...)
}

func (s *MetaAlarmState) GetChildrenOpenTime() int64 {
	if len(s.ChildrenTimestamps) != 0 {
		return s.ChildrenTimestamps[0]
	}

	return math.MaxInt64
}

func (s *MetaAlarmState) GetParentOpenTime() int64 {
	if len(s.ParentsTimestamps) != 0 {
		return s.ParentsTimestamps[0]
	}

	return math.MaxInt64
}

func (s *MetaAlarmState) RemoveChildrenBefore(timestamp int64) {
	idx := sort.Search(len(s.ChildrenTimestamps), func(i int) bool {
		return s.ChildrenTimestamps[i] >= timestamp
	})

	s.ChildrenTimestamps = s.ChildrenTimestamps[idx:]
	s.ChildrenEntityIDs = s.ChildrenEntityIDs[idx:]
}

func (s *MetaAlarmState) RemoveParentsBefore(timestamp int64) {
	idx := sort.Search(len(s.ParentsTimestamps), func(i int) bool {
		return s.ParentsTimestamps[i] >= timestamp
	})

	s.ParentsTimestamps = s.ParentsTimestamps[idx:]
	s.ParentsEntityIDs = s.ParentsEntityIDs[idx:]
}

func (s *MetaAlarmState) GetEntityID() string {
	if s.MetaAlarmName == "" {
		return ""
	}

	if s.MetaAlarmComponentName == "" {
		return s.MetaAlarmName
	}

	return s.MetaAlarmName + "/" + s.MetaAlarmComponentName
}
