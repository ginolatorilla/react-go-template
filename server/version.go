package server

import "fmt"

type VersionInfo struct {
	Version    string
	CommitHash string
}

func (v VersionInfo) String() string {
	return fmt.Sprintf("%s (%s)", v.Version, v.CommitHash)
}
