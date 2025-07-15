package repository

type LandlineMockRepository struct {
	SearchNameByPhoneCalled bool
	SearchPhoneByNameCalled bool
}

// AddEntry implements LandlineRepository.
func (l *LandlineMockRepository) AddEntry(name string, landline string) error {
	return nil
}

// SearchNameByPhone implements LandlineRepository.
func (l *LandlineMockRepository) SearchNameByPhone(landline string) string {
	l.SearchNameByPhoneCalled = true
	return "fulano"
}

// SearchPhoneByName implements LandlineRepository.
func (l *LandlineMockRepository) SearchPhoneByName(name string) string {
	l.SearchPhoneByNameCalled = true
	return "+5511999999999"
}
