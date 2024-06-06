package config

import (
	"bytes"
	"strconv"
)

func writeDSNParam(buf *bytes.Buffer, hasParam *bool, name, value string) {
	buf.Grow(1 + len(name) + 1 + len(value))
	if !*hasParam {
		*hasParam = true
		buf.WriteByte('?')
	} else {
		buf.WriteByte('&')
	}
	buf.WriteString(name)
	buf.WriteByte('=')
	buf.WriteString(value)
}

func (cfg *PGConfig) FormatDSN() string {
	var buf bytes.Buffer
	// postgres://
	buf.WriteString("postgres://")

	// [username[:password]@]
	if len(cfg.DBUser) > 0 {
		buf.WriteString(cfg.DBUser)
		if len(cfg.DBPassword) > 0 {
			buf.WriteByte(':')
			buf.WriteString(cfg.DBPassword)
		}
		buf.WriteByte('@')
	}

	// host:port
	buf.WriteString(cfg.DBHost)
	buf.WriteByte(':')
	buf.WriteString(cfg.DBPort)

	// .../dbname
	buf.WriteByte('/')
	buf.WriteString(cfg.DBName)

	// [?param1=value1&...&paramN=valueN]
	hasParam := false

	// Schema
	if len(cfg.DBSchema) > 0 {
		writeDSNParam(&buf, &hasParam, "search_path", cfg.DBSchema)
	}

	if len(cfg.DBSSLMode) > 0 {
		writeDSNParam(&buf, &hasParam, "sslmode", cfg.DBSSLMode)
	}

	if cfg.DBConnTimeout > 0 {
		writeDSNParam(&buf, &hasParam, "connect_timeout", strconv.FormatInt(cfg.DBConnTimeout, 10))
	}

	return buf.String()
}
