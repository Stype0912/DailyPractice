package context_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "收到信号，监控退出,time=", time.Now().Unix())
			return
		default:
			fmt.Println(name, "goroutine监控中,time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}
func TestContext(t *testing.T) {
	m := make(map[int]int, 10)
	for i := 1; i <= 10; i++ {
		m[i] = i
	}

	for k, v := range m {
		go func() {
			fmt.Println("k ->", k, "v ->", v)
		}()
	}
}
