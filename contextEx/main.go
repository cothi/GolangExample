package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
golang, context 패키지

작업을 지시할 때, 작업 가능 시간, 취서, 조건, 작업 명세서 역할

특정 서버에서 request륿 보내는데 production level에서는 적절한 timeout을 설정해주는 것이 좋다.
*/

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "api-key", "my-super-secret-api-key")
}

func doSomethingCool1(ctx context.Context) {
	apiKey := ctx.Value("api-key")
	fmt.Println(apiKey)
}

/* func main() {
	fmt.Println("Go Context Tutorial")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	doSomethingCool(ctx)
} */

func doSomethingCool2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Go Context Tutorial")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()
	go doSomethingCool2(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("oh no, i've exceeded the deadline")
	}

	time.Sleep(2 * time.Second)
}
