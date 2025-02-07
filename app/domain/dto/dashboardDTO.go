package dto

// Dashboard dto
type Dashboard struct {
	ID      uint64
	AdSoyad string
	Telefon string
	Adres   string
}

// CharitableWhoAddedMostSacrife dto
type CharitableWhoAddedMostSacrife struct {
	UserID       uint64
	KisiID       uint64
	Counts       uint64
	NameLastname string
}

// UsersWhoAddedMostSacrifeAndBranch  En cok kurban ekleyen hocamiz ve subemiz
type UsersWhoAddedMostSacrifeAndBranch struct {
	SacrifeUserID uint64
	BranchID      uint64
	Counts        int
	FirstName     string
	LastName      string
	Email         string
	Branch        string
}
