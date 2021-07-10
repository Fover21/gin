package util

import "testing"

func TestRandomInt32(t *testing.T) {
	var rand int32
	for i := 0; i < 100; i++ {
		rand = RandomInt32(10)
		if rand < 0 || rand > 9 {
			t.Error("invalid random result ", rand)
		}
	}

	rand = RandomInt32(-1)
	if rand != 0 {
		t.Error("not equal 0 as expect")
	}
}
