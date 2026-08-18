package domain

type FindingWindow struct{ scratch []Finding }

func (w *FindingWindow) Snapshot(findings []Finding) []Finding {
	w.scratch = append(w.scratch[:0], findings...)
	return w.scratch
}
