package services

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func RunScheduledHeartbeatCheckService() {
	cronInstance := cron.New()
	cronInstance.AddFunc("* * * * *", func() {
		fmt.Println("Running cron job...")
		go CheckResourceHeartbeat()
	})
	fmt.Println("Starting cron scheduler...")
	cronInstance.Start()
}
