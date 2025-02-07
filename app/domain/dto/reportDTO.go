package dto

// usersWhoAddedMostSacrifeAndBranchList dto layer
type UsersWhoAddedMostSacrifeAndBranchList struct {
	ID           uint64 `json:"id"`
	Counts       int    `json:"counts"`
	NameLastname int    `json:"nameLastname"`
}
