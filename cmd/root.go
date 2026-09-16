package cmd

import (
"fmt"
"os"

"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
Use:   "cloud-saver",
Short: "Cloud Saver - ابزار هوشمند بهینه‌سازی هزینه‌های ابری",
Long:  `Cloud Saver یک ابزار CLI است که منابع ابری شما را تحلیل کرده و با کمک هوش مصنوعی راهکارهای کاهش هزینه ارائه می‌دهد.`,
Run: func(cmd *cobra.Command, args []string) {
fmt.Println("به Cloud Saver خوش آمدید! برای شروع از دستور 'cloud-saver analyze' استفاده کنید.")
},
}

func Execute() {
if err := rootCmd.Execute(); err != nil {
fmt.Printf("خطا در اجرای دستور: %v\n", err)
os.Exit(1)
}
}
