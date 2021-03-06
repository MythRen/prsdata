package main

import (
	"fmt"
	"sync"
	"time"

	logger "github.com/sirupsen/logrus"
)

type RealCommand struct {
	round   int
	total   int
	command *Command
	pcap    *Pcap
	realJob *RealJob
	g       *sync.WaitGroup
}

func (r *RealCommand) String() string {
	if r.command.Type == "shell" {
		return fmt.Sprintf("%s [%d/%d] %s", r.realJob, r.round, r.total, r.command)
	} else {
		return fmt.Sprintf("%s [%d/%d] %s %s", r.realJob, r.round, r.total, r.command, r.pcap)
	}
}

func (r *RealCommand) run() {
	logger.Infoln(fmt.Sprintf("%s executing", r))
	start := time.Now()
	result := execRealCommand(r)
	end := time.Now()
	duration := end.Sub(start)

	if !result.succeed {
		err := result.err.Error()
		reason := result.output
		if reason == "" {
			reason = "<empty output>"
		}
		logger.Errorln(fmt.Sprintf("%s execute failed, use: %s, err is: %s, output is:\n%s\n", r, duration, err, reason))
	} else {
		logger.Infoln(fmt.Sprintf("%s execute succeed, use: %s", r, duration))
	}
}
