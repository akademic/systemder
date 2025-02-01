package systemder

import (
	"path/filepath"
)

const defaultDir = "/etc/systemd/system"

// GenerateAndWriteService generates a service unit file and writes it to the default directory.
func (s *Systemder) GenerateAndWriteService(desc, name string) error {
	unit, err := s.GenerateService(desc)
	if err != nil {
		return err
	}

	serviceName := name + ".service"

	return s.writeUnit(unit, serviceName)
}

// GenerateAndWriteOneshot generates a oneshot unit file and writes it to the default directory.
func (s *Systemder) GenerateAndWriteTimer(desc, onCalendar, name string, args []string) error {
	oneshot, err := s.GenerateOneshot(desc, name, args)
	if err != nil {
		return err
	}

	oneshotName := name + ".service"

	timer, err := s.GenerateTimer(desc, oneshotName, onCalendar)
	if err != nil {
		return err
	}

	err = s.writeUnit(oneshot, oneshotName)
	if err != nil {
		return err
	}

	timerName := name + ".timer"

	return s.writeUnit(timer, timerName)
}

func (s *Systemder) writeUnit(unit string, name string) error {
	unitPath := filepath.Join(defaultDir, name)

	return s.writeFile(unitPath, []byte(unit), 0644)
}
