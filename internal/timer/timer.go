package timer

import (
	. "bots/internal/cache"
	"fmt"
	"time"
)

//check on leak/locks
func StartTimer() {
	for {
		timer := time.NewTimer(15 * time.Second)
		reqPerUserSnap := make(map[string]int)
		for i, v := range CacheApp.RequestsPerUser {
			reqPerUserSnap[i] = v
		}
		<-timer.C
		CacheApp.Bots = 0
		for i := range CacheApp.RequestsPerUser {
			delta := CacheApp.RequestsPerUser[i] - reqPerUserSnap[i]
			fmt.Printf("delta = %v \n", delta)
			if delta > 100 {
				CacheApp.Bots++
			}
		}
	}
}
