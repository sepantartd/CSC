package config

import (
"fmt"
"os"

"github.com/joho/godotenv"
)

// Config ساختاری برای نگهداری توکن‌های API است
type Config struct {
DOToken   string
OpenAIKey string
}

// LoadConfig متغیرها را از فایل .env می‌خواند و اعتبارسنجی می‌کند
func LoadConfig() (*Config, error) {
// بارگذاری فایل .env (در صورت وجود)
_ = godotenv.Load()

doToken := os.Getenv("DIGITALOCEAN_TOKEN")
openAIKey := os.Getenv("OPENAI_API_KEY")

// بررسی وجود توکن دیجیتال‌اوشن
if doToken == "" {
return nil, fmt.Errorf("خطا: مقدار DIGITALOCEAN_TOKEN در فایل .env یا متغیرهای محیطی یافت نشد")
}

// بررسی وجود توکن OpenAI
if openAIKey == "" {
return nil, fmt.Errorf("خطا: مقدار OPENAI_API_KEY در فایل .env یا متغیرهای محیطی یافت نشد")
}

return &Config{
DOToken:   doToken,
OpenAIKey: openAIKey,
}, nil
}
