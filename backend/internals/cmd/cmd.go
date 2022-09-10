package cmd

import (
	"errors"
	"os"
	"path/filepath"

	c "github.com/MrTimeout/go-home/backend/internals/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	ConfigFileType = "yaml"
	ConfigFileName = "go-home"
)

var (
	// ErrCheckConfigFile is used to inform that the config file was not
	// readed successfully. File maybe doesn't exist or user has not the rights.
	ErrCheckConfigFile = errors.New("checking config file")

	// TODO: we have to fix this global variable, we can't have a global variable to the configuration
	// It is not well encapsulated.
	cfg c.Config
)

// NewRootCmd is the main entrypoint of the application. When the program
// starts executing, it will trigger all config files and parameters needed
// to get the job done
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "go-home",
		Short: "Just the main entrypoint to execute go-home API",
		Long:  "Just the main entrypoint to execute go-home API",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

func readConfig() {
	if err := checkConfigFile(); err != nil {
		panic(err)
	}

	viper.SetConfigType(ConfigFileType)
	viper.SetConfigName(ConfigFileName)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	c.ConfigureLogger(cfg.Logger)
	c.ConfigureDB(cfg.Database)
}

func checkConfigFile() error {
	if pwd, err := os.Getwd(); err == nil && fileExists(pwd) {
		viper.AddConfigPath(pwd)
		return nil
	}

	if home, err := os.UserHomeDir(); err == nil && fileExists(home) {
		viper.AddConfigPath(home)
		return nil
	}

	return ErrCheckConfigFile
}

func fileExists(dir string) bool {
	_, err := os.OpenFile(filepath.Join(dir, ConfigFileName+"."+ConfigFileType), os.O_RDONLY, 0444)
	return err == nil
}

// Execute is the method called by main file to start the application.
func Execute() error {
	cobra.OnInitialize(readConfig)

	rootCmd := NewRootCmd()

	return rootCmd.Execute()
}
