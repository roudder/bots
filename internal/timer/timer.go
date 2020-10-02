package timer

import (
	. "bots/internal/cache"
	"time"
)

//check on leak/locks
func StartTimer() {
	for {
		timer := time.NewTimer(time.Minute)
		reqPerUserSnap := make(map[string]int)
		for i, v := range CacheApp.RequestsPerUser {
			reqPerUserSnap[i] = v
		}
		<-timer.C
		CacheApp.Bots = 0
		for i := range CacheApp.RequestsPerUser {
			delta := CacheApp.RequestsPerUser[i] - reqPerUserSnap[i]
			if delta > 100 {
				CacheApp.Bots++
			}
		}
	}
}
