package entity

type Error struct {
	Error  error
	Cause  string
	Status int
}

func (err Error) IsValid() {
}
