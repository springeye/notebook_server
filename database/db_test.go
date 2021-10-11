package database

import "testing"

func TestUser(t *testing.T) {
	Database.Explain("select version();")
}
