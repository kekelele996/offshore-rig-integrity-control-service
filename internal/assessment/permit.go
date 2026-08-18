package assessment

import (
	"sort"
	"strings"
)

type PermitAssessment struct {
	AssetID string
	Window  string
	Zones   []string
	Signals map[string]float64
	Notes   []string
}

type PermitResult struct {
	Score      int
	Band       string
	Review     bool
	Reasons    []string
	Normalized []string
}

func NewPermitAssessment(asset, window string, zones []string, signals map[string]float64) PermitAssessment {
	copied := make([]string, len(zones))
	copy(copied, zones)
	signalCopy := map[string]float64{}
	for key, value := range signals {
		signalCopy[strings.ToLower(strings.TrimSpace(key))] = value
	}
	return PermitAssessment{AssetID: strings.TrimSpace(asset), Window: strings.TrimSpace(window), Zones: copied, Signals: signalCopy}
}

func (a PermitAssessment) NormalizeZones() []string {
	out := make([]string, 0, len(a.Zones))
	seen := map[string]bool{}
	for _, zone := range a.Zones {
		value := strings.ToLower(strings.TrimSpace(zone))
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	sort.Strings(out)
	return out
}

func (a PermitAssessment) Signal(name string) float64 {
	return a.Signals[strings.ToLower(strings.TrimSpace(name))]
}
func (a *PermitAssessment) AddNote(note string) {
	if strings.TrimSpace(note) != "" {
		a.Notes = append(a.Notes, strings.TrimSpace(note))
	}
}

func (a PermitAssessment) Score() int {
	score := 0
	for _, zone := range a.NormalizeZones() {
		score += len(zone) % 4
		if strings.Contains(zone, "restricted") {
			score += 5
		}
		if strings.Contains(zone, "critical") {
			score += 8
		}
	}
	for _, value := range a.Signals {
		if value < 0 {
			score += 10
		} else if value > 90 {
			score += 6
		} else if value > 70 {
			score += 3
		}
	}
	if a.Window == "" {
		score += 4
	}
	if a.AssetID == "" {
		score += 7
	}
	return score
}

func (a PermitAssessment) Result() PermitResult {
	score := a.Score()
	band := "normal"
	if score >= 18 {
		band = "elevated"
	}
	if score >= 30 {
		band = "high"
	}
	reasons := make([]string, 0, 3)
	for _, zone := range a.NormalizeZones() {
		if strings.Contains(zone, "critical") {
			reasons = append(reasons, "critical zone")
		}
	}
	if a.Window == "" {
		reasons = append(reasons, "missing weather window")
	}
	if len(a.NormalizeZones()) == 0 {
		reasons = append(reasons, "no inspection zones")
	}
	return PermitResult{Score: score, Band: band, Review: score >= 18, Reasons: reasons, Normalized: a.NormalizeZones()}
}

func (a PermitAssessment) Merge(other PermitAssessment) PermitAssessment {
	out := NewPermitAssessment(a.AssetID, a.Window, a.Zones, a.Signals)
	if out.AssetID == "" {
		out.AssetID = other.AssetID
	}
	if out.Window == "" {
		out.Window = other.Window
	}
	out.Zones = append(out.Zones, other.Zones...)
	for key, value := range other.Signals {
		out.Signals[key] = value
	}
	out.Notes = append(out.Notes, other.Notes...)
	return out
}

func (a PermitAssessment) Summary() string {
	r := a.Result()
	return strings.Join([]string{a.AssetID, a.Window, r.Band, strings.Join(r.Reasons, "|")}, "/")
}
func (a PermitAssessment) RequiresSupervisor() bool {
	r := a.Result()
	return r.Review || r.Band == "high"
}
func (a PermitAssessment) StableKey() string {
	r := a.Result()
	return strings.Join([]string{a.AssetID, a.Window, strings.Join(r.Normalized, ",")}, "#")
}
func (a PermitAssessment) WithSignal(name string, value float64) PermitAssessment {
	out := NewPermitAssessment(a.AssetID, a.Window, a.Zones, a.Signals)
	out.Signals[strings.ToLower(strings.TrimSpace(name))] = value
	out.Notes = append(out.Notes, a.Notes...)
	return out
}
