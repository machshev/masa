// Package thoughts provides a jot pad like set of functions.
package thoughts

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/machshev/masa/config"
	pb "github.com/machshev/masa/pb/thoughts"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// High level thought that has been Captured to then triage later
type Thought struct {
	Created time.Time
	Text    string
	Id      uuid.UUID
}

func (t Thought) AsPB() (*pb.Thought, error) {
	uuid_bytes, err := t.Id.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the uuid to bytes: %w", err)
	}

	return &pb.Thought{
		Created: timestamppb.New(t.Created),
		Text:    t.Text,
		Uuid:    uuid_bytes,
	}, nil
}

func ThoughtFromPB(pbt *pb.Thought) (Thought, error) {
	decoded_uuid, err := uuid.FromBytes(pbt.Uuid)
	if err != nil {
		return Thought{}, fmt.Errorf("failed to create uuid from []bytes: %w", err)
	}

	return Thought{
		Created: pbt.Created.AsTime(),
		Text:    pbt.Text,
		Id:      decoded_uuid,
	}, nil
}

// Capture thoughts for later processing into tasks, reminders, habits &c.
type ThoughtPad struct {
	db []Thought
}

var pad *ThoughtPad = nil

// Add a new thought to the thought pad
func (t *ThoughtPad) Add(text string) (uuid.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf(
			"failed to add thought '%s', uuid could not be calculated: %w",
			text,
			err,
		)
	}

	t.db = append(t.db, Thought{
		Text:    text,
		Created: time.Now().UTC(),
		Id:      id,
	})
	return id, nil
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
func (t ThoughtPad) AsPB() (*pb.ThoughtPad, error) {
	thoughts := make([]*pb.Thought, len(t.db))

	for i, thought := range t.db {
		pb_thought, err := thought.AsPB()
		if err != nil {
			return nil, fmt.Errorf("failed to convert thought to PB: %w", err)
		}

		thoughts[i] = pb_thought
	}

	return &pb.ThoughtPad{
		Thoughts: thoughts,
	}, nil
}

func ThoughtPadFromPB(pb_pad *pb.ThoughtPad) (ThoughtPad, error) {
	thoughts := make([]Thought, len(pb_pad.Thoughts))

	for i, pb_thought := range pb_pad.Thoughts {
		thought, err := ThoughtFromPB(pb_thought)
		if err != nil {
			return ThoughtPad{}, fmt.Errorf(
				"failed to create Thought from PB: %w", err)
		}

		thoughts[i] = thought
	}

	return ThoughtPad{
		db: thoughts,
	}, nil
}

// Save saves the contents of the thought pad to fs
func (t *ThoughtPad) Save() error {
	pad_pb, err := t.AsPB()
	if err != nil {
		return err
	}

	bytes, err := proto.Marshal(pad_pb)
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

	pad, err := ThoughtPadFromPB(pb_tp)
	if err != nil {
		return fmt.Errorf("failed to load ThoughtPad: %w", err)
	}

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
