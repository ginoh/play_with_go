package app

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// テスト全体の共通セットアップ/後始末を書く場所
	println("before m.Run")

	// ここで各テストが実行される
	code := m.Run()

	println("after m.Run")

	os.Exit(code)
}

func TestAppEcho(t *testing.T) {
	expect := "App"
	actual := AppEcho()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
