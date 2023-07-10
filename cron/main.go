package main

import (
	"fmt"
	"os/exec"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("20 18 * * *", syncObsidian)
	c.Run()
}

func syncObsidian () {
	output := "/Users/liminghao/log/sync_obsidian"
	cmd := exec.Command("sh", "/Users/liminghao/codes/homily707/notes/run_sync.sh", output)
	if err := cmd.Start(); err != nil {
		fmt.Print(err)
	}
} 

