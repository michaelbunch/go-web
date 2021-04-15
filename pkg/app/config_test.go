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
