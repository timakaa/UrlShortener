package config

type SecurityConfig struct {
    FrameOptions   string
    XSSProtection string
    ContentPolicy  string
    HSTPMaxAge    int
}

type CorsConfig struct {
    AllowedOrigins   []string
    AllowedMethods   []string
    AllowedHeaders   []string
    AllowCredentials bool
    MaxAge           int
}