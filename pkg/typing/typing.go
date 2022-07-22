package typing

import (
	"fmt"
	"time"

	"github.com/BenHesketh21/tutorial-scripts/pkg/logger"
)

func SimulateType(message string, speed float64) {
	for i := range message {
		fmt.Print(string(message[i]))
		interval := fmt.Sprintf("%f", 60/float64(speed*60)) + "s"
		timeInterval, err := time.ParseDuration(interval)
		if err != nil {
			logger.Error.Fatal(err)
		}
		//fmt.Print(interval, timeInterval.Seconds(), time.Second*time.Duration(timeInterval.Seconds()))
		time.Sleep(time.Nanosecond * time.Duration(timeInterval))
	}
	fmt.Println("")
}
