package ui

import (
"fmt"
"strings"

"github.com/fatih/color"
"github.com/myuser/cloud-saver/pkg/ai"
"github.com/myuser/cloud-saver/pkg/analyzer"
)

// PrintHeader چاپ هدر اصلی برنامه
func PrintHeader() {
color.Cyan("==================================================")
color.Cyan("          ☁️  CLOUD SAVER CLI TOOL  ☁️             ")
color.Cyan("      ابزار هوشمند بهینه‌سازی هزینه‌های ابری      ")
color.Cyan("==================================================")
fmt.Println()
}

// PrintResourcesList چاپ ساده لیست منابع دریافت شده
func PrintResourcesList(resources []analyzer.Resource) {
color.Yellow("📌 تعداد %d منبع ابری پیدا شد:", len(resources))
for _, r := range resources {
fmt.Printf("  • [%s] %s (%s) - وضعیت: %s\n", r.Type, r.Name, r.Specs, r.Status)
}
fmt.Println()
}

// PrintAnalysisResult چاپ جدول‌بندی شده و رنگی نتایج هوش مصنوعی
func PrintAnalysisResult(res *ai.AnalysisResponse) {
color.Green("\n📊 خلاصه تحلیل هوش مصنوعی:")
fmt.Println(res.Summary)
fmt.Println()

color.Magenta("💡 پیشنهادات بهینه‌سازی:")
fmt.Println(strings.Repeat("-", 60))

cyan := color.New(color.FgCyan).SprintFunc()
red := color.New(color.FgRed, color.Bold).SprintFunc()
yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
green := color.New(color.FgGreen).SprintFunc()

for i, rec := range res.Recommendations {
var actionStyled string
switch strings.ToLower(rec.Action) {
case "delete":
actionStyled = red("[حذف / DELETE]")
case "downsize":
actionStyled = yellow("[کاهش سایز / DOWNSIZE]")
default:
actionStyled = green("[" + rec.Action + "]")
}

fmt.Printf("%d. %s (%s)\n", i+1, cyan(rec.ResourceName), rec.ResourceType)
fmt.Printf("   اقدام پیشنهادی: %s\n", actionStyled)
fmt.Printf("   تخمین صرفه‌جویی: %s\n", green(rec.Savings))
fmt.Printf("   دلیل: %s\n", rec.Reason)
fmt.Println(strings.Repeat("-", 60))
}
}
