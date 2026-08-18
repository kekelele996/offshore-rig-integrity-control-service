package store

type HandoverHistory struct {
	roster []string
}

func NewHandoverHistory(roster []string) *HandoverHistory {
	return &HandoverHistory{roster: append([]string(nil), roster...)}
}

func (h *HandoverHistory) Visible(skip string) []string {
	visible := h.roster[:0]
	for _, name := range h.roster {
		if name == skip {
			continue
		}
		visible = append(visible, name)
	}
	for index := len(visible); index < len(h.roster); index++ {
		h.roster[index] = ""
	}
	return visible
}

func (h *HandoverHistory) Roster() []string {
	return h.roster
}
