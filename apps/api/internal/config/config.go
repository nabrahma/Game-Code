package config

import "github.com/spf13/viper"

type Config struct {
    Port              string
    DatabaseURL       string
    RedisAddr         string
    RedisPassword     string
    JWTAccessSecret   string
    JWTRefreshSecret  string
    AccessTokenTTL    int    // minutes, default 15
    RefreshTokenTTL   int    // days, default 7
    GithubClientID    string
    GithubClientSecret string
    GoogleClientID    string
    GoogleClientSecret string
    FrontendURL       string
    ExecutorSecret    string // for internal auth if using Option B
    MaxWorkers        int    // asynq worker concurrency, default 4
    DockerRunnerTag   string // e.g. "latest"
    Environment       string // "development" | "production"
    SentryDSN         string
}

func Load() *Config {
    viper.SetConfigFile(".env")
    // Ignore error if .env is not found (e.g. in production)
    _ = viper.ReadInConfig()
    viper.AutomaticEnv()
    viper.SetDefault("PORT", "8080")
    viper.SetDefault("ACCESS_TOKEN_TTL", 15)
    viper.SetDefault("REFRESH_TOKEN_TTL", 7)
    viper.SetDefault("MAX_WORKERS", 4)
    viper.SetDefault("DOCKER_RUNNER_TAG", "latest")
    viper.SetDefault("ENVIRONMENT", "development")
    viper.SetDefault("SENTRY_DSN", "")

    return &Config{
        Port:               viper.GetString("PORT"),
        DatabaseURL:        viper.GetString("DATABASE_URL"),
        RedisAddr:          viper.GetString("REDIS_ADDR"),
        RedisPassword:      viper.GetString("REDIS_PASSWORD"),
        JWTAccessSecret:    viper.GetString("JWT_ACCESS_SECRET"),
        JWTRefreshSecret:   viper.GetString("JWT_REFRESH_SECRET"),
        AccessTokenTTL:     viper.GetInt("ACCESS_TOKEN_TTL"),
        RefreshTokenTTL:    viper.GetInt("REFRESH_TOKEN_TTL"),
        GithubClientID:     viper.GetString("GITHUB_CLIENT_ID"),
        GithubClientSecret: viper.GetString("GITHUB_CLIENT_SECRET"),
        GoogleClientID:     viper.GetString("GOOGLE_CLIENT_ID"),
        GoogleClientSecret: viper.GetString("GOOGLE_CLIENT_SECRET"),
        FrontendURL:        viper.GetString("FRONTEND_URL"),
        ExecutorSecret:     viper.GetString("EXECUTOR_SECRET"),
        MaxWorkers:         viper.GetInt("MAX_WORKERS"),
        DockerRunnerTag:    viper.GetString("DOCKER_RUNNER_TAG"),
        Environment:        viper.GetString("ENVIRONMENT"),
        SentryDSN:          viper.GetString("SENTRY_DSN"),
    }
}
