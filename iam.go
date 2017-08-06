package confelo

// IAM defines iam domain service interface to
// a foreign bounded context
type IAM interface {
	AsPublisher(string) (Publisher, error)
}
