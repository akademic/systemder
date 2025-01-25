package templates

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed units/service.unit
var serviceTemplate string

//go:embed units/oneshot.unit
var serviceOneshot string

//go:embed units/timer.unit
var serviceTimer string

type Service struct {
	Description      string
	Restart          RestartType
	RestartSec       int
	ExecStart        string // command to run
	WorkingDirectory string // directory where the ExecStart command is run
}

type Oneshot struct {
	Description      string
	Wants            string // timer unit file name
	ExecStart        string // command to run
	WorkingDirectory string // directory where the ExecStart command is run
}

type Timer struct {
	Description string
	OnCalendar  string
	Service     string // service unit file name
}

func GetServiceUnit(s Service) string {
	templ := template.Must(template.New("name").Parse(serviceTemplate))

	buf := new(bytes.Buffer)

	s = fillWithDefaults(s)

	err := templ.Execute(buf, s)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func GetOneshotUnit(s Oneshot) string {
	templ := template.Must(template.New("name").Parse(serviceOneshot))

	buf := new(bytes.Buffer)

	err := templ.Execute(buf, s)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func GetTimerUnit(t Timer) string {
	templ := template.Must(template.New("name").Parse(serviceTimer))

	buf := new(bytes.Buffer)

	err := templ.Execute(buf, t)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func fillWithDefaults(s Service) Service {
	if s.Restart == "" {
		s.Restart = "always"
	}

	if s.RestartSec == 0 {
		s.RestartSec = 5
	}

	return s
}
