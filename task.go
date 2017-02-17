package sch

import (
	"time"
)


type tasklet struct {
	ts time.Time
}

func newTasklet(ts time.Time) *tasklet {
	var t tasklet
	t.ts = ts
	return &t
}
