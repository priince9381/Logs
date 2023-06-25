package models

import "time"

type Log struct {
	ID        int    `json:"id"`
	UnixTS    int64  `json:"unix_ts"`
	UserID    int    `json:"user_id"`
	EventName string `json:"event_name"`
}

func (l *Log) GetId() int {
	if l == nil {
		return 0
	}
	return l.ID
}

func (l *Log) GetUnixTs() time.Time {
	if l == nil {
		return time.Now()
	}
	return time.Unix(l.UnixTS, 0)
}

func (l *Log) GetUserId() int {
	if l == nil {
		return 0
	}
	return l.UserID
}

func (l *Log) GetEventName() string {
	if l == nil {
		return ""
	}
	return l.EventName
}
