package services

// NameChecker checks the availability of a name
// Returns true when the name is available, false when unavailable, error if the status cannot be determined
type NameChecker interface {
	Check(name string) (bool, error)
}
