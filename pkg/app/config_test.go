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

	if c.JWT.Issuer != "go-web" {
		t.Errorf("JWT Issuer was incorrect: %s", c.JWT.Issuer)
	}

	if c.JWT.Secret != "m8vh34h83u4v9h94v9v23lpd08n5h39d" {
		t.Errorf("JWT Secret was incorrect: %s", c.JWT.Secret)
	}

	if c.JWT.ClientKey != "93mf8v98h5rnegh7849hf3uwrgb93h20hgety" {
		t.Errorf("JWT ClientKey was incorrect: %s", c.JWT.ClientKey)
	}
}
