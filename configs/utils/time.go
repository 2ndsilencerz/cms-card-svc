package utils

import "time"

// InstantTimeString used to return time in string which useful for logging
func InstantTimeString() string {
	return time.Now().Format("2006/01/02 15:04:05")
}
