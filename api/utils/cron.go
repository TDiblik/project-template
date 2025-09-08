package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

type CronJob struct {
	Name     string
	Interval time.Duration
	Task     func(context.Context) error
}

func SetupCronJobs(ctx context.Context) {
	if fiber.IsChild() {
		return // only run cron jobs once (in master process)
	}

	log.Println("Setting up cron jobs (should be only once in logs): start")
	jobs := []CronJob{
		{
			Name:     "Test 1",
			Interval: 30 * time.Minute,
			Task:     exampleCronOne,
		},
		{
			Name:     "Test 2",
			Interval: 50 * time.Minute,
			Task:     exampleCronTwo,
		},
		{
			Name:     "Test 3",
			Interval: 80 * time.Minute,
			Task:     exampleCronThree,
		},
	}
	for _, job := range jobs {
		go startJob(ctx, job)
	}
	log.Println("Setting up cron jobs (should be only once in logs): done")
}

func startJob(ctx context.Context, job CronJob) {
	ticker := time.NewTicker(job.Interval)
	defer ticker.Stop()

	var running bool
	for {
		select {
		case <-ctx.Done():
			Log(fmt.Sprintf("[Cron] Shutting down job: %s", job.Name))
			return
		case <-ticker.C:
			if running {
				Log(fmt.Sprintf("[Cron] Skipping job (still running): %s", job.Name))
				continue
			}

			running = true
			go func() {
				defer func() {
					if r := recover(); r != nil {
						Log(fmt.Sprintf("[Cron] Panic recovered in job %s: %v", job.Name, r))
					}
					running = false
				}()

				if err := job.Task(ctx); err != nil {
					LogErr(fmt.Errorf("[Cron] Error in job %s: %v", job.Name, err))
				}
			}()
		}
	}
}

func exampleCronOne(context context.Context) error {
	Log("[CRON] - test 1, every 30 minutes")
	return nil
}
func exampleCronTwo(context context.Context) error {
	Log("[CRON] - test 2, every 50 minutes")
	return nil
}
func exampleCronThree(context context.Context) error {
	Log("[CRON] - test 3, every 80 minutes, simulates error")
	return fmt.Errorf("fuck")
}
