package singleton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	assert.Equal(t, ins1, ins2)
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]Singleton{}

	for i := 0; i < parCount; i++ {
		go func(index int) {
			// 协程阻塞，等待channel被关闭才能继续运行
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}

	// 关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)
	wg.Wait()

	for i := 1; i < parCount; i++ {
		assert.Equal(t, instances[i], instances[i-1])
	}
}
