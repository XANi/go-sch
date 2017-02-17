package sch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
	"time"
)

func TestQueue(t *testing.T) {
	q := newQueue()
	Convey("Empty queue is empty",t,func() {
		task, ts := q.pop(time.Unix(1,0))
		So(len(task),ShouldEqual,0)
		So(ts,ShouldResemble,time.Unix(0,0))
	})

	Convey("Insert 10 tasks", t, func() {
		var i int64
		for  i = 1; i <= 10; i++ {
			err := q.insert(newTasklet(time.Unix(0,1000 * i)))
			So(err, ShouldEqual, nil)
		}
	})
	Convey("Fail if task is before last element", t, func() {
		err := q.insert(newTasklet(time.Unix(0,9500)))
		So(err, ShouldNotEqual, nil)
	})
	Convey("Succeed if TS is same as last element", t, func() {
		err := q.insert(newTasklet(time.Unix(0,10000)))
		So(err, ShouldEqual, nil)
	})
	Convey("Get tasks before third one",t,func() {
		taskq1, next_ts1 := q.pop(time.Unix(0,3000))
		So(len(taskq1),ShouldEqual,3)
		So(next_ts1,ShouldResemble,time.Unix(0,4000))
		Convey(fmt.Sprintf("Get next task after %s",next_ts1),func() {
			taskq2,next_ts2 := q.pop(next_ts1)
			So(len(taskq2),ShouldEqual,1)
			So(next_ts2,ShouldResemble,time.Unix(0,5000))
			Convey(fmt.Sprintf("Get next task after %s",next_ts2),func() {
				next_taskq3,next_ts3 := q.pop(time.Unix(0,5000))
				So(len(next_taskq3),ShouldEqual,1)
				So(next_ts3,ShouldResemble,time.Unix(0,6000))
			})
		})
	})
	Convey("no tasks from time that was already popped",t,func() {
			task, ts := q.pop(time.Unix(0,5000))
		_ = ts
			So(len(task),ShouldEqual,0)
	})

}
