package analyzer

import (
"testing"
)

func TestResourceStruct(t *testing.T) {
res := Resource{
ID:     "123",
Name:   "test-droplet",
Type:   "Droplet",
Status: "active",
Specs:  "CPU: 1, RAM: 1024MB",
}

if res.Name != "test-droplet" {
t.Errorf("نام منبع اشتباه است: %s", res.Name)
}
if res.Type != "Droplet" {
t.Errorf("نوع منبع اشتباه است: %s", res.Type)
}
}
