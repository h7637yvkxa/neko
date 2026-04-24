package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Server holds configuration for the HTTP server.
type Server struct {
	Host string
	Port int
	Cert string
	Key  string

	// CORS
	CORSAllowOrigins []string

	// Static file serving
	StaticDir string

	// Path prefix for all routes
	PathPrefix string
}

// Bind registers server-related flags on the given cobra command
// and binds them to viper for configuration file / env var support.
func (s *Server) Bind(cmd *cobra.Command) error {
	cmd.PersistentFlags().String("server.host", "", "host to bind the server to (empty means all interfaces)")
	if err := viper.BindPFlag("server.host", cmd.PersistentFlags().Lookup("server.host")); err != nil {
		return err
	}

	// Changed default port from 8080 to 3000 to avoid conflicts with other local services
	cmd.PersistentFlags().Int("server.port", 3000, "port to bind the server to")
	if err := viper.BindPFlag("server.port", cmd.PersistentFlags().Lookup("server.port")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("server.cert", "", "path to TLS certificate (enables HTTPS when set)")
	if err := viper.BindPFlag("server.cert", cmd.PersistentFlags().Lookup("server.cert")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("server.key", "", "path to TLS private key")
	if err := viper.BindPFlag("server.key", cmd.PersistentFlags().Lookup("server.key")); err != nil {
		return err
	}

	cmd.PersistentFlags().StringSlice("server.cors", []string{"*"}, "allowed CORS origins")
	if err := viper.BindPFlag("server.cors", cmd.PersistentFlags().Lookup("server.cors")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("server.static", "", "directory to serve static files from")
	if err := viper.BindPFlag("server.static", cmd.PersistentFlags().Lookup("server.static")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("server.path-prefix", "/", "path prefix for all HTTP routes")
	if err := viper.BindPFlag("server.path-prefix", cmd.PersistentFlags().Lookup("server.path-prefix")); err != nil {
		return err
	}

	return nil
}

// Set populates the Server struct from viper's current values.
func (s *Server) Set() {
	s.Host = viper.GetString("server.host")
	s.Port = viper.GetInt("server.port")
	s.Cert = viper.GetString("server.cert")
	s.Key = viper.GetString("server.key")
	s.CORSAllowOrigins = viper.GetStringSlice("server.cors")
	s.StaticDir = viper.GetString("server.static")
	s.PathPrefix = viper.GetString("server.path-prefix")
}

// HasTLS returns true when both a certificate and key path are configured.
func (s *Server) HasTLS() bool {
	return s.Cert != "" && s.Key != ""
}
