// We have two monkeys, a and b, and the parameters aSmile and bSmile indicate
// if each is smiling. We are in trouble if they are both smiling or if neither
// of them is smiling. Return true if we are in trouble.
//
// monkeyTrouble(true, true) → true
// monkeyTrouble(false, false) → true
// monkeyTrouble(true, false) → false


package main

import "log"

func monkeyTrouble(a_smile, b_smile bool) bool {
  // #TODO: Your code goes here

}

var monkeyTroubleTests = []struct {
  a_smile  bool
  b_smile  bool
  out      bool
}{
  {
    a_smile:  true,
    b_smile: true,
    out:      true,
  },
  {
    a_smile:  false,
    b_smile: true,
    out:      false,
  },
  {
    a_smile:  true,
    b_smile: false,
    out:      false,
  },
  {
    a_smile:  false,
    b_smile: false,
    out:      true,
  },
}

func main() {
  var rcvd bool
  for test := range monkeyTroubleTests {
    rcvd = monkeyTrouble(monkeyTroubleTests[test].a_smile, monkeyTroubleTests[test].b_smile)
    if rcvd != monkeyTroubleTests[test].out {
      log.Println("FAIL: Test Failed a_smile: ",monkeyTroubleTests[test].a_smile," b_smile: ",monkeyTroubleTests[test].b_smile," expected: ",monkeyTroubleTests[test].out," received: ", rcvd)

    }else {
      log.Println("PASS: Test Passed a_smile: ",monkeyTroubleTests[test].a_smile," b_smile: ",monkeyTroubleTests[test].b_smile," expected: ",monkeyTroubleTests[test].out," received: ", rcvd)

    }
  }
}
