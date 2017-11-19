package files

// Tracker provides a way to track a bounded slice of file paths
type Tracker struct {
	s []string
}

// NewTracker returns a tracker with the specified length
func NewTracker(length int) *Tracker {
	t := &Tracker{}
	t.s = make([]string, length)
	return t
}

// Add appends a path to the tracker
func (t *Tracker) Add(path string) {
	t.s = t.s[1:]
	t.s = append(t.s, path)
}

// In determines if the path exists in the tracker
func (t *Tracker) In(path string) bool {
	for _, item := range t.s {
		if item == path {
			return true
		}
	}

	return false
}
