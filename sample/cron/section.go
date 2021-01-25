package cron

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/kakami/gocron"
	"github.com/spf13/cobra"
)

func newSectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "section",
		RunE: runSection,
	}
	return cmd
}

func runSection(_ *cobra.Command, _ []string) error {
	// secString := "8:00-12:00|14:00-18:00"

	cron := gocron.NewScheduler()

	for i := 0; i < 5; i++ {
		w := &worker{
			id:        fmt.Sprintf("worker_%d", i),
			secString: buildSecString(),
		}
		w.parseSection(cron)
	}

	_, t := cron.NextRun()
	fmt.Println(t)

	cron.Start()
	return nil
}

func buildSecString() string {
	var idx int32
	var secString string
	secs := rand.Int31n(3) + 1
	delay := time.Duration(rand.Int31n(2))

	t := time.Now().Add(delay * time.Minute)
	tn := t.Add(time.Minute)
	secString = fmt.Sprintf("%02d:%02d-%02d:%02d", t.Hour(), t.Minute(), tn.Hour(), tn.Minute())
	for idx = 1; idx < secs; idx++ {
		t = tn.Add(time.Minute)
		tn = t.Add(time.Minute)
		secString += fmt.Sprintf("|%02d:%02d-%02d:%02d", t.Hour(), t.Minute(), tn.Hour(), tn.Minute())
	}
	fmt.Println(secString)

	return secString
}

type worker struct {
	id        string
	secString string
}

func (w *worker) parseSection(cron *gocron.Scheduler) {
	loc, _ := time.LoadLocation("Local")
	parts := strings.Split(w.secString, "|")
	// var daySeconds int64 = 24 * 3600
	// d := time.Now().Unix() / daySeconds * daySeconds
	tnow := time.Now()
	y, m, n := tnow.Date()
	for idx := range parts {
		fields := strings.Split(parts[idx], "-")
		if len(fields) != 2 {
			continue
		}
		form := "15:04"
		t1, err1 := time.Parse(form, fields[0])
		t2, err2 := time.Parse(form, fields[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("parse [%s-%s] failed\n", fields[0], fields[1])
			continue
		}
		// t1 = time.Unix(d+int64(t1.Hour()*3600+t1.Minute()*60), 0).In(loc)
		// t2 = time.Unix(d+int64(t2.Hour()*3600+t2.Minute()*60), 0).In(loc)
		t1 = time.Date(y, m, n, t1.Hour(), t1.Minute(), 0, 0, loc)
		t2 = time.Date(y, m, n, t2.Hour(), t2.Minute(), 0, 0, loc)
		fmt.Println("start at", t1, "end at", t2)
		if t1.Before(tnow) && t2.After(tnow) {
			cron.Every(24 * time.Hour).From(t1).Immediately().Do(w.start)
		} else {
			cron.Every(24 * time.Hour).From(t1).Do(w.start)
		}
		cron.Every(24 * time.Hour).From(t2).Do(w.stop)
	}
}

func (w *worker) start() {
	fmt.Println(w.id, "start...")
}

func (w *worker) stop() {
	fmt.Println(w.id, "stop...")
}
