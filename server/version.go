package server

import "fmt"

type VersionInfo struct {
	Version    string
	CommitHash string
}

var (
	AppName    = "react-go-template" // Name of the application
	Version    = "<not-set>"         // Version of the application
	CommitHash = "<not-set>"         // Commit hash of the application
)

func (v VersionInfo) String() string {
	return fmt.Sprintf("%s (%s)", v.Version, v.CommitHash)
}
