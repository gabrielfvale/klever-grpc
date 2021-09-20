package test

import (
	"os"
	"testing"

	"github.com/gabrielfvale/klever-grpc/pkg"
)

// TestEnvEmpty tests if LoadEnv() loads an invalid .env file to the Environ
func TestEnvEmpty(t *testing.T) {
	enverr := pkg.LoadEnv()

	if _, err := os.Stat("../.env"); os.IsNotExist(err) {
		if enverr != nil {
			t.Error("Env does not exist")
		}
	}
}

// TestEnvVar tests if GetEnvVar() returns the proper value from Environ
func TestEnvVar(t *testing.T) {
	os.Setenv("TEST", "OK")

	res := pkg.GetEnvVar("TEST")
	if res != "OK" {
		t.Error("Env variable was wrongly read")
	}
}
