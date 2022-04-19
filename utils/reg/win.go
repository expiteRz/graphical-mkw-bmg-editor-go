//go:build windows

package reg

import (
	"github.com/sqweek/dialog"
	"golang.org/x/sys/windows/registry"
	"os"
	"path/filepath"
)

// ReadFileDir reads filepath for the app from registry
func ReadFileDir() string {
	exitFunc := func(err error) {
		dialog.Message("%v\n\nPress OK to exit the application", err).Title("Fatal error").Info()
		os.Exit(1)
	}

	key, existed, err := registry.CreateKey(registry.CURRENT_USER, `Software\Rz\KMP Editor`, registry.ALL_ACCESS)
	if err != nil {
		exitFunc(err)
	}

	// If key is already existed
	if existed {
		value, _, err := key.GetBinaryValue("FilePath")
		if err != nil {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				exitFunc(err)
			}
			err = key.SetStringValue("FilePath", homeDir)
			if err != nil {
				exitFunc(err)
			}
			return homeDir
		}
		return string(value)
	}

	// If not
	homeDir, err := os.UserHomeDir()
	if err != nil {
		exitFunc(err)
	}
	err = key.SetStringValue("FilePath", homeDir)
	if err != nil {
		exitFunc(err)
	}

	return homeDir
}

// SetDirToRegistry save the filepath to registry
func SetDirToRegistry(dir string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Rz\KMP Editor`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	dirc, _ := filepath.Split(dir)
	if err = key.SetStringValue("FilePath", dirc); err != nil {
		return err
	}

	return nil
}
