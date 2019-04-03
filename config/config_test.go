package config

import "testing"

func TestGetEnv(t *testing.T) {
	actual := GetEnv()
	expected := ":8080"
	if actual.Port != expected {
		t.Errorf("TestGetEnv test expected [%+v],\n actual [%+v]", expected, actual.Port)
	}
}
