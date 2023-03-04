package main

import (
	"fmt"
	"time"

	util "github.com/DevopsGuyXD/IAM-Access-Key-Rotation/Utils"
)

var check_status string = "(2/2)Sending notification ."

//======================= FUNCTION =======================
func SendNotifications(users []string, accesskey_dates []string) {

	current_date := time.Now()

	for i := 0; i < len(accesskey_dates); i++ {

		if accesskey_dates[i] == "No access key" {
			continue
		} else {
			date_string := accesskey_dates[i]
			date_parsed, err := time.Parse("2006-01-02", date_string)
			util.CheckForNil(err)

			difference := current_date.Sub(date_parsed)

			fmt.Printf("%v",check_status)
			check_status = "."

			if int(difference.Hours()/24) > 90 {
				email_handler(users[i])
			}
		}
	}
}