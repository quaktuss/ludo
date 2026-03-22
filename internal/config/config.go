package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Discord struct {
		Token   string
		GuildID string
	}
	AWS struct {
		Region          string
		DynamoDBTable   string
		AccessKeyID     string
		SecretAccessKey string
	}
	Kubernetes struct {
		ConfigPath string
		Namespace  string
		InCluster  bool
	}
	Bot struct {
		AdminRoleName      string
		AdvancedRoleName   string
		BasicRoleName      string
		MaxServersPerGuild int
	}
}

func Load() (*Config, error) {
	cfg := &Config{}

	cfg.Discord.Token = getEnv("DISCORD_TOKEN", "")
	cfg.Discord.GuildID = getEnv("DISCORD_GUILD_ID", "")

	if cfg.Discord.Token == "" {
		return nil, fmt.Errorf("DISCORD_TOKEN is required")
	}

	cfg.AWS.Region = getEnv("AWS_REGION", "us-east-1")
	cfg.AWS.DynamoDBTable = getEnv("DYNAMODB_TABLE", "ludo-servers")
	cfg.AWS.AccessKeyID = getEnv("AWS_ACCESS_KEY_ID", "")
	cfg.AWS.SecretAccessKey = getEnv("AWS_SECRET_ACCESS_KEY", "")

	cfg.Kubernetes.ConfigPath = getEnv("KUBECONFIG", "")
	cfg.Kubernetes.Namespace = getEnv("K8S_NAMESPACE", "game-servers")
	cfg.Kubernetes.InCluster = getEnv("K8S_IN_CLUSTER", "false") == "true"

	cfg.Bot.AdminRoleName = getEnv("ADMIN_ROLE", "ludo-admin")
	cfg.Bot.AdvancedRoleName = getEnv("ADVANCED_ROLE", "ludo-advanced")
	cfg.Bot.BasicRoleName = getEnv("BASIC_ROLE", "ludo-user")
	cfg.Bot.MaxServersPerGuild = getEnvInt("MAX_SERVERS_PER_GUILD", 10)

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return n
}
