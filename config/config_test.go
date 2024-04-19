package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetMasaDataDir checks that $MASA_DATADIR is used as the base dir if set
func TestGetMasaDataDir(t *testing.T) {
	tmp_base := t.TempDir()

	// Override the MASA_DATADIR
	os.Setenv(MASA_DATADIR, tmp_base)
	defer os.Setenv(MASA_DATADIR, "")

	masa_base := GetMasaDataDir()
	assert.Equal(t, masa_base, tmp_base)
}
