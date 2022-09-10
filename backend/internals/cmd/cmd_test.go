package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	c "github.com/MrTimeout/go-home/backend/internals/config"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func backupConfigFile(t *testing.T, dir, name string) {
	f, err := os.OpenFile(filepath.Join(dir, name), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err == nil && len(b) >= 0 {
		os.Remove(f.Name()) //nolint:errcheck
		t.Cleanup(func() {
			os.WriteFile(f.Name(), b, fileInfo.Mode())
		})
	}
}

func createConfigFile(t *testing.T, dir, name, text string) *os.File {
	f, err := os.OpenFile(filepath.Join(dir, name), os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if b, err := io.ReadAll(f); err == nil && len(b) > 0 {
		if err := os.Remove(f.Name()); err == nil {
			createConfigFile(t, dir, name, text)
		}
	}

	if strings.TrimSpace(text) != "" {
		f.WriteString(text)
	}

	if b, err := io.ReadAll(f); err == nil {
		t.Log("this is the string: ", b)
	}

	t.Cleanup(func() {
		os.Remove(f.Name())
	})

	return f
}

func TestReadConfig(t *testing.T) {
	var configFile = ConfigFileName + "." + ConfigFileType
	var config = c.Config{
		Logger: c.Logger{
			Production: false,
			FileAppenders: []c.FileLoggerAppender{
				{
					LoggerFileLevel: c.DebugLevel,
					LoggerFileName:  "/tmp/go-home.json",
					DateTimeFormat:  c.RFC3339,
				},
			},
		},
	}

	want, err := yaml.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("config file at home directory successfully readed", func(t *testing.T) {
		backupConfigFile(t, pwd, configFile)
		backupConfigFile(t, home, configFile)

		_ = createConfigFile(t, home, configFile, string(want))

		assert.Nil(t, Execute())
		assert.Equal(t, config, cfg)
	})

	t.Run("config file at current directory successfully readed", func(t *testing.T) {
		backupConfigFile(t, pwd, configFile)
		backupConfigFile(t, home, configFile)

		_ = createConfigFile(t, pwd, configFile, string(want))

		assert.Nil(t, Execute())
		assert.Equal(t, config, cfg)
	})

	t.Run("there is no config file, so it should return err checking config file", func(t *testing.T) {
		backupConfigFile(t, pwd, configFile)
		backupConfigFile(t, home, configFile)

		assert.ErrorIs(t, ErrCheckConfigFile, checkConfigFile())
	})
}
