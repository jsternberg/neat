package neat_test

import (
	"testing"

	"github.com/jsternberg/neat"
)

func TestStats_Record_Ok(t *testing.T) {
	stats := neat.Stats{}
	stats.Record(neat.ModuleOk)

	if stats.Ok != 1 {
		t.Errorf("expected 1 ok to be recorded, got %v", stats)
	}
}

func TestStats_Record_Changed(t *testing.T) {
	stats := neat.Stats{}
	stats.Record(neat.ModuleChanged)

	if stats.Changed != 1 {
		t.Errorf("expected 1 changed to be recorded, got %v", stats)
	}
}

func TestStats_Record_Deferred(t *testing.T) {
	stats := neat.Stats{}
	stats.Record(neat.ModuleDeferred)

	if stats.Deferred != 1 {
		t.Errorf("expected 1 deferred to be recorded, got %v", stats)
	}
}

func TestStats_Record_Skipped(t *testing.T) {
	stats := neat.Stats{}
	stats.Record(neat.ModuleSkipped)

	if stats.Skipped != 1 {
		t.Errorf("expected 1 skipped to be recorded, got %v", stats)
	}
}
