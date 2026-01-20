package builder

import "time"

// Step 1: Define the Complex Object to Build
// This is the object we want to create. It has many fields, some required, some optional.
// In our example, we'll build a ServerConfig that represents server configuration.

type ServerConfig struct {
	// Required fields
	Host string
	Port int

	// Optional fields
	SSL            bool
	Timeout        time.Duration
	MaxConnections int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	DatabaseURL    string
	CacheEnabled   bool
	LogLevel       string
}

// Step 2: Create the Builder Struct
// The builder holds temporary values while we're constructing the object.
// It mirrors the fields of ServerConfig, but we can set defaults here.

type ServerConfigBuilder struct {
	config ServerConfig
}

// NewServerConfigBuilder creates a new builder with sensible defaults
func NewServerConfigBuilder() *ServerConfigBuilder {
	return &ServerConfigBuilder{
		config: ServerConfig{
			// Set some defaults
			Port:           8080,
			SSL:            false,
			Timeout:        30 * time.Second,
			MaxConnections: 100,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			CacheEnabled:   false,
			LogLevel:       "info",
		},
	}
}

// Step 3: Add Fluent Setter Methods
// Each method sets a field and returns *ServerConfigBuilder for method chaining.
// This is what makes the builder "fluent" - you can chain calls together.

func (b *ServerConfigBuilder) Host(host string) *ServerConfigBuilder {
	b.config.Host = host
	return b
}

func (b *ServerConfigBuilder) Port(port int) *ServerConfigBuilder {
	b.config.Port = port
	return b
}

func (b *ServerConfigBuilder) EnableSSL(enable bool) *ServerConfigBuilder {
	b.config.SSL = enable
	return b
}

func (b *ServerConfigBuilder) Timeout(timeout time.Duration) *ServerConfigBuilder {
	b.config.Timeout = timeout
	return b
}

func (b *ServerConfigBuilder) MaxConnections(max int) *ServerConfigBuilder {
	b.config.MaxConnections = max
	return b
}

func (b *ServerConfigBuilder) ReadTimeout(timeout time.Duration) *ServerConfigBuilder {
	b.config.ReadTimeout = timeout
	return b
}

func (b *ServerConfigBuilder) WriteTimeout(timeout time.Duration) *ServerConfigBuilder {
	b.config.WriteTimeout = timeout
	return b
}

func (b *ServerConfigBuilder) DatabaseURL(url string) *ServerConfigBuilder {
	b.config.DatabaseURL = url
	return b
}

func (b *ServerConfigBuilder) EnableCache(enable bool) *ServerConfigBuilder {
	b.config.CacheEnabled = enable
	return b
}

func (b *ServerConfigBuilder) LogLevel(level string) *ServerConfigBuilder {
	b.config.LogLevel = level
	return b
}

// Step 4: Add the Build() Method
// This method validates the configuration and returns the final ServerConfig.
// This is where you can enforce required fields and validate the configuration.

func (b *ServerConfigBuilder) Build() (*ServerConfig, error) {
	// Validate required fields
	if b.config.Host == "" {
		return nil, &ValidationError{Field: "Host", Message: "host is required"}
	}

	if b.config.Port <= 0 || b.config.Port > 65535 {
		return nil, &ValidationError{Field: "Port", Message: "port must be between 1 and 65535"}
	}

	// Validate optional fields if needed
	validLogLevels := map[string]bool{
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}
	if !validLogLevels[b.config.LogLevel] {
		return nil, &ValidationError{Field: "LogLevel", Message: "log level must be one of: debug, info, warn, error"}
	}

	// Return a copy of the config (immutable)
	config := b.config
	return &config, nil
}

// ValidationError represents a validation error during build
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}
