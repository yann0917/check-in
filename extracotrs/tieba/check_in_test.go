package tieba

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

func TestGetForumList(t *testing.T) {
	tbs, list := GetForumList()
	t.Log(tbs)
	t.Log(list)
}

func TestSignAdd(t *testing.T) {
	SignAdd()
}

func TestGetTbs(t *testing.T) {
	tbs := GetTbs()
	t.Log(tbs)
}

func TestCheckIn(t *testing.T) {
	tbs := GetTbs()
	OneKeySignIn(tbs.Tbs)
}
