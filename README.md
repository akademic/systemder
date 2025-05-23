# Systemder

A Go library for generating and managing systemd service and timer units programmatically.

## Features

- Generate systemd service units from your Go applications
- Create oneshot services with timers for scheduled tasks
- Automatically setup and manage systemd services
- Simple API for common systemd operations

## Installation

```bash
go get github.com/akademic/systemder
```

## Quick Start

### Basic Service Setup

```go
package main

import (
    "log"
    "github.com/akademic/systemder"
)

func main() {
    // Create a logger (or use NullLogger for no logging)
    logger := &systemder.NullLogger{}

    // Create systemder instance
    sys := systemder.NewSystemder(logger)

    // Setup a service that will run your current executable
    err := sys.SetupService("My Go Application", "myapp")
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Service 'myapp.service' created and started!")
}
```

### Scheduled Tasks with Timers

```go
package main

import (
    "log"
    "github.com/akademic/systemder"
)

func main() {
    logger := &systemder.NullLogger{}
    sys := systemder.NewSystemder(logger)

    // Create a timer that runs every hour with arguments
    args := []string{"--cleanup", "--verbose"}
    err := sys.SetupTimer(
        "Hourly cleanup task",
        "hourly",           // OnCalendar specification
        "cleanup-task",     // Service name
        args,              // Arguments to pass
    )
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Timer 'cleanup-task.timer' created and started!")
}
```

## API Reference

### Methods

#### SetupService

Creates a systemd service unit, enables it, and starts it.

```go
func (s *Systemder) SetupService(desc, name string) error
```

**Parameters:**
- `desc`: Description for the service
- `name`: Name of the service (without .service extension)

**Example:**
```go
err := sys.SetupService("My web server", "webserver")
// Creates: webserver.service
```

#### SetupTimer

Creates a systemd timer with associated oneshot service, enables it, and starts it.

```go
func (s *Systemder) SetupTimer(desc, onCalendar, name string, args []string) error
```

**Parameters:**
- `desc`: Description for the timer and service
- `onCalendar`: Systemd calendar specification (e.g., "daily", "hourly", "Mon *-*-* 00:00:00")
- `name`: Base name for the timer and service
- `args`: Command line arguments to pass to the executable

**Example:**
```go
// Run daily at midnight
err := sys.SetupTimer("Daily backup", "daily", "backup", []string{"--full"})
// Creates: backup.service and backup.timer

// Run every 5 minutes
err := sys.SetupTimer("Frequent check", "*:0/5", "healthcheck", nil)
```

#### GenerateService

Generates a systemd service unit content as a string.

```go
func (s *Systemder) GenerateService(desc string) (string, error)
```

#### GenerateOneshot

Generates a oneshot service unit content.

```go
func (s *Systemder) GenerateOneshot(desc, name string, args []string) (string, error)
```

#### GenerateTimer

Generates a timer unit content.

```go
func (s *Systemder) GenerateTimer(desc string, service string, onCalendar string) (string, error)
```

#### GenerateAndWriteService

Generates and writes a service unit file to `/etc/systemd/system/`.

```go
func (s *Systemder) GenerateAndWriteService(desc, name string) error
```

#### GenerateAndWriteTimer

Generates and writes both oneshot service and timer unit files.

```go
func (s *Systemder) GenerateAndWriteTimer(desc, onCalendar, name string, args []string) error
```

## Calendar Specifications

Common `onCalendar` values for timers:

- `"minutely"` - Every minute
- `"hourly"` - Every hour
- `"daily"` - Every day at midnight
- `"weekly"` - Every week
- `"monthly"` - Every month
- `"*:0/5"` - Every 5 minutes
- `"Mon,Tue *-*-* 10:00:00"` - Monday and Tuesday at 10 AM
- `"*-*-* 06:00:00"` - Every day at 6 AM

## Advanced Usage

### Manual Unit Generation

```go
sys := systemder.NewSystemder(&systemder.NullLogger{})

// Just generate the unit content without writing
serviceUnit, err := sys.GenerateService("My service description")
if err != nil {
    log.Fatal(err)
}

fmt.Println(serviceUnit)
// Output: systemd service unit content
```

## Requirements

- Linux system with systemd
- Root privileges (for writing to `/etc/systemd/system/`)

## Notes

- All unit files are written to `/etc/systemd/system/`
- Services use the current executable path automatically
- Working directory is set to the executable's directory

