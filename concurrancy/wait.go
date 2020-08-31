// Package concurrancy provides functions for easy concurrancy.
package concurrancy

import (
	"context"
	"sync"
)

// WaitEither will run funcs parallelly, block till one of them finishing and return. Others will receive cancel signals through context.
// If ctx is cancelled, it will return immediately with sending cancel signals.
func WaitEither(ctx context.Context, funcs ...func(context.Context)) {
	ctx, cancel := context.WithCancel(ctx)
	done := make(chan struct{}, 1)

	for _, f := range funcs {
		fn := f
		go func() {
			fn(ctx)

			select {
			case done <- struct{}{}:
			default:
			}
		}()
	}

	select {
	case <-ctx.Done():
	case <-done:
	}

	cancel()
}

// WaitAll will run funcs parallelly, block till all of them finishing and return.
// If ctx is canceled, it will wait till all of funcs cancelled.
func WaitAll(ctx context.Context, funcs ...func(context.Context)) {
	var wg sync.WaitGroup
	wg.Add(len(funcs))

	for _, f := range funcs {
		fn := f
		go func() {
			defer wg.Done()
			fn(ctx)
		}()
	}

	wg.Wait()
}
