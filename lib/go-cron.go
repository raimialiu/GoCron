package lib

import "GoSimpleCron/helpers"

type GoSimpleCron struct {
	config GoSimpleCronConfig
	_obj   func()
}

type GoSimpleCronConfig struct {
	RecoverOnError bool
}

func Instance(config GoSimpleCronConfig) GoSimpleCron {
	return GoSimpleCron{
		config: config,
	}
}

func (g *GoSimpleCron) Schedule(everyTime int32, f func()) {
	helpers.Tick(everyTime, f)
}

func (g *GoSimpleCron) DoEvery(f func()) GoCronSchedule {
	g._obj = f
	return GoCronScheduleInstance(f)
}

func (g *GoSimpleCron) Validate(cronExpression string) {

}

func (g *GoSimpleCron) SetInterval(timeout int32, f func()) {
	helpers.SetInterval(timeout, f)
}
