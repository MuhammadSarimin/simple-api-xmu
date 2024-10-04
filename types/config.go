package types

import "strings"

type Config struct {
	Name string `env:"APP_NAME"`
	Host string `env:"APP_HOST" envDefault:"0.0.0.0"`
	Port string `env:"APP_PORT" envDefault:"9999"`
	DB   DBConfig
}

type DBConfig struct {
	Host        string `env:"DB_HOST"`
	Port        string `env:"DB_PORT"`
	User        string `env:"DB_USER"`
	Pass        string `env:"DB_PASS"`
	Name        string `env:"DB_NAME"`
	SSL         string `env:"DB_SSL" envDefault:"disable"`
	AutoMigrate bool   `env:"DB_AUTOMIGRATE" envDefault:"true"`
}

func (c *Config) Address() string {
	return c.Host + ":" + c.Port
}

func (c *DBConfig) DSN() string {
	var s strings.Builder

	if c.Host != "" {
		c.writeParam(&s, "host", c.Host)
	}

	if c.Port != "" {
		c.writeParam(&s, "port", c.Port)
	}

	if c.User != "" {
		c.writeParam(&s, "user", c.User)
	}

	if c.Pass != "" {
		c.writeParam(&s, "password", c.Pass)
	}

	if c.Name != "" {
		c.writeParam(&s, "dbname", c.Name)
	}

	if c.SSL == "" {
		c.SSL = "disable"
	}

	c.writeParam(&s, "sslmode", c.SSL)

	return s.String()
}

func (c *DBConfig) writeParam(s *strings.Builder, key, value string) {
	s.WriteByte(' ')
	s.WriteString(key)
	s.WriteByte('=')
	s.WriteString(value)
}
