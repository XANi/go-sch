package sch

import (
	"sync"
	"fmt"
	"time"
)

// scheduling queue
// is a series of task in monotonic order.
// only appending is possible
type queue struct {
	sync.Mutex
	tasks []*tasklet
}

func newQueue() *queue {
	var q queue
	return &q
}

//
func (q *queue) insert(t *tasklet) error {
	q.Lock()
	if len(q.tasks) > 0 && t.ts.Before(q.tasks[len(q.tasks)-1].ts) {
		q.Unlock()
		return fmt.Errorf("ts of new task before end of queue: q:%s, t:%s",q.tasks[len(q.tasks)-1].ts,t.ts)
	} else {
		q.tasks = append(q.tasks,t)
		q.Unlock()
	}
	return nil
}


//  get a list of runnable tasks + ts of next task in line; remove them from queue
func (q *queue) pop(ts time.Time) ([]*tasklet, time.Time) {
	var mark int
	var t *tasklet
	if len(q.tasks) < 1 {
		return []*tasklet{},time.Unix(0,0)
	}

	q.Lock()
	for mark, t = range q.tasks {
		if ts.Before(t.ts) {break}
	}
	ret := q.tasks[:mark]
	q.tasks = q.tasks[mark:]
	q.Unlock()

	return ret, q.tasks[0].ts
}
