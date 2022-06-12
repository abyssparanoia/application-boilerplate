package now

import (
	"flag"
	"time"
)

var JST = time.FixedZone("JST", 9*60*60)

var Now = func() time.Time {
	n := time.Now()
	nJST := n.In(JST)
	return nJST
}

func FakeNow(t time.Time) {
	t = t.In(JST)
	if flag.Lookup("test.v") != nil {
		Now = func() time.Time { return t }
	}
}
