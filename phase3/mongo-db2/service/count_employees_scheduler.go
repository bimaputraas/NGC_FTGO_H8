package service

import (
	"fmt"
	"ngc2_p3/repository"
	"time"

	"github.com/robfig/cron/v3"
)

func CountEmployeesScheduler(r repository.Repository) {
	scheduler := cron.New()
	scheduler.AddFunc("*/1 * * * *",func() {
		count,err := r.CountEmployees()
		if err != nil {
			fmt.Println(err)
			return
		}
		formatString := fmt.Sprintf("%v - Total registered employees now = %d",time.Now().Format("2006-01-02 15:04:05"),count)

		fmt.Println(formatString)
	})

	scheduler.Start()
}
