package systemder

import (
	"os"
	"path/filepath"

	"github.com/akademic/systemder/templates"
)

func (s *Systemder) GenerateService(desc string) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	service := templates.Service{
		Description:      desc,
		ExecStart:        ex,
		WorkingDirectory: filepath.Dir(ex),
	}

	return templates.GetServiceUnit(service), nil
}

func (s *Systemder) GenerateOneshot(desc, name string) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	oneshot := templates.Oneshot{
		Description:      desc,
		Wants:            name + ".timer",
		ExecStart:        ex,
		WorkingDirectory: filepath.Dir(ex),
	}

	return templates.GetOneshotUnit(oneshot), nil
}

func (s *Systemder) GenerateTimer(desc string, service string, onCalendar string) (string, error) {
	timer := templates.Timer{
		Description: desc,
		Service:     service,
		OnCalendar:  onCalendar,
	}

	return templates.GetTimerUnit(timer), nil
}
