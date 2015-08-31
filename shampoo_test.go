package main

import "testing"

func TestCommandExecutesNTimes(t *testing.T) {
	var count int
	executor = func(cmd string, args ...string) ([]byte, error) {
		count++
		return nil, nil
	}
	defer func() {
		executor = execCommand
	}()

	executeSerial(10, "foo")
	if count != 10 {
		t.Errorf("Expected command to be run 10 times, but ran %d times", count)
	}
}
