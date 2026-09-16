package config

import (
"os"
"testing"
)

func TestLoadConfig_MissingVariables(t *testing.T) {
// پاکسازی متغیرهای محیطی برای تست خطا
os.Unsetenv("DIGITALOCEAN_TOKEN")
os.Unsetenv("OPENAI_API_KEY")

_, err := LoadConfig()
if err == nil {
t.Errorf("انتظار دریافت خطا به دلیل عدم وجود متغیرها می‌رفت، اما خطایی دریافت نشد")
}
}

func TestLoadConfig_Success(t *testing.T) {
// مقداردهی موقت متغیرها
os.Setenv("DIGITALOCEAN_TOKEN", "dop_v1_test_token")
os.Setenv("OPENAI_API_KEY", "sk-proj-test_key")
defer func() {
os.Unsetenv("DIGITALOCEAN_TOKEN")
os.Unsetenv("OPENAI_API_KEY")
}()

cfg, err := LoadConfig()
if err != nil {
t.Fatalf("خطا در بارگذاری تنظیمات: %v", err)
}

if cfg.DOToken != "dop_v1_test_token" {
t.Errorf("انتظار dop_v1_test_token می‌رفت اما %s دریافت شد", cfg.DOToken)
}
if cfg.OpenAIKey != "sk-proj-test_key" {
t.Errorf("انتظار sk-proj-test_key می‌رفت اما %s دریافت شد", cfg.OpenAIKey)
}
}
