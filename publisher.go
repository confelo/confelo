package confelo

type UserID int64

// Publisher represents publisher role (value object)
type Publisher struct {
	userID UserID
	name   string
}

func NewPublisher(uid UserID, name string) Publisher {
	return Publisher{uid, name}
}
