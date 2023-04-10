package typing

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func SimulateType(message string, shellPrompt string, speed float64) {
	if shellPrompt == "" {
		shellPrompt = "$"
	}
	fmt.Print(shellPrompt + " ")
	for i := range message {
		fmt.Print(string(message[i]))
		interval := fmt.Sprintf("%f", 60/float64(speed*60)) + "s"
		timeInterval, err := time.ParseDuration(interval)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Print(interval, timeInterval.Seconds(), time.Second*time.Duration(timeInterval.Seconds()))
		time.Sleep(time.Nanosecond * time.Duration(timeInterval))
	}
	time.Sleep(time.Millisecond * 500)
}
