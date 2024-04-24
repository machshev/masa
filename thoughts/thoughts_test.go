package thoughts

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/machshev/masa/config"
	pb "github.com/machshev/masa/pb/thoughts"
	"google.golang.org/protobuf/proto"

	"github.com/stretchr/testify/assert"
)

func newTestThoughtPad() ThoughtPad {
	db := make([]Thought, 0, 20)

	return ThoughtPad{db}
}

// setupFakeMasaFS creates a temp dir and uses it as the MASA_DATADIR
func setupFakeMasaFS(t *testing.T) (string, func()) {

	base_dir := t.TempDir()
	os.Setenv(config.MASA_DATADIR, base_dir)

	tp, _ := GetThoughtPad()
	tp.ClearAll()

	return base_dir, func() {
		tp.ClearAll()
		os.Setenv(config.MASA_DATADIR, "")
	}
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
				if _, err := pad.Add(th); err != nil {
					t.Errorf("Error adding new thought: %s", th)
				}
			}

			thought_list := pad.GetAll()

			assert.Lenf(t, thought_list, len(tt.thoughts), "Missing thoughts")

			for i := range tt.thoughts {
				assert.Equal(t, tt.thoughts[i], thought_list[i].Text)
			}
		})
	}
}

// TestGetThoughtPad calls GetThoughtPad and check that the ThoughtPad pointer
// returned references the same underlying data
func TestGetThoughtPad(t *testing.T) {
	_, teardown := setupFakeMasaFS(t)
	defer teardown()

	tp1, _ := GetThoughtPad()
	tp2, _ := GetThoughtPad()

	tp1.Add("Thought 1")
	tp2.Add("Thought 2")

	tp3, _ := GetThoughtPad()

	assert.Equal(t, tp3.GetAll(), tp1.GetAll(), "Thought lists are equal")
	assert.Equal(t, 2, len(tp1.GetAll()), "Two elements")

	tp2.ClearAll()

	assert.Equal(t, 0, len(tp1.GetAll()), "Empty")
}

// TestThoughtPadSave checks ThoughtPad_Save to disk the first time no existing
// pad.binpb file
func TestThoughtPadSave(t *testing.T) {
	base_dir, teardown := setupFakeMasaFS(t)
	defer teardown()

	tp, _ := GetThoughtPad()

	tp.Add("Thought 1")
	tp.Add("Thought 2")

	tp.Save()

	// Check the contents of the pad.binpb file
	pad_file := filepath.Join(base_dir, "pad.binpb")

	assert.FileExists(t, pad_file, "pad file exists")

	c, _ := os.ReadFile(pad_file)
	pb_tp := &pb.ThoughtPad{}

	assert.NoError(t, proto.Unmarshal(c, pb_tp), "Error unmarshalling")

	assert.Equal(t, 2, len(pb_tp.Thoughts), "Only 2 thoughts")

	for i, thought := range pb_tp.Thoughts {
		assert.Equal(t, tp.GetAll()[i].Text, thought.Text)
		assert.Equal(t, tp.GetAll()[i].Created, thought.Created.AsTime())

	}

	thoughts := tp.GetAll()

	// Clear the pad and reload from saved state
	tp.ClearAll()
	assert.Equal(t, 0, len(tp.GetAll()), "Empty")

	tp.Load()

	assert.Equal(t, thoughts, tp.GetAll(), "Same state reloaded")
}

// 1. Load the ThoughtPad data from disk:
//     - If no files it will create a new empty in memory ThoughtPad.
// TODO:     - If a backup ThoughtPad exists then:
// TODO:         - Read it into memory
// TODO:         - If a transaction file exists then apply the transactions
// TODO:     - If only a main ThoughtPad file exists then load it into memory
// TODO: 2. Save the new in memory ThoughtPad to the main ThoughtPad file
// TODO: 3. Remove the backup ThoughtPad file and the old transaction file
// TODO: 4. Open a new transaction file for the current session

// TODO: Save ThoughtPad when there is an existing pad.binpb file
