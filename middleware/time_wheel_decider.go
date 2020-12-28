package middleware

import (
	"time"

	"github.com/funkygao/golib/timewheel"
)

var tw = timewheel.NewTimeWheel(1*time.Second, 35*60)

func Start() {

}
