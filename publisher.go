package confelo

// UserID represents User identity
type UserID int64

// Publisher represents publisher role (value object)
type Publisher struct {
	userID UserID
	name   string
}

// NewPublisher creates new Publisher value
func NewPublisher(uid UserID, name string) Publisher {
	return Publisher{uid, name}
}
