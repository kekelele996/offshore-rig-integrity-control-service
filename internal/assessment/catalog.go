package assessment

import "sort"

type Catalog struct {
	names  []string
	scores map[string]int
}

func NewCatalog() *Catalog { return &Catalog{scores: map[string]int{}} }
func (c *Catalog) Put(name string, score int) {
	if c.scores == nil {
		c.scores = map[string]int{}
	}
	c.scores[name] = score
	c.names = append(c.names, name)
}
func (c *Catalog) Names() []string {
	out := append([]string(nil), c.names...)
	sort.Strings(out)
	return out
}
func (c *Catalog) Total() int {
	n := 0
	for _, value := range c.scores {
		n += value
	}
	return n
}
func (c *Catalog) Highest() string {
	best := ""
	score := -1
	for name, value := range c.scores {
		if value > score {
			best = name
			score = value
		}
	}
	return best
}
func (c *Catalog) NeedsReview() bool { return c.Total() >= 40 || len(c.scores) >= 6 }
