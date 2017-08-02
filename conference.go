package confelo

type ConferenceID int64

type CreateConferenceCommand struct {
	Name string
}

type Conference struct {
	Name string
}
