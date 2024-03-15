package process

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Process struct {
	User string
	PID  int

	CPU  float32
	MEM  float32
	VZS  int
	RSS  int
	TT   string
	Stat string

	Started string
	Time    string
	Command string
}

func NewProcess(line string) *Process {
	// ps aux
	f := strings.Fields(line)

	// User
	user := f[0]

	// PID
	pid, err := strconv.Atoi(f[1])
	if err != nil {
		log.Fatal(err)
	}

	// CPU
	rawCpu, err := strconv.ParseFloat(f[2], 32)
	if err != nil {
		log.Fatal(err)
	}

	cpu := float32(rawCpu)

	// MEM
	rawMem, err := strconv.ParseFloat(f[3], 32)
	if err != nil {
		log.Fatal(err)
	}

	mem := float32(rawMem)

	// VZS
	vzs, err := strconv.Atoi(f[4])
	if err != nil {
		log.Fatal(err)
	}

	// RSS
	rss, err := strconv.Atoi(f[5])
	if err != nil {
		log.Fatal(err)
	}

	// TT
	tt := f[6]

	// STAT
	stat := f[7]

	// STARTED
	started := f[8]

	// TIME
	time := f[9]

	// COMMAND
	command := f[10]

	return &Process{
		User:    user,
		PID:     pid,
		CPU:     cpu,
		MEM:     mem,
		VZS:     vzs,
		RSS:     rss,
		TT:      tt,
		Stat:    stat,
		Started: started,
		Time:    time,
		Command: command,
	}
}

func (p *Process) ToString() string {
	return fmt.Sprintf("[%s] %d mb\n----\n", p.Command, p.RSS/1024)
}
