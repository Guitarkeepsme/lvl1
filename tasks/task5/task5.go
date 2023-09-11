/*
	Разработать программу, которая будет последовательно отправлять значения в канал,

а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second))
	defer cancel()
	
	go func(ctx context.Context()) {

	}
	// startTime := time.now()
	// go func() {
	// 	for {
	// 		if startTime.After(startTime.Add(time.Second * 5)) {
	// 			return
	// 		}
	// 	}
	// }
}