//go:build !windows

// Unix has not supported Registry, so use environment variables instead

package reg

import (
	"os"
	"path/filepath"
)

const envName = "RZ_BMG_EDITOR_OPENDIR"

// ReadFileDir reads filepath for the app from environment variables
func ReadFileDir() string {
	str, b := os.LookupEnv(envName)
	if !b {
		home, err := os.UserHomeDir()
		if err != nil {
			errorDialog(err)
			return ""
		}
		if err = os.Setenv(envName, home); err != nil {
			errorDialog(err)
			return ""
		}

		return home
	}

	return str
}

// SetFileDir save the filepath to environment variables
func SetFileDir(dir string) error {
	dirc, _ := filepath.Split(dir)
	if err := os.Setenv(envName, dirc); err != nil {
		return err
	}

	return nil
}
