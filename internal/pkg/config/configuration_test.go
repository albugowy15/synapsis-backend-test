package config_test

import (
	"testing"

	"github.com/albugowy15/synapsis-backend-test/internal/pkg/config"
)

func TestLoadConfig(t *testing.T) {
	config.LoadConfig("../../../")
	conf := config.GetConfig()

	if conf == nil {
		t.Fatal("Configuration is nil")
	}
	if conf.DBDriver != "postgres" {
		t.Errorf("Expected DbDriver config to be postgres, got %s", conf.DBDriver)
	}
	if conf.DBSource != "postgres://postgres:postgres@localhost:5432/synapsis_db" {
		t.Errorf("Expected DBSource config to be postgres://postgres:postgres@localhost:5432/synapsis_db, got %s", conf.DBSource)
	}
	if conf.Port != "8080" {
		t.Errorf("Expected Port config to be 8080, got %s", conf.Port)
	}
}
