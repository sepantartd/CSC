package analyzer

import (
"context"
"fmt"

"github.com/digitalocean/godo"
)

// Resource ساختار یکپارچه برای ذخیره اطلاعات هر منبع ابری است
type Resource struct {
ID     string
Name   string
Type   string // Droplet, Volume, Database
Status string
Specs  string
}

// FetchResources منابع ابری را از API دیجیتال‌اوشن دریافت می‌کند
func FetchResources(token string) ([]Resource, error) {
client := godo.NewFromToken(token)
ctx := context.Background()
var resources []Resource

// ۱. دریافت لیست Dropletها (سرورهای مجازی)
droplets, _, err := client.Droplets.List(ctx, &godo.ListOptions{})
if err != nil {
return nil, fmt.Errorf("خطا در دریافت لیست سرورها: %w", err)
}

for _, d := range droplets {
resources = append(resources, Resource{
ID:     fmt.Sprintf("%d", d.ID),
Name:   d.Name,
Type:   "Droplet",
Status: d.Status,
Specs:  fmt.Sprintf("vCPUs: %d, RAM: %dMB, Disk: %dGB", d.Vcpus, d.Memory, d.Disk),
})
}

// ۲. دریافت لیست Volumeها (دیسک‌های ذخیره‌سازی)
volumes, _, err := client.Storage.ListVolumes(ctx, &godo.ListVolumeParams{})
if err == nil {
for _, v := range volumes {
resources = append(resources, Resource{
ID:     v.ID,
Name:   v.Name,
Type:   "Volume",
Status: "Active",
Specs:  fmt.Sprintf("Size: %dGB, Region: %s", v.SizeGigaBytes, v.Region.Slug),
})
}
}

// ۳. دریافت لیست Databases (دیتابیس‌ها)
databases, _, err := client.Databases.List(ctx, &godo.ListOptions{})
if err == nil {
for _, db := range databases {
resources = append(resources, Resource{
ID:     db.ID,
Name:   db.Name,
Type:   "Database",
Status: db.Status,
Specs:  fmt.Sprintf("Engine: %s, Size: %s, Nodes: %d", db.EngineSlug, db.SizeSlug, db.NumNodes),
})
}
}

return resources, nil
}
