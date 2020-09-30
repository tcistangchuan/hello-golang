package singleton

import "testing"

func TestGetInstance2(t *testing.T) {
	instance1 := GetInstance2()
	instance2 := GetInstance2()
	if instance1 != instance2 {
		t.Fatal("instance is not equal")
	}
	t.Log(" server success")
}
