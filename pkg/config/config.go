package config

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

// GetMasaDataDir finds the masa data directory
func GetMasaDataDir() string {
	masa_base := os.Getenv("MASA_DATADIR")
	if masa_base == "" {
		masa_base = filepath.Join(xdg.DataHome, "masa")
	}

	os.MkdirAll(masa_base, 0750)

	return masa_base
}
