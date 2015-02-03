// The parameter weekday is True if it is a weekday, and the parameter
// vacation is True if we are on vacation. We sleep in if it is not a
// weekday or we're on vacation. Return True if we sleep in.
//
// SleepIn(False, False) → True
// SleepIn(True, False) → False
// SleepIn(False, True) → True

package sleepin

import "testing"

var sleepInTests = []struct {
	weekday  bool
	vacation bool
	out      bool
}{
	{
		weekday:  true,
		vacation: true,
		out:      true,
	},
	{
		weekday:  false,
		vacation: true,
		out:      true,
	},
	{
		weekday:  true,
		vacation: false,
		out:      false,
	},
	{
		weekday:  false,
		vacation: false,
		out:      true,
	},
}

func TestSleepIn(t *testing.T) {
  var rcvd bool
	for  test := range sleepInTests {
    rcvd = SleepIn(sleepInTests[test].weekday, sleepInTests[test].vacation)
		if rcvd != sleepInTests[test].out {
			t.Error("Test Failed", sleepInTests[test].weekday,sleepInTests[test].vacation, sleepInTests[test].out,rcvd)
     }
	}
}

func SleepIn(weekday, vacation bool) bool {


}
