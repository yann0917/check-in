package juejin

import (
	"os"
	"testing"

	"github.com/yann0917/check-in/pkg"
)

func TestMain(m *testing.M) {
	pkg.Viper("../../config.yaml")

	code := m.Run()
	os.Exit(code)
}

func TestGetCheckInStatus(t *testing.T) {
	GetCheckInStatus()
}
