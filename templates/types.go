package templates

type RestartType string

const (
	RestartAlways    RestartType = "always"
	RestartOnFailure RestartType = "on-failure"
	RestartNo        RestartType = "no"
)
