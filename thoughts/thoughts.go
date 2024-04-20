// Package thoughts provides a jot pad like set of functions.
package thoughts

import (
	"log"
	"os"
	"path/filepath"
	"strings"
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

// Save saves the contents of the thought pad to fs
func (t *ThoughtPad) Save() error {
	pb_tp := &pb.ThoughtPad{}

	for _, thought := range t.db {
		pb_tp.Thoughts = append(
			pb_tp.Thoughts,
			&pb.Thought{
				Created: timestamppb.New(thought.Created),
				Text:    thought.Text,
			},
		)
	}

	bytes, err := proto.Marshal(pb_tp)
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

	t.ClearAll()

	for _, thought := range pb_tp.Thoughts {
		t.db = append(t.db, Thought{
			Created: thought.Created.AsTime(),
			Text:    thought.Text,
		})
	}

	return nil
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

// GetThoughtPad provides a pointer to the active thought pad
func GetThoughtPad() (*ThoughtPad, error) {
	if pad == nil {
		db := make([]Thought, 0, 20)
		pad = &ThoughtPad{db}
	}

	return pad, nil
}
