package server

import "fmt"

// VersionInfo contains the version and commit hash of the application.
type VersionInfo struct {
	Version    string
	CommitHash string
}

var (
	AppName    = "react-go-template" // Name of the application, which should be set with linker flags.
	Version    = "<not-set>"         // Version of the application, which should be set with linker flags.
	CommitHash = "<not-set>"         // Commit hash of the application, which should be set with linker flags.
)

func (v VersionInfo) String() string {
	return fmt.Sprintf("%s (%s)", v.Version, v.CommitHash)
}
