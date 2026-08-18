package store

type HandoverHistory struct {
	roster []string
}

func NewHandoverHistory(roster []string) *HandoverHistory {
	return &HandoverHistory{roster: append([]string(nil), roster...)}
}

func (h *HandoverHistory) Visible(skip string) []string {
	visible := make([]string, 0, len(h.roster))
	for _, name := range h.roster {
		if name == skip {
			continue
		}
		visible = append(visible, name)
	}
	return visible
}

func (h *HandoverHistory) Roster() []string {
	return h.roster
}
