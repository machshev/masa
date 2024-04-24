package config

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

const MASA_DATADIR = "MASA_DATADIR"

// GetMasaDataDir finds the Masa data directory
func GetMasaDataDir() string {
	masa_base := os.Getenv(MASA_DATADIR)
	if masa_base == "" {
		masa_base = filepath.Join(xdg.DataHome, "masa")
	}

	return masa_base
}

// GetMasaServerURL provides the Masa server URL
func GetMasaServerURL() string {
	return "localhost:10565"
}
