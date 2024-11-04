package service

type TeamDto struct {
	Owner int64
	Name  string
}

func NewTeam(team TeamDto) error {
	return nil
}
