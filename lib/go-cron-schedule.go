package lib

import (
	"GoCron/helpers"
)

type GoCronSchedule struct {
	cronFunc func()
}

func GoCronScheduleInstance(f func()) GoCronSchedule {
	return GoCronSchedule{
		f,
	}
}

func (g GoCronSchedule) Minute() {
	helpers.Tick(helpers.EVERY_MINUTE, g.cronFunc)
}

func (g GoCronSchedule) Hour() {
	helpers.Tick(helpers.EVERY_HOUR, g.cronFunc)
}

func (g GoCronSchedule) Day() {
	helpers.Tick(helpers.EVERY_DAY, g.cronFunc)
}

func (g GoCronSchedule) Second() {
	helpers.Tick(helpers.EVERY_SECOND, g.cronFunc)
}
