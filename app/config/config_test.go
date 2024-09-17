package config

import (
	"ekolo/pkg/assert"
	"os"
	"testing"
)

var env_vars = map[string]string{
	envHTTP:   ":8080",
	envDBHost: "db.koko.com",
	envDBPort: "5432",
	envDBName: "koko",
	envDBUser: "koko",
	envDBPass: "kokopwd",
}

func TestConfig(t *testing.T) {
	for key, val := range env_vars {
		os.Setenv(key, val)
	}

	cf := New()

	assert.Assert(t, cf.HTTPAddr, env_vars["EKOLO_HTTP"])
	assert.Assert(t, cf.DBHost, env_vars["EKOLO_DB_HOST"])
	assert.Assert(t, cf.DBPort, env_vars["EKOLO_DB_PORT"])
	assert.Assert(t, cf.DBName, env_vars["EKOLO_DB_NAME"])
	assert.Assert(t, cf.DBUser, env_vars["EKOLO_DB_USER"])
	assert.Assert(t, cf.DBPass, env_vars["EKOLO_DB_PASS"])
}
