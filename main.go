package systemder

import "os"

type Systemder struct {
	log       Logger
	writeFile func(path string, data []byte, perm os.FileMode) error
}

func NewSystemder(log Logger) *Systemder {
	return &Systemder{
		log:       log,
		writeFile: os.WriteFile,
	}
}
