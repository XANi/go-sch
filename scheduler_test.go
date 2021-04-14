package sch

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := newQueue()
	t.Run("empty queue", func(t *testing.T) {
		task, ts := q.pop(time.Unix(1,0))
		assert.Len(t,task,0,"empty task list")
		assert.True(t,ts.Equal(time.Unix(0,0)),"zero ts with empty task list")
	})
	t.Run("insert 10 tasks", func(t *testing.T) {
		for  i  := int64(1); i <= 10; i++ {
			err := q.insert(newTasklet(time.Unix(0,1000 * i)))
			assert.Nil(t,err)
		}
	})

	t.Run("fail if task is added before last element", func(t *testing.T) {
		err := q.insert(newTasklet(time.Unix(0,9500)))
		assert.Error(t,err)

	})
	t.Run("succeed if new element ts is same as last", func(t *testing.T) {
		err := q.insert(newTasklet(time.Unix(0,10000)))
		assert.Nil(t,err)
	})
	t.Run("get task before 3rd", func(t *testing.T) {
		taskq1, next_ts1 := q.pop(time.Unix(0,3000))
		assert.Len(t,taskq1,3)
		assert.Equal(t,time.Unix(0,4000),next_ts1)
		t.Run("get task after 4us", func(t *testing.T) {
			taskq2,next_ts2 := q.pop(next_ts1)
			assert.Len(t,taskq2,1)
			assert.Equal(t,time.Unix(0,5000),next_ts2)

			t.Run("get task after 5us", func(t *testing.T) {
				taskq3,next_ts3 := q.pop(next_ts2)
				assert.Len(t,taskq3,1)
				assert.Equal(t,time.Unix(0,6000),next_ts3)
			})
		})


	})
	t.Run("no tasks from time that was already popped",func(t *testing.T) {
		task, ts := q.pop(time.Unix(0,5000))
		_ = ts
		assert.Len(t,task,0)
	})

}



func BenchmarkInsert(b *testing.B) {
	// run the Fib function b.N times
	q := newQueue()
	for n := 0; n < b.N; n++ {
			q.insert(newTasklet(time.Unix(0,1000 * int64(n))))
	}
}

func BenchmarkPop(b *testing.B) {
	q := newQueue()
	for n := 0; n < b.N; n++ {
		q.insert(newTasklet(time.Unix(0,1000 * int64(n))))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		q.pop((time.Unix(0, 1000*int64(n))))
	}

}