package handlers

import (
	"fmt"
	"time"
)

func StringToDate(str string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
	}
	return t
}

func DateNow() string {
	dt := time.Now()
	//Format DD-MM-YYYY
	return dt.Format("02/01/2006")
}
