package systemder

import (
	"github.com/akademic/systemder/systemctl"
)

// SetupService generates a service unit file, writes it to the default directory and starts the service.
func (s *Systemder) SetupService(desc, name string) error {
	err := s.GenerateAndWriteService(desc, name)
	if err != nil {
		return err
	}

	unitName := name + ".service"

	err = systemctl.EnableService(unitName)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			dErr := systemctl.DisableService(unitName)
			if dErr != nil {
				s.log.Error("failed to disable service: %v", err)
			}
		}
	}()

	err = systemctl.StartService(unitName)
	if err != nil {
		return err
	}

	return nil
}

func (s *Systemder) SetupTimer(desc, onCalendar, name string, args []string) error {
	err := s.GenerateAndWriteTimer(desc, onCalendar, name, args)
	if err != nil {
		return err
	}

	unitName := name + ".timer"

	err = systemctl.EnableService(unitName)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			dErr := systemctl.DisableService(unitName)
			if dErr != nil {
				s.log.Error("failed to disable service: %v", err)
			}
		}
	}()

	err = systemctl.StartService(unitName)
	if err != nil {
		return err
	}

	return nil
}
