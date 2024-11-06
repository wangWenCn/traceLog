package traceLog

import (
	"bytes"
	"context"
	"runtime"
	"sync"
)

var contextMap = sync.Map{}

func getGoroutineID() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	return string(bytes.Fields(buf[:n])[1])
}

func GetGoroutineContext() context.Context {
	value, ok := contextMap.Load(getGoroutineID())
	if !ok {
		return nil
	}
	ctx, ok := value.(context.Context)
	if !ok {
		return nil
	}
	return ctx
}

func SetGoroutineContext(ctx context.Context) {
	contextMap.Store(getGoroutineID(), ctx)
}

func DelGoroutineContext() {
	contextMap.Delete(getGoroutineID())
}
