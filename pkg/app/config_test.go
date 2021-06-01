package app

import "testing"

func TestLoadConfig(t *testing.T) {
	c, err := LoadConfig("../../config.yaml")
	if err != nil {
		t.Error(err)
	}

	if c.AppName != "Go-Web" {
		t.Errorf("AppName was incorrect: %s", c.AppName)
	}
}

func TestServerConfig(t *testing.T) {
	c, err := LoadConfig("../../config.yaml")
	if err != nil {
		t.Error(err)
	}

	if c.Server.Port != 3000 {
		t.Errorf("Server Port was incorrect: %d", c.Server.Port)
	}

	if c.Server.JWT.Issuer != "go-web-api" {
		t.Errorf("Server JWT Issuer was incorrect: %s", c.Server.JWT.Issuer)
	}

	if c.Server.JWT.Secret != "secret-jwt-key" {
		t.Errorf("Server JWT Secret was incorrect: %s", c.Server.JWT.Secret)
	}
}
