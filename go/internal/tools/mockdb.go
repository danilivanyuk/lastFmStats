package tools

import (
	"time"
)

type mockDB struct {}

var mockUserDetails = map[string]UserDetails{
	"Max": {
		Username: "Max",
	},
	"Tylor": {
		Username: "Tylor",
	},
	"XX": {
		Username: "XX",
	},
}
var mockSongsDetails = map[string]SongsDetails{
	"Max": {
		Username: "Max",
		AmountOfSongs: 120,
	},
	"Tylor": {
		Username: "Tylor",
		AmountOfSongs: 53,
	},
	"XX": {
		Username: "XX",
		AmountOfSongs: 84,
	},
}

func (d *mockDB) GetUserDetails(username string) *UserDetails {
}
func (d *mockDB) GetSongsDetails(username string) *SongsDetails {
}
func (d *mockDB) SetupDatabase() error {
	return nil
}

