package helpers

import "time"

func Tick(timeout int32, taskFunc func()) {
	for range time.Tick(time.Second * time.Duration(timeout)) {
		taskFunc()
	}
}

func SetInterval(timeout int32, f func()) {
	ticker := time.NewTicker(time.Second * time.Duration(timeout))

	for {
		<-ticker.C
		f()
	}

}

func SetTimeout(timeout int32, f func()) {
	time.AfterFunc(time.Second*time.Duration(timeout), f)
}
