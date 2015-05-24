package logger

import (
	"fmt"
	"runtime"
	"testing"
)

func TestFailLog(t *testing.T) {
	fmt.Println(runtime.GOARCH)
}
