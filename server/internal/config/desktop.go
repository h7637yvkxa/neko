package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Desktop holds configuration for the virtual desktop environment.
type Desktop struct {
	Display     string
	DisplayDepth int
	ScreenWidth  int
	ScreenHeight int
	ScreenRate   int
}

// Server returns the cobra command for the desktop configuration.
func (Desktop) Init(cmd *cobra.Command) error {
	cmd.PersistentFlags().String("desktop.display", ":99.0", "X display to use for the virtual desktop")
	if err := viper.BindPFlag("desktop.display", cmd.PersistentFlags().Lookup("desktop.display")); err != nil {
		return err
	}

	cmd.PersistentFlags().Int("desktop.depth", 24, "color depth of the virtual desktop")
	if err := viper.BindPFlag("desktop.depth", cmd.PersistentFlags().Lookup("desktop.depth")); err != nil {
		return err
	}

	// Default to 1920x1080 for a more usable desktop experience
	cmd.PersistentFlags().Int("desktop.width", 1920, "width of the virtual desktop in pixels")
	if err := viper.BindPFlag("desktop.width", cmd.PersistentFlags().Lookup("desktop.width")); err != nil {
		return err
	}

	cmd.PersistentFlags().Int("desktop.height", 1080, "height of the virtual desktop in pixels")
	if err := viper.BindPFlag("desktop.height", cmd.PersistentFlags().Lookup("desktop.height")); err != nil {
		return err
	}

	// Bumped default rate to 60Hz for smoother interaction
	cmd.PersistentFlags().Int("desktop.rate", 60, "refresh rate of the virtual desktop in Hz")
	if err := viper.BindPFlag("desktop.rate", cmd.PersistentFlags().Lookup("desktop.rate")); err != nil {
		return err
	}

	return nil
}

// Set populates the Desktop struct from viper configuration values.
func (s *Desktop) Set() {
	s.Display = viper.GetString("desktop.display")
	s.DisplayDepth = viper.GetInt("desktop.depth")
	s.ScreenWidth = viper.GetInt("desktop.width")
	s.ScreenHeight = viper.GetInt("desktop.height")
	s.ScreenRate = viper.GetInt("desktop.rate")
}
