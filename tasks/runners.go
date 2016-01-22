package tasks
import "github.schq.secious.com/Logrhythm/easy/conf"


type Feedback interface {
	// IsDone() should return true if the feedback is complete due
	// to the completion of the Execution.
	IsDone() bool

	// Progress() provides update messages as execution is carried out.
	// The result of calling this function provides a human readable
	// update of progress, followed by how much progress has been made
	// in between [0.0, 1.0], negative value mean 'unknown' amount of
	// progress, and finally if an error has occured this method
	// provides the error.
	Progress() (message string, completedPct float32, failure error)
}

// Gatherers should not require a constructor, and a pointer literal
// should suffice.
type Gatherer interface {
	// Usage explains the function of this gatherer.  Used when listing
	// the Gatherers in the entire system.
	Usage() string

	// Exec() accepts a few data points in order to direct the processing
	// carried out by this Gatherer.
	Exec(*conf.Conf) chan Feedback
}

