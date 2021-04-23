package simpleFactory

import "testing"

// TestType1 test get hiapi with factory
func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s!="Hi, Tom" {
		t.Fatal("Type1 test fail!")
	}
	t.Log("Success 1")
}

// TestType2 test get hiapi with factory
func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "Hello, Tom" {
		t.Fatal("Type2 test fail!")
	}
	t.Log("Success 2")
}





