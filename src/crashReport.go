package main

/*
Crash-Reporter taken from GitHubs hub: https://github.com/github/hub/blob/master/github/crash_report.go
*/

/*
func captureCrash() {
	if rec := recover(); rec != nil {
		if err, ok := rec.(error); ok {
			reportCrash(err)
		} else if err, ok := rec.(string); ok {
			reportCrash(errors.New(err))
		}
	}
}

func reportCrash(err error) {
	if err == nil {
		return
	}

	buf := make([]byte, 10000)
	runtime.Stack(buf, false)
	stack := formatStack(buf)
	errType := reflect.TypeOf(err).String()

	fmt.Println(errType)
	fmt.Printf("%v\n\n", err)
	fmt.Println(stack)

	os.Exit(1)
}

func runtimeInfo() string {
	return fmt.Sprintf("GOOS: %s\nGOARCH: %s", runtime.GOOS, runtime.GOARCH)
}

func formatStack(buf []byte) string {
	buf = bytes.Trim(buf, "\x00")

	stack := strings.Split(string(buf), "\n")
	stack = append(stack[0:1], stack[5:]...)

	return strings.Join(stack, "\n")
}
*/
