package bash

import (
	"bytes"
	"os/exec"
	"strings"
	"syscall"
)

const defaultFailedCode = 1

type Out struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

// RunTest case by command
// output will be as struct
// stdout, stderr, exitcode
func RunTest(command string) Out {
	var outbuf, errbuf bytes.Buffer
	var exitCode int

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout := strings.Replace(outbuf.String(), "\n", "", -1)
	stderr := strings.Replace(errbuf.String(), "\n", "", -1)

	if err != nil {
		// try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			exitCode = defaultFailedCode
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}

	result := Out{
		StdOut:   stdout,
		StdErr:   stderr,
		ExitCode: exitCode,
	}
	return result
}

// CustomSetEnv execute value before set
//func CustomSetEnv(key string, value string) {
//	logger := logging.MainLog().Logger
//	value = RunTest(value).StdOut
//	err := os.Setenv(key, value)
//	if err != nil {
//		logger.Warn().Str("var", key).Msg("failed when set env variable")
//	}
//}
