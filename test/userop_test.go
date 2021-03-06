package test

import (
	"GoEchoton/repository"
	"testing"
)

func TestUserOPVaild(t *testing.T) {
	userop := repository.NewUserOP()
	r, err := userop.Valid("jon", "hahha")
	if err != nil {
		t.Fatal(err)
	}
	if !r {
		t.Fatal("Error data")
	}
}
