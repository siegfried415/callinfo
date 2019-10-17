
package callinfo

import (
        "fmt"
        "runtime"
        "strings"
	"strconv"
)

/*
func Prefix() string {
        pc := make([]uintptr, 1024)
        frames := runtime.CallersFrames(pc)

        more := true
        var prefix string  
        for more {
                _, more = frames.Next()
                prefix += " "  
        }

        return prefix
}
*/

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

		/*
		// To keep this example's output stable
		// even if there are changes in the testing package,
		// stop unwinding when we leave package runtime.
		if !strings.Contains(frame.File, "runtime/") {
			break
		}
		fmt.Printf("- more:%v | %s\n", more, frame.Function)
		*/

                prefix += " "  

		if !more {
			break
		}
	}

        return prefix
}

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
