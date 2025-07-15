package repository

type LandlineSpyRepository struct {
	SearchNameByPhoneCalled bool
	SearchPhoneByNameCalled bool
}

// AddEntry implements LandlineRepository.
func (l *LandlineSpyRepository) AddEntry(name string, landline string) error {
	return nil
}

// SearchNameByPhone implements LandlineRepository.
func (l *LandlineSpyRepository) SearchNameByPhone(landline string) string {
	l.SearchNameByPhoneCalled = true
	return ""
}

// SearchPhoneByName implements LandlineRepository.
func (l *LandlineSpyRepository) SearchPhoneByName(name string) string {
	l.SearchPhoneByNameCalled = true
	return ""
}
