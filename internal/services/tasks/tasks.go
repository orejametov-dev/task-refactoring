package tasks

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"task/internal/app"
	"task/internal/domain"
	"text/tabwriter"
	"time"
)

func Creator(newTasks chan domain.Task) {
	for {
		go func() {
			task := domain.Task{
				ID:        uuid.New().ID(),
				CreatedAt: time.Now(),
				Status:    domain.Status(domain.CREATED),
			}

			if task.CreatedAt.Nanosecond()%2 > 0 { // if it is compiled in macOS, you need to change nanosec to millisec
				task.Status = domain.Status(domain.CREATION_ERROR)
			}

			newTasks <- task
		}()

		time.Sleep(app.ARTIFICAIAL_APPLICATION_PAUSE)
	}
}

func Worker(newTasks chan domain.Task, completedTasks chan domain.Task) {
	for task := range newTasks {
		go func(task domain.Task) {
			if task.CreatedAt.After(time.Now().Add(app.MAX_EXECUTION_TIME)) {
				task.Status = domain.Status(domain.EXECUTED)
			}

			task.FinishedAt = time.Now()
			completedTasks <- task
		}(task)
	}
}

func Print(tasks chan domain.Task) {
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 4, ' ', 0)

	for task := range tasks {
		output := fmt.Sprintf(
			"%d\t%s\t%s\t%s\n",
			task.ID,
			task.CreatedAt.Format(time.RFC3339),
			task.FinishedAt.Format(time.RFC3339),
			task.Status,
		)

		if _, err := w.Write([]byte(output)); err != nil {
			log.Println(err)
		}

		if err := w.Flush(); err != nil {
			log.Println(err)
		}
	}
}
