package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/distatus/battery"
	"github.com/keybase/go-notifier"
)

const notificationTitle = "Low Battery!"
const notificationMsgTemplate = "Battery at %.2f%%"

var batteryThreshold float64
var minutesBetweenCheck int

const envKey = "XDG_RUNTIME_DIR"
const envValueTemplate = "/run/user/%d"

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

func setEnv() {
	uid := os.Geteuid()
	value := fmt.Sprintf(envValueTemplate, uid)
	if err := os.Setenv(envKey, value); err != nil {
		log.Fatalf("Could not set the environment variable: %s\n", envKey)
	}
}

func init() {
	flag.Float64Var(&batteryThreshold, "t", 30.0, "Battery threshold. Battery below this level (and battery in discharging state) will cause sending notification.")
	flag.Parse()
}

func main() {
	setEnv()
	checkBatteries()
}
