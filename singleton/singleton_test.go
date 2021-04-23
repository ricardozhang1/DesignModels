package singleton

import (
	"sync"
	"testing"
)

const partCount = 100

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()

	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParalleSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(partCount)
	instances := [partCount]*Singleton{}
	for i:=0; i<partCount; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i:=1; i<partCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}



