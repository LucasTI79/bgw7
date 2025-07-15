package repository

type LandlineStubRepository struct {
}

// AddEntry implements LandlineRepository.
func (l *LandlineStubRepository) AddEntry(name string, landline string) error {
	return nil
}

// SearchNameByPhone implements LandlineRepository.
func (l *LandlineStubRepository) SearchNameByPhone(landline string) string {
	return "fulano"
}

// SearchPhoneByName implements LandlineRepository.
func (l *LandlineStubRepository) SearchPhoneByName(name string) string {
	return "+5511999999999"
}
