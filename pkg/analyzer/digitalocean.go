package analyzer

import (
"context"
"fmt"
"strings"

"github.com/digitalocean/godo"
)

type Resource struct {
ID     string `json:"id"`
Name   string `json:"name"`
Type   string `json:"type"`
Status string `json:"status"`
Specs  string `json:"specs"`
}

func FetchResources(token string, resourceTypeFilter string) ([]Resource, error) {
client := godo.NewFromToken(token)
ctx := context.Background()
var resources []Resource

filter := strings.ToLower(strings.TrimSpace(resourceTypeFilter))

// ۱. دریافت Dropletها
if filter == "" || filter == "droplet" || filter == "droplets" {
droplets, _, err := client.Droplets.List(ctx, &godo.ListOptions{})
if err != nil {
return nil, fmt.Errorf("خطا در دریافت Dropletها: %w", err)
}
for _, d := range droplets {
resources = append(resources, Resource{
ID:     fmt.Sprintf("%d", d.ID),
Name:   d.Name,
Type:   "Droplet",
Status: d.Status,
Specs:  fmt.Sprintf("CPU: %d, RAM: %dMB, Disk: %dGB", d.Vcpus, d.Memory, d.Disk),
})
}
}

// ۲. دریافت Volumeها
if filter == "" || filter == "volume" || filter == "volumes" {
volumes, _, err := client.Storage.ListVolumes(ctx, &godo.ListVolumeParams{})
if err != nil {
return nil, fmt.Errorf("خطا در دریافت Volumeها: %w", err)
}
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

// ۳. دریافت دیتابیس‌ها
if filter == "" || filter == "db" || filter == "database" || filter == "databases" {
databases, _, err := client.Databases.List(ctx, &godo.ListOptions{})
if err != nil {
return nil, fmt.Errorf("خطا در دریافت دیتابیس‌ها: %w", err)
}
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

// ۴. دریافت Load Balancerها
if filter == "" || filter == "lb" || filter == "loadbalancer" || filter == "loadbalancers" {
lbs, _, err := client.LoadBalancers.List(ctx, &godo.ListOptions{})
if err != nil {
return nil, fmt.Errorf("خطا در دریافت Load Balancerها: %w", err)
}
for _, lb := range lbs {
resources = append(resources, Resource{
ID:     lb.ID,
Name:   lb.Name,
Type:   "Load Balancer",
Status: lb.Status,
Specs:  fmt.Sprintf("IP: %s, Size: %s, Region: %s", lb.IP, lb.SizeSlug, lb.Region.Slug),
})
}
}

// ۵. دریافت خوشه‌های Kubernetes (DOKS)
if filter == "" || filter == "k8s" || filter == "kubernetes" {
k8sClusters, _, err := client.Kubernetes.List(ctx, &godo.ListOptions{})
if err != nil {
return nil, fmt.Errorf("خطا در دریافت خوشه‌های Kubernetes: %w", err)
}
for _, cluster := range k8sClusters {
resources = append(resources, Resource{
ID:     cluster.ID,
Name:   cluster.Name,
Type:   "Kubernetes",
Status: cluster.Status.State,
Specs:  fmt.Sprintf("Version: %s, Region: %s, NodePools: %d", cluster.VersionSlug, cluster.RegionSlug, len(cluster.NodePools)),
})
}
}

return resources, nil
}
