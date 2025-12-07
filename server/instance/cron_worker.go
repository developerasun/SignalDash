package instance

import (
	"log"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

type cronWorker struct {
	viper *viper.Viper
}

type CronWorker interface {
	Run() error
}

func NewCronWorker(v *viper.Viper) *cronWorker {
	return &cronWorker{
		viper: v,
	}
}

func (cw *cronWorker) Run() error {
	c := cron.New(cron.WithSeconds())
	expression := cw.viper.GetString("server.cron.expression.every1min")

	count := 0
	_, aErr := c.AddFunc(expression, func() {
		count++
		log.Println("count: ", count)
	})

	if aErr != nil {
		return aErr
	}

	c.Start()
	log.Println("cron_worker.go: worker started")
	return nil
}
