package systemctl

import "os/exec"

func StartService(name string) error {
	cmd := exec.Command("systemctl", "start", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return NewError(string(out), err)
	}

	return nil
}

func StopService(name string) error {
	cmd := exec.Command("systemctl", "stop", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return NewError(string(out), err)
	}
	return nil
}

func EnableService(name string) error {
	cmd := exec.Command("systemctl", "enable", name)
	out, err := cmd.CombinedOutput()

	if err != nil {
		return NewError(string(out), err)
	}
	return nil
}

func DisableService(name string) error {
	cmd := exec.Command("systemctl", "disable", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return NewError(string(out), err)
	}

	return nil
}

func RestartService(name string) error {
	cmd := exec.Command("systemctl", "restart", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return NewError(string(out), err)
	}

	return nil
}

func ReloadService(name string) error {
	cmd := exec.Command("systemctl", "reload", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return NewError(string(out), err)
	}

	return nil
}

func StatusService(name string) error {
	cmd := exec.Command("systemctl", "status", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return NewError(string(out), err)
	}

	return nil
}
