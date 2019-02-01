package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("here 3")
				return
			default:
				// Do other stuff
				fmt.Println("here 4")
				time.Sleep(10 * time.Second)
				fmt.Println("here 5")
			}
		}
	}()

	// Do stuff
	fmt.Println("here 1")

	time.Sleep(2 * time.Second)
	fmt.Println("here 6")
	// Quit goroutine
	quit <- struct{}{}
	fmt.Println("here 2")

}

// func cancelByContext(ctx context.Context, c net.Conn, wg *sync.WaitGroup) {
// 	defer c.Close()
// 	defer wg.Done()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("cancel goroutine by context:", ctx.Err())
// 			return
// 		default:
// 			_, err := io.WriteString(c, "hello cancelByContext")
// 			if err != nil {
// 				return
// 			}
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
//cannel()
