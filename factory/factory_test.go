package factory

import "testing"

func TestFactory(t *testing.T) {
	log := Factory("log")
	startRes := log.Start()
	expected := "logStart"
	if startRes != expected {
		t.Errorf("expected err to be: %s, but it was %s", expected, startRes)
	}

	stopRes := log.Stop()
	expected = "logStop"
	if stopRes != expected {
		t.Errorf("expected err to be: %s, but it was %s", expected, stopRes)
	}
}
