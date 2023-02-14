package main

import (
	"testing"
)

func TestAddition3(t *testing.T) {
    sum := addition3(1,2);
    if sum == 3 {
        t.Log("Recieved as expected: 3")
    } else {
        t.Error(sum)
    }
}
