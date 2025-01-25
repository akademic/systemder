package templates

import (
	"testing"
)

func TestGetServiceUnit(t *testing.T) {
	tests := []struct {
		name    string
		service Service
		want    string
	}{
		{
			name: "simple",
			service: Service{
				Description:      "test service",
				Restart:          RestartAlways,
				RestartSec:       10,
				ExecStart:        "/usr/bin/echo hello",
				WorkingDirectory: "/tmp",
			},
			want: "[Unit]\nDescription=test service\nRequires=network-online.target\nAfter=network-online.target\n\n[Service]\nRestart=always\nRestartSec=10\nExecStart=/usr/bin/echo hello\nWorkingDirectory=/tmp\n\n[Install]\nWantedBy=multi-user.target\n",
		},
		{
			name: "defaults",
			service: Service{
				Description:      "test service",
				ExecStart:        "/usr/bin/echo hello",
				WorkingDirectory: "/",
			},
			want: "[Unit]\nDescription=test service\nRequires=network-online.target\nAfter=network-online.target\n\n[Service]\nRestart=always\nRestartSec=5\nExecStart=/usr/bin/echo hello\nWorkingDirectory=/\n\n[Install]\nWantedBy=multi-user.target\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServiceUnit(tt.service); got != tt.want {
				t.Errorf("GetServiceUnit() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestGetOneshotUnit(t *testing.T) {
	tests := []struct {
		name    string
		service Oneshot
		want    string
	}{
		{
			name: "simple",
			service: Oneshot{
				Description:      "test service",
				ExecStart:        "/usr/bin/echo hello",
				WorkingDirectory: "/tmp",
				Wants:            "test_service.timer",
			},
			want: "[Unit]\nDescription=test service\nWants=test_service.timer\n\n[Service]\nType=oneshot\nExecStart=/usr/bin/echo hello\nWorkingDirectory=/tmp\n\n[Install]\nWantedBy=multi-user.target\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOneshotUnit(tt.service); got != tt.want {
				t.Errorf("GetOneshotUnit() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestGetTimerUnit(t *testing.T) {
	tests := []struct {
		name  string
		timer Timer
		want  string
	}{
		{
			name: "simple",
			timer: Timer{
				Description: "test timer",
				OnCalendar:  "*-*-* *:*:00",
				Service:     "test_service.service",
			},
			want: "[Unit]\nDescription=test timer\nRequires=test_service.service\n\n[Timer]\nUnit=test_service.service\nOnCalendar=*-*-* *:*:00\n\n[Install]\nWantedBy=timers.target\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTimerUnit(tt.timer); got != tt.want {
				t.Errorf("GetTimerUnit() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
