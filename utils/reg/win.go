//go:build windows

// Use registry for the author's experiences

package reg

import (
	"golang.org/x/sys/windows/registry"
	"os"
	"path/filepath"
)

// ReadFileDir reads filepath for the app from registry
func ReadFileDir() string {
	key, existed, err := registry.CreateKey(registry.CURRENT_USER, `Software\Rz\KMP Editor`, registry.ALL_ACCESS)
	if err != nil {
		errorDialog(err)
	}

	// If key is already existed
	if existed {
		value, _, err := key.GetBinaryValue("FilePath")
		if err != nil {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				errorDialog(err)
			}
			err = key.SetStringValue("FilePath", homeDir)
			if err != nil {
				errorDialog(err)
			}
			return homeDir
		}
		return string(value)
	}

	// If not
	homeDir, err := os.UserHomeDir()
	if err != nil {
		errorDialog(err)
	}
	err = key.SetStringValue("FilePath", homeDir)
	if err != nil {
		errorDialog(err)
	}

	return homeDir
}

// SetFileDir save the filepath to registry
func SetFileDir(dir string) error {
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
