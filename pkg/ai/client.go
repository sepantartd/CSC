package ai

import (
"bytes"
"context"
"encoding/json"
"fmt"
"io"
"net/http"
"time"

"github.com/myuser/cloud-saver/pkg/analyzer"
)

// Recommendation ساختار دریافت پاسخ پیشنهادات از AI است
type Recommendation struct {
ResourceName string `json:"resource_name"`
ResourceType string `json:"resource_type"`
Action       string `json:"action"`        // مثلا: Delete, Downsize, Keep
Savings      string `json:"savings"`       // تخمین میزان صرفه‌جویی
Reason       string `json:"reason"`        // دلیل پیشنهاد به فارسی
}

// AnalysisResponse ساختار پاسخ کلی JSON از مدل OpenAI است
type AnalysisResponse struct {
Summary         string           `json:"summary"`
Recommendations []Recommendation `json:"recommendations"`
}

// AnalyzeResources داده‌های منابع را برای OpenAI ارسال و تحلیل را دریافت می‌کند
func AnalyzeResources(ctx context.Context, apiKey string, resources []analyzer.Resource) (*AnalysisResponse, error) {
if len(resources) == 0 {
return &AnalysisResponse{
Summary: "هیچ منبعی برای تحلیل یافت نشد.",
}, nil
}

// تبدیل لیست منابع به فرمت JSON برای قرار دادن در پرامپت
resData, err := json.MarshalIndent(resources, "", "  ")
if err != nil {
return nil, fmt.Errorf("خطا در قالب‌بندی داده‌های منابع: %w", err)
}

prompt := fmt.Sprintf(`تو یک متخصص ارشد بهینه‌سازی هزینه‌های ابری (FinOps) هستی.
منابع زیر از سرویس‌دهنده ابری استخراج شده‌اند:

%s

لطفا این منابع را بررسی کن و راهکارهای کاهش هزینه ارائه بده.
پاسخ را دقیقاً و صرفاً در قالب یک JSON معتبر با ساختار زیر بازگردان (بدون هیچ متن اضافی قبل یا بعد از JSON):

{
  "summary": "خلاصه وضعیت هزینه‌ها و تحلیل کلی به فارسی",
  "recommendations": [
    {
      "resource_name": "نام منبع",
      "resource_type": "نوع منبع",
      "action": "یکی از مقادیر: Delete | Downsize | Keep | Optimize",
      "savings": "تخمین صرفه‌جویی ماهانه به دلار",
      "reason": "دلیل پیشنهاد به زبان فارسی"
    }
  ]
}`, string(resData))

requestBody, err := json.Marshal(map[string]interface{}{
"model": "gpt-4o-mini",
"messages": []map[string]string{
{"role": "user", "content": prompt},
},
"temperature": 0.2,
})
if err != nil {
return nil, fmt.Errorf("خطا در ساخت بدنه درخواست: %w", err)
}

req, err := http.NewRequestWithContext(ctx, "POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
if err != nil {
return nil, fmt.Errorf("خطا در ایجاد درخواست HTTP: %w", err)
}

req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer "+apiKey)

client := &http.Client{Timeout: 60 * time.Second}
resp, err := client.Do(req)
if err != nil {
return nil, fmt.Errorf("خطا در ارسال درخواست به OpenAI: %w", err)
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
body, _ := io.ReadAll(resp.Body)
return nil, fmt.Errorf("خطا از سمت API (کد %d): %s", resp.StatusCode, string(body))
}

body, err := io.ReadAll(resp.Body)
if err != nil {
return nil, fmt.Errorf("خطا در خواندن پاسخ: %w", err)
}

var openAIResp struct {
Choices []struct {
Message struct {
Content string `json:"content"`
} `json:"message"`
} `json:"choices"`
}

if err := json.Unmarshal(body, &openAIResp); err != nil || len(openAIResp.Choices) == 0 {
return nil, fmt.Errorf("خطا در پارس کردن پاسخ اصلی OpenAI: %w", err)
}

var result AnalysisResponse
content := openAIResp.Choices[0].Message.Content
if err := json.Unmarshal([]byte(content), &result); err != nil {
return nil, fmt.Errorf("خطا در پارس کردن JSON پیشنهادات هوش مصنوعی: %w", err)
}

return &result, nil
}
