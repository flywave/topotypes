# topotypes — AGENTS.md

## Project

Go library (`go 1.23`, module `github.com/flywave/topotypes`) defining a type system for 3D topological/geometric objects. Used to serialize/deserialize a heterogeneous family of 3D scene elements (shapes, pipes, catenaries, lights, materials, anchors, geology features, GIM utility components, hydropower tunnels, construction schedules) as JSON.

## Commands

- **Build & test** — standard Go tooling:
  - `go build ./...`
  - `go test ./...`
  - `go vet ./...` (no linter configured beyond what `go vet` provides)
- **Single test**: `go test -run TestMk -v .`
- The only existing test reads `tests/cat.json` (gitignored). Run with `go test` from repo root after placing fixtures in `tests/`.

## Architecture

- **Entrypoint**: `TopoUnmarshal(js []byte)` in `topo.go` — dispatches JSON payload to the correct struct by the `"type"` string field. All leaf types implement `ToposInterface`.
  - "Shape" types → `shape.go` (box, cone, cylinder, sphere, torus, wedge, revolution, pipe)
  - Parametric types (Pipe, Prism, Revol, Catenary, MultiPipe, Parametric, Borehole) → embedded `TopoParametric`
  - Point types (CrossPoint, CrossMultiPoint) → `cross.go`
  - Symbol types (Symbol, SymbolPath, SymbolSurface) → `symbol.go`
  - Light (Spot, Point, Directional, Area) → `light.go`
  - Others (Mask, Feature, Decal, Board, Camera, PipeJoint, SectionLine, Fault, CollapsePillar)
- **Pool of `*Unmarshal()` functions** in dedicated files — each handles post-JSON deserialization of nested interfaces like `Profile`, `Shape`, `Materials`.
- **Profile system** (`profile/`, `profile.go`): Triangle, Rectangle, Circ, Elips, Polygon — used by Pipe, Prism, Revol, Catenary, MultiPipe.
- **Anchor system** (`anchor/`): `TopoAnchor{Id, Position, Link}` and link types — used by Pipe, Catenary, MultiPipe, CrossPoint.
- **Material system** (`material/`): PBR, Lambert, Phong, Base with texture properties. `TopoMaterialMap` accepts either `map[string]material` or `[]material` in JSON (auto-converts array to named map).
- **Shape hierarchy** (three parallel representations): `shape/` (legacy), `base/` (primitive parametric), `gim/` and `hydropower/` (domain-specific parametric). `TopoParametric.UnmarshalJSON` in `parametric.go` dispatches by `"type"` prefix (`"GIM/"`, `"HYDROPOWER/"`, or other → `base/`).
- **GIM subsystem** (`gim/`): subtypes `EC` (electrical — 30+ component types like CableWire, CableTunnel, PipeRow, etc.), `GS`, `GT`. Type strings follow `"GIM/EC/ComponentName"` pattern.
- **Topo4D** (`topo4d/`): construction scheduling — WorkPlans containing WorkSchedules containing WorkTasks.
- **Organizational packages**:
  - `component/` — component abstraction layer
  - `geology/` — Borehole, Fault, CollapsePillar, SectionLine
  - `joint/`, `anchor/` — connection primitives
  - `utils/` — only `StrEquals` (case-insensitive string compare via `strings.EqualFold`)

## JSON Type System

Top-level JSON object uses `"type"` field to select representation:
- `"Shape"` → `TopoShape` (wraps `shape.Box/Cone/etc.`)
- `"Prism"`, `"Revol"` → extrusion/revolution of a profile
- `"Pipe"`, `"MultiSegmentPipe"` → tubular with start/end anchors and profiles
- `"Catenary"` → catenary curve with anchors and slack
- `"CrossPoint"`, `"CrossMultiPoint"` → point with model reference
- `"Symbol"`, `"SymbolPath"`, `"SymbolSurface"` → instance placement
- `"Light"` + `"light"` sub-field → `TopoSpotLight`, `TopoPointLight`, `TopoDirectionalLight`, `TopoAreaLight`
- `"Mask"`, `"Feature"`, `"Decal"`, `"Board"`, `"Camera"`, `"PipeJoint"`
- `"Borehole"`, `"Fault"`, `"CollapsePillar"`, `"SectionLine"`
- `"Parametric"` → `TopoParametric` with sub-`"shape"` having `"type"` like `"GIM/EC/CableWire"` or `"HYDROPOWER/WaterTunnel"`

Type-name ↔ iota mappings are in `topo.go` constants and string-conversion functions.

## Conventions

- String comparisons everywhere use `utils.StrEquals` (case-insensitive `strings.EqualFold`).
- Enum-to-string and string-to-enum functions use parallel `XxxToString` / `StringToXxx` naming.
- `NewTopoXxx()` factory functions set `.Type` via `TopoTypeToString(TOPO_TYPE_XXX)`.
- Custom `UnmarshalJSON` by embedded-anonymous struct pattern (most parametric types): a local `struct` with `Topos` + typed fields, then field-by-field copy to the outer struct.
- `tests/` is gitignored — test fixtures not checked in.
- All packages under `github.com/flywave/topotypes/` — no external test packages.
