package indicator

import (
	"testing"
	"time"

	"github.com/developerasun/SignalDash/server/config"
	"github.com/developerasun/SignalDash/server/models"
	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestIndicatorInsert(t *testing.T) {
	t.Skip()
	assert := assert.New(t)
	testdb, oErr := gorm.Open(sqlite.Open("../../sd_app.test.db"), &gorm.Config{})

	if oErr != nil {
		t.Logf("failed to open test db: %s", oErr.Error())
		t.Fail()
	}
	testdb.AutoMigrate(&models.Indicator{})
	indicatorRepo := IndicatorRepo{db: testdb}
	cErr := indicatorRepo.Create(&models.Indicator{
		Name:   "apple",
		Ticker: "aapl",
		Value:  1.00,
		Domain: "apple.com",
	})

	assert.Nil(cErr)
}

func TestCronExpression(t *testing.T) {
	assert := assert.New(t)
	c := cron.New()

	options := &config.ViperOptions{
		Filename:  "options",
		ConfigDir: "../../config",
	}
	v := options.InitConfig()

	expression := v.GetString("cron.expression.every1min")

	count := 0
	c.AddFunc(expression, func() {
		if count < 3 {
			t.Log(time.Now().UTC())
			count++
		} else {
			c.Stop()
		}
	})

	<-time.After(time.Second * 3)

	assert.Equal(count, 2)
}
