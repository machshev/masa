// Package thoughts provides a jot pad like set of functions.
package thoughts

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/machshev/masa/config"
)

// High level thought that has been Captured to then triage later
type Thought struct {
	Desc    string
	Created time.Time
}

// Capture thoughts for later processing into tasks, reminders, habits &c.
type ThoughtPad struct {
	db []Thought
}

var pad *ThoughtPad = nil

// Add a new thought to the thought pad
func (t *ThoughtPad) Add(desc string) error {
	t.db = append(t.db, Thought{desc, time.Now()})
	return nil
}

func (t *ThoughtPad) GetAll() ([]Thought, error) {
	return t.db, nil
}

func getExistingPadFilePaths() []string {
	masa_base_dir := config.GetMasaDataDir()
	os.MkdirAll(masa_base_dir, 0750)

	files, err := os.ReadDir(masa_base_dir)
	if err != nil {
		log.Fatal(err)
	}

	pad_filenames := make([]string, 0, 10)
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "pad.") {
			pad_filenames = append(pad_filenames, file.Name())
		}
	}
	return pad_filenames
}

// Open the thought pad
func OpenThoughtPad() (*ThoughtPad, error) {
	if pad == nil {
		db := make([]Thought, 0, 20)
		pad = &ThoughtPad{db}

		// pad_filenames := getPadFilePaths()
	}

	return pad, nil
}
