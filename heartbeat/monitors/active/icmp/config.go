package icmp

import (
	"time"

	"github.com/elastic/beats/heartbeat/monitors"
)

type Config struct {
	Name string `config:"name"`

	Mode      monitors.IPSettings `config:",inline"`
	Interface map[string][]string `config:"interface" validate:"required"`

	Timeout time.Duration `config:"timeout"`
	Wait    time.Duration `config:"wait"`
}

var DefaultConfig = Config{
	Name: "icmp",
	Mode: monitors.DefaultIPSettings,

	Timeout: 16 * time.Second,
	Wait:    1 * time.Second,
}
