package errorinfo

import (
	"log"
)

func Check(err error, fatal bool) {
	if err != nil {
		if fatal {
			log.Fatal("[FATAL ERROR] ", err)
		} else {
			log.Printf("[ERROR] %v\n", err)
		}
	}
}
