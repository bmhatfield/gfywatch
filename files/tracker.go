package files

// Tracker provides a way to track a bounded slice of file paths
type Tracker []string

// NewTracker returns a tracker with the specified length
func NewTracker(length int) *Tracker {
	t := make(Tracker, length)
	return &t
}

// Add appends a path to the tracker
func (t *Tracker) Add(path string) {
	tracker := *t
	tracker = tracker[1:]
	tracker = append(tracker, path)
	t = &tracker
}

// In determines if the path exists in the tracker
func (t *Tracker) In(path string) bool {
	for _, item := range *t {
		if item == path {
			return true
		}
	}

	return false
}
