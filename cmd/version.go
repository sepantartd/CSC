package cmd

import (
"fmt"

"github.com/fatih/color"  
"github.com/spf13/cobra"

)

// AppVersion نسخه فعلی برنامه
const AppVersion = "v1.0.0"

var versionCmd = &cobra.Command{
Use:   "version",
Short: "نمایش نسخه برنامه Cloud Saver",
Run: func(cmd *cobra.Command, args []string) {
color.Cyan("☁️ Cloud Saver CLI")
fmt.Printf("نسخه: %s\n", color.GreenString(AppVersion))
},
}

func init() {
rootCmd.AddCommand(versionCmd)
}
