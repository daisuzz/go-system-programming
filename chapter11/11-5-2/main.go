package main

import (
	"github.com/google/shlex"
	"strings"
)

func parseCmd(cmdStr string) (cmd string, args []string, err error) {
	l := shlex.NewLexer(strings.NewReader(cmdStr))
	cmd, err = l.Next()
	if err != nil {
		return
	}
	for token, err := l.Next(); err != nil; {
		args = append(args, token)
	}
	return
}
