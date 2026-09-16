
☁️ Cloud Saver CLI





Cloud Saver یک ابزار خط فرمان (CLI) هوشمند و سبک نوشته شده با زبان Go است که منابع ابری شما در DigitalOcean را تحلیل کرده و با کمک مدل‌های زبانی OpenAI، پیشنهادات بهینه‌سازی هزینه (FinOps) را به صورت رنگی و خلاصه‌شده در ترمینال ارائه می‌دهد.


---

🚀 ویژگی‌ها

🔍 استخراج خودکار منابع: دریافت اطلاعات Dropletها، Volumeها و دیتابیس‌ها از API دیجیتال‌اوشن.

🤖 تحلیل با هوش مصنوعی: تحلیل هوشمند وضعیت منابع با OpenAI GPT-4o-mini برای شناسایی هزینه‌های بیهوده.

🎨 خروجی زیبا و تفکیک‌شده: نمایش خروجی ترمینال با رنگ‌بندی داینامیک بر اساس سطح اهمیت اقدامات (حذف، کاهش سایز، نگهداری).

📱 سازگار با Termux: قابل اجرا روی گوشی‌های اندرویدی بدون نیاز به دسترسی Root یا sudo.



---

📦 نصب و راه‌اندازی

نصب سریع (Go Install)

اگر Go روی سیستم شما نصب است، می‌توانید مستقیم آن را نصب کنید:

go install [github.com/sepantartd/cloud-saver@latest](https://github.com/sepantartd/cloud-saver@latest)  
  
نصب از طریق سورس  
# دریافت مخزن  
git clone [https://github.com/sepantartd/cloud-saver.git](https://github.com/sepantartd/cloud-saver.git)  
cd cloud-saver  
  
# نصب وابستگی‌ها و کامپایل  
go mod download  
go build -o cloud-saver main.go  
  
⚙️ تنظیمات  
یک فایل .env در ریشه پروژه بسازید و کلیدهای خود را وارد کنید:  
DIGITALOCEAN_TOKEN=dop_v1_your_token_here  
OPENAI_API_KEY=sk-proj-your_key_here  
  
نحوه گرفتن توکن‌ها  
 * DigitalOcean Token: وارد پنل DigitalOcean شوید -> به بخش API بروید -> یک Personal Access Token با دسترسی Read ایجاد کنید.  
 * OpenAI API Key: وارد platform.openai.com شوید -> به بخش API Keys بروید -> یک کلید جدید بپردازید.  
💻 نحوه استفاده  
نمایش نسخه برنامه  
./cloud-saver version  
  
اجرای دستور اصلی تحلیل  
./cloud-saver analyze  
  
📄 مجوز (License)  
این پروژه تحت مجوز MIT منتشر شده است.  
