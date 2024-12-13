package roles

type Role int

const (
	Bingo Role = 1 << iota
	Bluey
	Chilli
	Bandit
)

func IsAuthorized(permissions Role, role Role) bool {
	return permissions&role != 0
}
