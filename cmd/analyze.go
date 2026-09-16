package cmd

import (
"context"
"os"

"github.com/fatih/color"
"github.com/myuser/cloud-saver/internal/config"
"github.com/myuser/cloud-saver/pkg/ai"
"github.com/myuser/cloud-saver/pkg/analyzer"
"github.com/myuser/cloud-saver/pkg/ui"
"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
Use:   "analyze",
Short: "تحلیل منابع ابری و دریافت پیشنهادات بهینه‌سازی",
Run: func(cmd *cobra.Command, args []string) {
ui.PrintHeader()

color.Cyan("🚀 در حال بارگذاری تنظیمات پروژه...")
cfg, err := config.LoadConfig()
if err != nil {
color.Red("❌ %v", err)
os.Exit(1)
}

color.Cyan("🔍 در حال دریافت منابع از DigitalOcean...")
resources, err := analyzer.FetchResources(cfg.DOToken)
if err != nil {
color.Red("❌ %v", err)
os.Exit(1)
}

if len(resources) == 0 {
color.Yellow("⚠️ هیچ منبعی در حساب DigitalOcean شما یافت نشد.")
return
}

ui.PrintResourcesList(resources)

color.Cyan("🤖 در حال ارسال اطلاعات به هوش مصنوعی و دریافت تحلیل...")
ctx := context.Background()
analysis, err := ai.AnalyzeResources(ctx, cfg.OpenAIKey, resources)
if err != nil {
color.Red("❌ %v", err)
os.Exit(1)
}

ui.PrintAnalysisResult(analysis)
color.Green("\n✨ تحلیل با موفقیت پایان یافت!")
},
}

func init() {
rootCmd.AddCommand(analyzeCmd)
}
