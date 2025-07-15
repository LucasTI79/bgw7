package repository

type landlineRepository struct {
}

type LandlineRepository interface {
	SearchPhoneByName(name string) string
	SearchNameByPhone(landline string) string
	AddEntry(name, landline string) error
}

func NewRepository() LandlineRepository {
	return &LandlineDummyRepository{}
}
