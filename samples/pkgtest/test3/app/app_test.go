package app_test

import (
	"pkgtest3/app/mock_app"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAppEcho(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := mock_app.NewMockApplication(ctrl)
	app.EXPECT().Name().Return("MockApp")

	// mock を利用したテストを書く
	if got := app.Name(); got != "MockApp" {
		t.Fatalf("Name() = %s, want %s", got, "MockApp")
	}
}
t