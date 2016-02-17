package config

import (
	"github.com/magiconair/properties"
)

func FromFile(filename string) (*Config, error) {
	p, err := properties.LoadFile(filename, properties.UTF8)
	if err != nil {
		return nil, err
	}
	return FromProperties(p)
}

func FromProperties(p *properties.Properties) (*Config, error) {
	var cfg *Config = &Config{}

	cfg.Listen.Port = p.GetString("receiver.addr", DefaultConfig.Listen.Port)

	return cfg, nil
}
