package systemctl

import "os/exec"

func StartService(name string) error {
	cmd := exec.Command("systemctl", "start", name)
	out, err := cmd.CombinedOutput()
	return NewError(string(out), err)
}

func StopService(name string) error {
	cmd := exec.Command("systemctl", "stop", name)
	out, err := cmd.CombinedOutput()
	return NewError(string(out), err)
}

func EnableService(name string) error {
	cmd := exec.Command("systemctl", "enable", name)
	_, err := cmd.CombinedOutput()
	return err
}

func DisableService(name string) error {
	cmd := exec.Command("systemctl", "disable", name)
	out, err := cmd.CombinedOutput()
	return NewError(string(out), err)
}

func RestartService(name string) error {
	cmd := exec.Command("systemctl", "restart", name)
	out, err := cmd.CombinedOutput()
	return NewError(string(out), err)
}

func ReloadService(name string) error {
	cmd := exec.Command("systemctl", "reload", name)
	out, err := cmd.CombinedOutput()
	return NewError(string(out), err)
}

func StatusService(name string) error {
	cmd := exec.Command("systemctl", "status", name)
	out, err := cmd.CombinedOutput()
	return NewError(string(out), err)
}
