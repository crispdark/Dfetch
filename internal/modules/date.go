package modules

import "time"

func Date() string {
	return time.Now().Format("2006-01-02")
}
