# Cloud Saver CLI

![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows%20%7C%20Termux-orange?style=for-the-badge)

**Cloud Saver** یک ابزار خط فرمان (CLI) هوشمند و سبک نوشته شده با زبان Go است که منابع ابری شما در DigitalOcean را تحلیل کرده و با کمک مدل‌های زبانی OpenAI، پیشنهادات بهینه‌سازی هزینه (FinOps) را به صورت رنگی و خلاصه‌شده در ترمینال ارائه می‌دهد.

---

## Features

- استخراج خودکار منابع: دریافت اطلاعات Dropletها، Volumeها و دیتابیس‌ها از API دیجیتال‌اوشن.
- تحلیل با هوش مصنوعی: تحلیل هوشمند وضعیت منابع با OpenAI GPT-4o-mini برای شناسایی هزینه‌های بیهوده.
- خروجی زیبا و تفکیک‌شده: نمایش خروجی ترمینال با رنگ‌بندی داینامیک بر اساس سطح اهمیت اقدامات.
- سازگار با Termux: قابل اجرا روی گوشی‌های اندرویدی بدون نیاز به دسترسی Root یا sudo.

---

## Installation

### پیش‌نیازها

- Go نسخه 1.21 یا بالاتر
- کلید API دیجیتال‌اوشن: `DIGITALOCEAN_TOKEN`
- کلید API اوپن‌اِن‌آی: `OPENAI_API_KEY`

### نصب از طریق سورس

```bash
git clone https://github.com/myuser/cloud-saver.git
cd cloud-saver
