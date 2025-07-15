package repository

type LandlineDummyRepository struct {
}

// AddEntry implements LandlineRepository.
func (l *LandlineDummyRepository) AddEntry(name string, landline string) error {
	return nil
}

// SearchNameByPhone implements LandlineRepository.
func (l *LandlineDummyRepository) SearchNameByPhone(landline string) string {
	return ""
}

// SearchPhoneByName implements LandlineRepository.
func (l *LandlineDummyRepository) SearchPhoneByName(name string) string {
	return ""
}
