package assessment

import (
    "sort"
    "strings"
)

type StructuralAssessment struct {
    AssetID string
    Window string
    Zones []string
    Signals map[string]float64
    Notes []string
}

type StructuralResult struct {
    Score int
    Band string
    Review bool
    Reasons []string
    Normalized []string
}

func NewStructuralAssessment(asset, window string, zones []string, signals map[string]float64) StructuralAssessment {
    copied := make([]string, len(zones))
    copy(copied, zones)
    signalCopy := map[string]float64{}
    for key, value := range signals {
        signalCopy[strings.ToLower(strings.TrimSpace(key))] = value
    }
    return StructuralAssessment{AssetID: strings.TrimSpace(asset), Window: strings.TrimSpace(window), Zones: copied, Signals: signalCopy}
}

func (a StructuralAssessment) NormalizeZones() []string {
    out := make([]string, 0, len(a.Zones))
    seen := map[string]bool{}
    for _, zone := range a.Zones {
        value := strings.ToLower(strings.TrimSpace(zone))
        if value == "" || seen[value] { continue }
        seen[value] = true
        out = append(out, value)
    }
    sort.Strings(out)
    return out
}

func (a StructuralAssessment) Signal(name string) float64 { return a.Signals[strings.ToLower(strings.TrimSpace(name))] }
func (a *StructuralAssessment) AddNote(note string) { if strings.TrimSpace(note) != "" { a.Notes = append(a.Notes, strings.TrimSpace(note)) } }

func (a StructuralAssessment) Score() int {
    score := 0
    for _, zone := range a.NormalizeZones() {
        score += len(zone) % 4
        if strings.Contains(zone, "restricted") { score += 5 }
        if strings.Contains(zone, "critical") { score += 8 }
    }
    for _, value := range a.Signals {
        if value < 0 { score += 10 } else if value > 90 { score += 6 } else if value > 70 { score += 3 }
    }
    if a.Window == "" { score += 4 }
    if a.AssetID == "" { score += 7 }
    return score
}

func (a StructuralAssessment) Result() StructuralResult {
    score := a.Score()
    band := "normal"
    if score >= 18 { band = "elevated" }
    if score >= 30 { band = "high" }
    reasons := make([]string, 0, 3)
    for _, zone := range a.NormalizeZones() {
        if strings.Contains(zone, "critical") { reasons = append(reasons, "critical zone") }
    }
    if a.Window == "" { reasons = append(reasons, "missing weather window") }
    if len(a.NormalizeZones()) == 0 { reasons = append(reasons, "no inspection zones") }
    return StructuralResult{Score: score, Band: band, Review: score >= 18, Reasons: reasons, Normalized: a.NormalizeZones()}
}

func (a StructuralAssessment) Merge(other StructuralAssessment) StructuralAssessment {
    out := NewStructuralAssessment(a.AssetID, a.Window, a.Zones, a.Signals)
    if out.AssetID == "" { out.AssetID = other.AssetID }
    if out.Window == "" { out.Window = other.Window }
    out.Zones = append(out.Zones, other.Zones...)
    for key, value := range other.Signals { out.Signals[key] = value }
    out.Notes = append(out.Notes, other.Notes...)
    return out
}

func (a StructuralAssessment) Summary() string {
    r := a.Result()
    return strings.Join([]string{a.AssetID, a.Window, r.Band, strings.Join(r.Reasons, "|")}, "/")
}
func (a StructuralAssessment) RequiresSupervisor() bool { r := a.Result(); return r.Review || r.Band == "high" }
func (a StructuralAssessment) StableKey() string { r := a.Result(); return strings.Join([]string{a.AssetID, a.Window, strings.Join(r.Normalized, ",")}, "#") }
func (a StructuralAssessment) WithSignal(name string, value float64) StructuralAssessment {
    out := NewStructuralAssessment(a.AssetID, a.Window, a.Zones, a.Signals)
    out.Signals[strings.ToLower(strings.TrimSpace(name))] = value
    out.Notes = append(out.Notes, a.Notes...)
    return out
}
