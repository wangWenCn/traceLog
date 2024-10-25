package traceLog

import (
	"bytes"
	"context"
	"runtime"
)

var (
	ContextMap = make(map[string]context.Context)
)

func getGoroutineID() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	return string(bytes.Fields(buf[:n])[1])
}

func GetGoroutineContext() context.Context {
	return ContextMap[getGoroutineID()]
}

func SetGoroutineContext(ctx context.Context) {
	ContextMap[getGoroutineID()] = ctx
}

func DelGoroutineContext() {
	delete(ContextMap, getGoroutineID())
}
