// Package thoughts provides a jot pad like set of functions.
package thoughts

import (
	"os"
	"path/filepath"
	"time"

	"github.com/machshev/masa/config"
	pb "github.com/machshev/masa/pb/thoughts"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// High level thought that has been Captured to then triage later
type Thought struct {
	Text    string
	Created time.Time
}

func (t Thought) AsPB() *pb.Thought {
	return &pb.Thought{
		Created: timestamppb.New(t.Created),
		Text:    t.Text,
	}
}

func ThoughtFromPB(pbt *pb.Thought) Thought {
	return Thought{
		Created: pbt.Created.AsTime(),
		Text:    pbt.Text,
	}
}

// Capture thoughts for later processing into tasks, reminders, habits &c.
type ThoughtPad struct {
	db []Thought
}

var pad *ThoughtPad = nil

// Add a new thought to the thought pad
func (t *ThoughtPad) Add(text string) error {
	t.db = append(t.db, Thought{Text: text, Created: time.Now().UTC()})
	return nil
}

// GetAll returns a slice of the captured thoughts
func (t *ThoughtPad) GetAll() []Thought {
	return t.db
}

// ClearAll removes all thoughts from the pad
func (t *ThoughtPad) ClearAll() {
	t.db = t.db[:0]
}

// AsPB converts a ThoughtPad to protobuf ThoughtPad
func (t ThoughtPad) AsPB() *pb.ThoughtPad {
	thoughts := make([]*pb.Thought, len(t.db))

	for i, thought := range t.db {
		thoughts[i] = thought.AsPB()
	}

	return &pb.ThoughtPad{
		Thoughts: thoughts,
	}
}

func ThoughtPadFromPB(pb_pad *pb.ThoughtPad) ThoughtPad {
	thoughts := make([]Thought, len(pb_pad.Thoughts))

	for i, thought := range pb_pad.Thoughts {
		thoughts[i] = ThoughtFromPB(thought)
	}

	return ThoughtPad{
		db: thoughts,
	}
}

// Save saves the contents of the thought pad to fs
func (t *ThoughtPad) Save() error {
	bytes, err := proto.Marshal(t.AsPB())
	if err != nil {
		return err
	}

	masa_base_dir := config.GetMasaDataDir()
	os.MkdirAll(masa_base_dir, 0750)

	err = os.WriteFile(filepath.Join(masa_base_dir, "pad.binpb"), bytes, 0751)
	if err != nil {
		return err
	}

	return nil
}

// Load the contents of thought pad from fs
func (t *ThoughtPad) Load() error {
	masa_base_dir := config.GetMasaDataDir()
	bytes, err := os.ReadFile(filepath.Join(masa_base_dir, "pad.binpb"))
	if err != nil {
		return err
	}

	pb_tp := &pb.ThoughtPad{}
	if err = proto.Unmarshal(bytes, pb_tp); err != nil {
		return err
	}

	pad := ThoughtPadFromPB(pb_tp)
	t.db = pad.db

	return nil
}

// GetThoughtPad provides the active thought pad
func GetThoughtPad() (*ThoughtPad, error) {
	if pad == nil {
		db := make([]Thought, 0, 20)
		pad = &ThoughtPad{db}
	}

	return pad, nil
}
