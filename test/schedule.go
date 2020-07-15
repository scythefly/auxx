package test

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

var (
	tBegin time.Time
)

type scheduleSession struct {
	idx  int64
	idxx int64
}

func (ss *scheduleSession) tickFunc() {
	ss.idxx++
	fmt.Println("session tick tick", ss.idxx, time.Now().Unix()-tBegin.Unix())
}

func ScheduleTest() {
	tBegin = time.Now()
	s := gocron.NewScheduler()
	// s.Every(2).Seconds().Do(taskPrint)
	s.Every(3).Seconds().Do(taskPrint)

	ss := &scheduleSession{idx: 0}
	s.Every(3).Seconds().Do(taskSessionGrow, ss)
	s.Every(1).Second().Do(ss.tickFunc)

	t := time.Unix((time.Now().Unix()/60+1)*60, 0)
	fmt.Println(t.Minute())
	s.Every(1).Minute().From(&t).Do(func() {
		ttt := time.Now()
		fmt.Println(">> do at right minute", ttt.Minute(), ttt.Second())
	})

	<-s.Start()
}

func taskPrint() {
	fmt.Println(">>> tick tick task print", time.Now().Unix()-tBegin.Unix())
}

func taskSessionGrow(ss *scheduleSession) {
	ss.idx++
	fmt.Println("grow session", ss.idx, time.Now().Unix()-tBegin.Unix())
}
