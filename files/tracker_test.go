package files

import (
	"testing"
)

func TestTracker(t *testing.T) {
	// New, empty Tracker
	tracker := NewTracker(3)

	// Tracker should in fact be empty
	if tracker.In("forks") {
		t.Logf("%+v", tracker)
		t.Error("forks in tracker unexpectedly")
	}

	tracker.Add("forks")

	// Tracker should now be ["", "", "forks"]
	if !tracker.In("forks") {
		t.Logf("%+v", tracker)
		t.Error("forks not in tracker")
	}

	if tracker.s[0] != "" && tracker.s[1] != "" {
		t.Logf("%+v", tracker)
		t.Error("first and second elements of tracker not emptystring")
	}

	tracker.Add("forks1")
	tracker.Add("forks2")
	tracker.Add("forks3")

	// Tracker should now be ["forks1", "forks2", "forks3"]
	// Original "forks" should be rotated out
	if tracker.In("forks") {
		t.Logf("%+v", tracker)
		t.Error("forks in tracker unexpectedly")
	}
}
