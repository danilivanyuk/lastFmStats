package tools

import (
	log "github.com/sirupsen/logrus"
)

type UserDetails struct {
	Username string
}
type SongsDetails struct {
	AmountOfSongs int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *UserDetails
	GetSongDetails(username string) *SongsDetails
	SetupDatabase() error
}
