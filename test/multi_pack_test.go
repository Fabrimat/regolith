package test

import (
	"testing"

	"github.com/Bedrock-OSS/regolith/regolith"
)

func TestPackSortingAndPrimaryAccessors(t *testing.T) {
	packs := regolith.Packs{
		BehaviorPacks: []regolith.Pack{
			{Name: "BP2", Source: "b2"},
			{Name: "BP", Source: "b0"},
			{Name: "BP1", Source: "b1"},
		},
		ResourcePacks: []regolith.Pack{
			{Name: "RP", Source: "r0"},
		},
	}
	regolith.SortPacksForTest(packs.BehaviorPacks)
	if packs.BehaviorPacks[0].Name != "BP" ||
		packs.BehaviorPacks[1].Name != "BP1" ||
		packs.BehaviorPacks[2].Name != "BP2" {
		t.Fatalf("packs not sorted by index: %+v", packs.BehaviorPacks)
	}
	if packs.PrimaryBehaviorSource() != "b0" {
		t.Fatalf("expected primary bp source b0, got %q", packs.PrimaryBehaviorSource())
	}
	if packs.PrimaryResourceSource() != "r0" {
		t.Fatalf("expected primary rp source r0, got %q", packs.PrimaryResourceSource())
	}
	if packs.IsZero() {
		t.Fatal("non-empty packs reported IsZero")
	}
	if !(regolith.Packs{}).IsZero() {
		t.Fatal("empty packs should report IsZero")
	}
}
