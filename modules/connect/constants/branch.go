package constants

type Branch interface {
	// GetCommit returns the Commit field.
	GetCommit() Commit

	// GetName returns the Name field if it's non-nil, zero value otherwise.
	GetName() string

	// GetProtected returns the Protected field if it's non-nil, zero value otherwise.
	GetProtected() bool
}
