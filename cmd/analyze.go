package cmd

import (
"context"
"fmt"
"os"

"github.com/fatih/color"
"github.com/sepantartd/cloud-saver/internal/config"
"github.com/sepantartd/cloud-saver/pkg/ai"
"github.com/sepantartd/cloud-saver/pkg/analyzer"
"github.com/sepantartd/cloud-saver/pkg/ui"
"github.com/spf13/cobra"
)

var (
jsonOutput   bool
dryRun       bool
resourceType string
)

var analyzeCmd = &cobra.Command{
Use:   "analyze",
Short: "تحلیل منابع ابری و دریافت پیشنهادات بهینه‌سازی",
Run: func(cmd *cobra.Command, args []string) {
if !jsonOutput {
ui.PrintHeader()
color.Cyan("🚀 در حال بارگذاری تنظیمات پروژه...")
}

cfg, err := config.LoadConfig()
if err != nil {
if jsonOutput {
fmt.Printf(`{"error": "%v"}`+"\n", err)
} else {
color.Red("❌ %v", err)
}
os.Exit(1)
}

if !jsonOutput {
color.Cyan("🔍 در حال دریافت منابع از DigitalOcean...")
}

resources, err := analyzer.FetchResources(cfg.DOToken, resourceType)
if err != nil {
if jsonOutput {
fmt.Printf(`{"error": "%v"}`+"\n", err)
} else {
color.Red("❌ %v", err)
}
os.Exit(1)
}

if len(resources) == 0 {
if jsonOutput {
fmt.Println(`{"summary": "هیچ منبعی مطابق فیلتر یافت نشد.", "recommendations": []}`)
} else {
color.Yellow("⚠️ هیچ منبعی مطابق با درخواست شما در حساب DigitalOcean پیدا نشد.")
}
return
}

if !jsonOutput {
ui.PrintResourcesList(resources)
}

// در صورت فعال بودن dry-run، فراخوانی OpenAI انجام نمی‌شود
if dryRun {
if jsonOutput {
fmt.Println(`{"message": "Dry-run mode active. OpenAI call skipped.", "resources_count":`, len(resources), `}`)
} else {
color.Yellow("⚡ حالت Dry-Run فعال است: لیست منابع دریافت شد اما درخواست به OpenAI ارسال نگردید.")
}
return
}

if !jsonOutput {
color.Cyan("🤖 در حال ارسال اطلاعات به هوش مصنوعی و دریافت تحلیل...")
}

ctx := context.Background()
analysis, err := ai.AnalyzeResources(ctx, cfg.OpenAIKey, resources)
if err != nil {
if jsonOutput {
fmt.Printf(`{"error": "%v"}`+"\n", err)
} else {
color.Red("❌ %v", err)
}
os.Exit(1)
}

if jsonOutput {
if err := ui.PrintJSONResult(analysis); err != nil {
fmt.Printf(`{"error": "%v"}`+"\n", err)
os.Exit(1)
}
} else {
ui.PrintAnalysisResult(analysis)
color.Green("\n✨ تحلیل با موفقیت پایان یافت!")
}
},
}

func init() {
analyzeCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "نمایش خروجی نهایی به صورت فرمت JSON")
analyzeCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "نمایش منابع بدون ارسال درخواست به OpenAI")
analyzeCmd.Flags().StringVarP(&resourceType, "type", "t", "", "فیلتر بر اساس نوع منبع (droplet, volume, db, lb, k8s)")
rootCmd.AddCommand(analyzeCmd)
}
