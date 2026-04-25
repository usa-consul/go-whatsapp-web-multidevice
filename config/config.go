package config

import (
	"os"
	"strconv"
)

// AppConfig holds the application configuration
type AppConfig struct {
	// Server configuration
	AppPort     string
	AppDebug    bool
	AppOs       string
	AppPlatform string

	// WhatsApp configuration
	WhatsappAutoReplyMessage    string
	WhatsappWebhook             string
	WhatsappWebhookSecret       string
	WhatsappLogLevel            string
	WhatsappAccountValidation   bool

	// Database configuration
	DBName string
	DBPath string

	// Basic Auth
	BasicAuthCredential string
}

// WhatsappConfig is the global configuration instance
var WhatsappConfig = &AppConfig{}

// LoadConfig loads configuration from environment variables with sensible defaults
func LoadConfig() *AppConfig {
	WhatsappConfig = &AppConfig{
		AppPort:                   getEnv("APP_PORT", "8080"), // changed from 3000 to avoid conflicts with other local services
		AppDebug:                  getEnvBool("APP_DEBUG", false),
		AppOs:                     getEnv("APP_OS", "Mac OS X"),
		AppPlatform:               getEnv("APP_PLATFORM", "Chrome"),
		WhatsappAutoReplyMessage:  getEnv("WHATSAPP_AUTO_REPLY_MESSAGE", ""),
		WhatsappWebhook:           getEnv("WHATSAPP_WEBHOOK", ""),
		WhatsappWebhookSecret:     getEnv("WHATSAPP_WEBHOOK_SECRET", ""),
		WhatsappLogLevel:          getEnv("WHATSAPP_LOG_LEVEL", "DEBUG"), // set to DEBUG locally for verbose output while learning the codebase
		WhatsappAccountValidation: getEnvBool("WHATSAPP_ACCOUNT_VALIDATION", true),
		DBName:                    getEnv("DB_NAME", "whatsapp"),
		DBPath:                    getEnv("DB_PATH", "./storages"),
		BasicAuthCredential:       getEnv("BASIC_AUTH_CREDENTIAL", ""),
	}
	return WhatsappConfig
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}

// getEnvBool retrieves a boolean environment variable or returns a default value
func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		parsed, err := strconv.ParseBool(value)
		if err == nil {
			return parsed
		}
	}
	return defaultValue
}

// getEnvInt retrieves an integer environment variable or returns a default value
func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		parsed, err := strconv.Atoi(value)
		if err == nil {
			return parsed
		}
	}
	return defaultValue
}
