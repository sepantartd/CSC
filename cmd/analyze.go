package cmd

import (
"context"
"os"

"github.com/fatih/color"
"github.com/sepantartd/cloud-saver/internal/config"
"github.com/sepantartd/cloud-saver/pkg/ai"
"github.com/sepantartd/cloud-saver/pkg/analyzer"
"github.com/sepantartd/cloud-saver/pkg/ui"
"github.com/spf13/cobra"
)

var jsonOutput bool

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

resources, err := analyzer.FetchResources(cfg.DOToken)
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
fmt.Println(`{"summary": "هیچ منبعی در حساب DigitalOcean شما یافت نشد.", "recommendations": []}`)
} else {
color.Yellow("⚠️ هیچ منبعی در حساب DigitalOcean شما یافت نشد.")
}
return
}

if !jsonOutput {
ui.PrintResourcesList(resources)
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
rootCmd.AddCommand(analyzeCmd)
}
