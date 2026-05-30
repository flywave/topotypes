package material

import (
	"testing"

	mst "github.com/flywave/go-mst"
)

func floatEq(t *testing.T, got, want float64, name string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: got %f, want %f", name, got, want)
	}
}

func TestGetPresetSteel(t *testing.T) {
	m := GetPreset(PresetSteel)
	if m == nil {
		t.Fatal("GetPreset(Steel) returned nil")
	}
	if m.Type != "pbr" {
		t.Errorf("Type: got %q, want pbr", m.Type)
	}
	if m.Metallic == nil || *m.Metallic != 1.0 {
		t.Error("Steel should be fully metallic")
	}
	if m.Roughness == nil || *m.Roughness != 0.4 {
		t.Error("Steel roughness mismatch")
	}
	if m.Color == nil || *m.Color != [3]byte{180, 185, 195} {
		t.Error("Steel color mismatch")
	}
}

func TestGetPresetConcrete(t *testing.T) {
	m := GetPreset(PresetConcrete)
	if m == nil {
		t.Fatal("GetPreset(Concrete) returned nil")
	}
	if m.Metallic == nil || *m.Metallic != 0.0 {
		t.Error("Concrete should not be metallic")
	}
	if m.Roughness == nil || *m.Roughness != 0.9 {
		t.Error("Concrete roughness mismatch")
	}
}

func TestGetPresetGlass(t *testing.T) {
	m := GetPreset(PresetGlass)
	if m == nil {
		t.Fatal("GetPreset(Glass) returned nil")
	}
	if m.Transparency != 0.4 {
		t.Errorf("Glass transparency: got %f, want 0.4", m.Transparency)
	}
	if m.Metallic == nil || *m.Metallic != 0.0 {
		t.Error("Glass should not be metallic")
	}
}

func TestGetPresetWood(t *testing.T) {
	m := GetPreset(PresetWood)
	if m == nil {
		t.Fatal("GetPreset(Wood) returned nil")
	}
	if m.Color == nil || *m.Color != [3]byte{165, 125, 85} {
		t.Error("Wood color mismatch")
	}
}

func TestGetPresetGalvanizedSteel(t *testing.T) {
	m := GetPreset(PresetGalvanizedSteel)
	if m == nil {
		t.Fatal("GetPreset(GalvanizedSteel) returned nil")
	}
	if *m.Roughness >= 0.4 {
		t.Errorf("GalvanizedSteel should be smoother, got roughness=%f", *m.Roughness)
	}
}

func TestGetPresetBrick(t *testing.T) {
	m := GetPreset(PresetBrick)
	if m == nil {
		t.Fatal("GetPreset(Brick) returned nil")
	}
	if m.Color == nil || *m.Color != [3]byte{185, 85, 65} {
		t.Error("Brick color mismatch")
	}
}

func TestGetPresetCopper(t *testing.T) {
	m := GetPreset(PresetCopper)
	if m == nil {
		t.Fatal("GetPreset(Copper) returned nil")
	}
	if *m.Metallic != 1.0 {
		t.Error("Copper should be metallic")
	}
	if m.Color == nil || *m.Color != [3]byte{210, 140, 90} {
		t.Error("Copper color mismatch")
	}
}

func TestGetPresetAluminum(t *testing.T) {
	m := GetPreset(PresetAluminum)
	if m == nil {
		t.Fatal("GetPreset(Aluminum) returned nil")
	}
	if *m.Roughness >= 0.3 {
		t.Errorf("Aluminum should be smooth, got roughness=%f", *m.Roughness)
	}
}

func TestGetPresetPlastic(t *testing.T) {
	m := GetPreset(PresetPlastic)
	if m == nil {
		t.Fatal("GetPreset(Plastic) returned nil")
	}
	if *m.Metallic != 0.0 {
		t.Error("Plastic should not be metallic")
	}
}

func TestGetPresetAsphalt(t *testing.T) {
	m := GetPreset(PresetAsphalt)
	if m == nil {
		t.Fatal("GetPreset(Asphalt) returned nil")
	}
	if m.Roughness == nil || *m.Roughness != 0.95 {
		t.Error("Asphalt roughness mismatch")
	}
}

func TestGetPresetRubber(t *testing.T) {
	m := GetPreset(PresetRubber)
	if m == nil {
		t.Fatal("GetPreset(Rubber) returned nil")
	}
	if m.Color == nil || *m.Color != [3]byte{40, 40, 45} {
		t.Error("Rubber color mismatch")
	}
}

func TestGetPresetCeramic(t *testing.T) {
	m := GetPreset(PresetCeramic)
	if m == nil {
		t.Fatal("GetPreset(Ceramic) returned nil")
	}
	if *m.Roughness >= 0.3 {
		t.Errorf("Ceramic should be smooth, got roughness=%f", *m.Roughness)
	}
}

func TestGetPresetUnknown(t *testing.T) {
	m := GetPreset("nonexistent")
	if m != nil {
		t.Error("GetPreset(nonexistent) should return nil")
	}
}

func TestGetPresetReturnsCopy(t *testing.T) {
	m1 := GetPreset(PresetSteel)
	m2 := GetPreset(PresetSteel)
	newColor := [3]byte{0, 0, 0}
	m1.Color = &newColor
	if m2.Color != nil && *m2.Color == [3]byte{0, 0, 0} {
		t.Error("GetPreset should return a copy")
	}
}

func TestPresetNames(t *testing.T) {
	ns := PresetNames()
	if len(ns) == 0 {
		t.Fatal("PresetNames() returned empty")
	}
	for _, n := range ns {
		if GetPreset(n) == nil {
			t.Errorf("GetPreset(%s) returned nil", n)
		}
	}
}

func TestPresetMtlToMeshMtlPBR(t *testing.T) {
	m := GetPreset(PresetSteel)
	mm := MtlToMeshMtl(m)
	if mm == nil {
		t.Fatal("MtlToMeshMtl returned nil")
	}

	pbr, ok := mm.(*mst.PbrMaterial)
	if !ok {
		t.Fatalf("expected *PbrMaterial, got %T", mm)
	}

	if pbr.Metallic != 1.0 {
		t.Errorf("Metallic: got %f", pbr.Metallic)
	}
	if pbr.Roughness != 0.4 {
		t.Errorf("Roughness: got %f", pbr.Roughness)
	}
}

func TestPresetMtlToMeshMtlGlass(t *testing.T) {
	m := GetPreset(PresetGlass)
	mm := MtlToMeshMtl(m)
	if mm == nil {
		t.Fatal("MtlToMeshMtl returned nil")
	}

	pbr, ok := mm.(*mst.PbrMaterial)
	if !ok {
		t.Fatalf("expected *PbrMaterial, got %T", mm)
	}

	if pbr.Transparency != 0.4 {
		t.Errorf("Transparency: got %f, want 0.4", pbr.Transparency)
	}
}

func TestAllPresetsConvertToMeshMtl(t *testing.T) {
	for _, n := range PresetNames() {
		m := GetPreset(n)
		mm := MtlToMeshMtl(m)
		if mm == nil {
			t.Errorf("%s: MtlToMeshMtl returned nil", n)
		}
	}
}

func TestAllPresetsHaveType(t *testing.T) {
	for _, n := range PresetNames() {
		m := GetPreset(n)
		if m.Type == "" {
			t.Errorf("%s: Type is empty", n)
		}
	}
}

func TestAllPresetsHaveColor(t *testing.T) {
	for _, n := range PresetNames() {
		m := GetPreset(n)
		if m.Color == nil || *m.Color == [3]byte{0, 0, 0} {
			t.Errorf("%s: Color should be set", n)
		}
	}
}
