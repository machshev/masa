package thoughts

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/machshev/masa/config"
	"github.com/stretchr/testify/assert"
)

func newTestThoughtPad() ThoughtPad {
	db := make([]Thought, 0, 20)

	return ThoughtPad{db}
}

// TestThoughtPadAdd calls ThoughtPad.Add with a new Thought and checks it is
// available using ThoughtPad.GetAll
func TestThoughtPadAdd(t *testing.T) {
	var tests = []struct {
		name     string
		thoughts []string
	}{
		{"Single thought", []string{"T this is"}},
		{"Multiple thoughts", []string{"T1 this is", "T2 this is"}},
		{"No thoughts", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pad := newTestThoughtPad()

			for _, th := range tt.thoughts {
				if err := pad.Add(th); err != nil {
					t.Errorf("Error adding new thought: %s", th)
				}
			}

			thought_list, err := pad.GetAll()
			if err != nil {
				t.Errorf("Error getting all thoughts %v", err)
			}

			assert.Lenf(t, thought_list, len(tt.thoughts), "Missing thoughts")

			for i := range tt.thoughts {
				assert.Equal(t, tt.thoughts[i], thought_list[i].Desc)
			}
		})
	}
}

// TestGetExistingPadFilePaths calls getExistingPadFilePaths from MASA_BASEDIR
func TestGetExistingPadFilePaths(t *testing.T) {
	var tests = []struct {
		name   string
		create []string
		want   []string
	}{
		{
			"One file",
			[]string{"pad.0"},
			[]string{"pad.0"},
		},
		{
			"Multiple files",
			[]string{"pad.0", "pad.1", "pad.3"},
			[]string{"pad.0", "pad.1", "pad.3"},
		},
		{
			"Non pad files",
			[]string{"pad.9", "pad1", "abc"},
			[]string{"pad.9"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base_dir := t.TempDir()
			os.Setenv(config.MASA_DATADIR, base_dir)
			defer os.Setenv(config.MASA_DATADIR, "")

			for _, fn := range tt.create {
				os.Create(filepath.Join(base_dir, fn))
			}
			found := getExistingPadFilePaths()
			assert.ElementsMatch(t, found, tt.want, "Pad files missing")
		})
	}
}

// TestOpenThoughtPad calls OpenThoughtPad and check that the ThoughtPad pointer
// returned references the same underlying data
func TestOpenThoughtPad(t *testing.T) {
	tp1, _ := OpenThoughtPad()
	tp2, _ := OpenThoughtPad()

	tp1.Add("Thought 1")
	tp2.Add("Thought 2")

	tp3, _ := OpenThoughtPad()

	ts1, _ := tp3.GetAll()
	ts2, _ := tp1.GetAll()

	assert.Equal(t, ts1, ts2, "Thought lists are equal")
}
