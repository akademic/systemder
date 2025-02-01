package systemder

import (
	"os"
	"strings"
	"testing"
)

type writeMock struct {
	callsCount int
	calls      []struct {
		writtenPath string
		writtenData []byte
		writtenPerm os.FileMode
	}
}

func (w *writeMock) writeFile(path string, data []byte, perm os.FileMode) error {
	if w.calls == nil {
		w.calls = make([]struct {
			writtenPath string
			writtenData []byte
			writtenPerm os.FileMode
		}, 0)
	}

	call := struct {
		writtenPath string
		writtenData []byte
		writtenPerm os.FileMode
	}{
		writtenPath: path,
		writtenData: data,
		writtenPerm: perm,
	}

	w.calls = append(w.calls, call)
	w.callsCount++
	return nil
}

func TestGenerateAndWriteService(t *testing.T) {
	log := &NullLogger{}
	write := &writeMock{}
	s := &Systemder{
		log:       log,
		writeFile: write.writeFile,
	}

	desc := "Test service"
	name := "testservice"
	err := s.GenerateAndWriteService(desc, name)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if write.callsCount != 1 {
		t.Fatalf("unexpected calls count: %d", write.callsCount)
	}

	if write.calls[0].writtenPath != "/etc/systemd/system/"+name+".service" {
		t.Fatalf("unexpected written path: %s", write.calls[0].writtenPath)
	}

	if !strings.Contains(string(write.calls[0].writtenData), desc) {
		t.Fatalf("unexpected written data: %s", write.calls[0].writtenData)
	}

	binaryName := "systemder.test" // go test makes a binary with this name
	if !strings.Contains(string(write.calls[0].writtenData), binaryName) {
		t.Fatalf("unexpected written data: %s", write.calls[0].writtenData)
	}
}

func TestGenerateAndWriteTimer(t *testing.T) {
	log := &NullLogger{}
	write := &writeMock{}
	s := &Systemder{
		log:       log,
		writeFile: write.writeFile,
	}

	desc := "Test timer"
	onCalendar := "* * * * *"
	name := "testtimer"
	err := s.GenerateAndWriteTimer(desc, onCalendar, name)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if write.callsCount != 2 {
		t.Fatalf("unexpected calls count: %d", write.callsCount)
	}

	// oneshot
	if write.calls[0].writtenPath != "/etc/systemd/system/"+name+".service" {
		t.Fatalf("unexpected written path: %s", write.calls[0].writtenPath)
	}

	if !strings.Contains(string(write.calls[0].writtenData), desc) {
		t.Fatalf("unexpected written data: %s", write.calls[0].writtenData)
	}

	binaryName := "systemder.test" // go test makes a binary with this name
	if !strings.Contains(string(write.calls[0].writtenData), binaryName) {
		t.Fatalf("unexpected written data: %s", write.calls[0].writtenData)
	}

	// timer
	if write.calls[1].writtenPath != "/etc/systemd/system/"+name+".timer" {
		t.Fatalf("unexpected written path: %s", write.calls[1].writtenPath)
	}

	if !strings.Contains(string(write.calls[1].writtenData), desc) {
		t.Fatalf("unexpected written data: %s", write.calls[1].writtenData)
	}

	if !strings.Contains(string(write.calls[1].writtenData), name+".service") {
		t.Fatalf("unexpected written data: %s", write.calls[1].writtenData)
	}

	if !strings.Contains(string(write.calls[1].writtenData), onCalendar) {
		t.Fatalf("unexpected written data: %s", write.calls[1].writtenData)
	}

}
