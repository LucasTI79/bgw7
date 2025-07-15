package repository

type LandlineFakeRepository struct {
	// key: Nome, value: Phone
	DB map[string]string
}

func (m *LandlineFakeRepository) SearchPhoneByName(name string) string {
	return m.DB[name]
}

func (m *LandlineFakeRepository) SearchNameByPhone(phone string) string {
	for key, value := range m.DB {
		if value == phone {
			return key
		}
	}
	return ""
}

func (m *LandlineFakeRepository) AddEntry(name, phone string) error {
	if m.DB == nil {
		m.DB = map[string]string{}
	}
	m.DB[name] = phone
	return nil
}
