package main

import (
	"os"

	"github.com/mjmcconnell/go_gke_pipeline/apps/proxy/pkg/app"
)

func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
}
