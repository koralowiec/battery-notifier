package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/distatus/battery"
	"github.com/keybase/go-notifier"
)

const notificationTitle = "Low Battery!"
const notificationMsgTemplate = "Battery at %.2f%%"

var batteryThreshold float64
var minutesBetweenCheck int

func sendNotification(currentPercent float64) {
	message := fmt.Sprintf(notificationMsgTemplate, currentPercent)

	notification := notifier.Notification{}
	notification.Title = notificationTitle
	notification.Message = message

	notifier, err := notifier.NewNotifier()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := notifier.DeliverNotification(notification); err != nil {
		log.Fatal(err)
		return
	}
}

func checkBatteries() {
	batteries, err := battery.GetAll()

	if err != nil {
		log.Fatal("Could not get battery info!")
		return
	}

	for _, bat := range batteries {
		percent := bat.Current / bat.Full * 100

		if percent <= batteryThreshold && bat.State == battery.Discharging {
			sendNotification(percent)
		}
	}
}

func init() {
	flag.Float64Var(&batteryThreshold, "t", 30.0, "Battery threshold. Battery below this level (and battery in discharging state) will cause sending notification.")
	flag.IntVar(&minutesBetweenCheck, "m", 1, "Checking interval in minutes.")
	flag.Parse()
}

func main() {
	ticker := time.NewTicker(time.Duration(minutesBetweenCheck) * time.Minute)

	for range ticker.C {
		go checkBatteries()
	}
}
