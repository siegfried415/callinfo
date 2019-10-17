
package callinfo

import (
        "fmt"
        "runtime"
        "strings"
	"strconv"
)

func Goid() int {
	defer func()  {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)     }
	}()
 
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	
	return id
}

func Prefix() string {
        var prefix string  

	// Ask runtime.Callers for up to 10 pcs, including runtime.Callers itself.
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	if n == 0 {
		// No pcs available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return "" 
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	// Loop to get frames.
	// A fixed number of pcs can expand to an indefinite number of Frames.
	for {
		_, more := frames.Next()
                prefix += " "  

		if !more {
			break
		}
	}

        return prefix
}

