package testing

import (
	"github.com/refaldyrk/gego/gego"
	"testing"
)

func TestResult(t *testing.T) {
	result := gego.Result("ini adalah kata rahasia")

	if result != "iginigi agadagalagah kagataga ragahagasigiaga" {
		t.Errorf("[ERROR]: Result got %s", result)
	} else {
		t.Logf("[SUCCESS]: Result got %s", result)
	}
}
