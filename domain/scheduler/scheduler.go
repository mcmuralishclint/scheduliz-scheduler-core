package scheduler

import (
	"context"
	"fmt"
	"github.com/robfig/cron"
)

func ScheduleJobs(ctx context.Context) error {
	schedulerCron := cron.New()
	schedulerCron.AddFunc("@every 1m", func() {
		fmt.Println("Starting the cron to fetch the items to queue")
		fmt.Println("Ending the cron to fetch the items to queue")
	})
	schedulerCron.Start()
	select {}
	return nil
}
