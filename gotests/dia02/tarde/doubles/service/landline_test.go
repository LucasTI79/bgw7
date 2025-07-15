package service_test

import (
	"testing"

	"github.com/bgw7/doubles/repository"
	"github.com/bgw7/doubles/service"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	// given
	repository := repository.LandlineDummyRepository{}
	service := service.NewLandlineService(&repository)
	expectedVersion := "1.0.0"

	// when
	result := service.GetVersion()

	// then
	assert.Equal(t, expectedVersion, result)
}

func TestSearchNameByPhone(t *testing.T) {
	// given
	repository := repository.LandlineStubRepository{}
	service := service.NewLandlineService(&repository)
	expectName := "fulano"
	phoneArg := "123456"

	// when
	result := service.SearchNameByPhone(phoneArg)

	// then
	assert.Equal(t, expectName, result)
}

func TestSearchPhoneByName(t *testing.T) {
	t.Run("test with spy", func(t *testing.T) {
		// given
		repository := repository.LandlineSpyRepository{}
		service := service.NewLandlineService(&repository)

		// when
		service.SearchPhoneByName("")

		// then
		assert.True(t, repository.SearchPhoneByNameCalled)
	})

	t.Run("test with mock", func(t *testing.T) {
		// given
		repository := repository.LandlineMockRepository{}
		service := service.NewLandlineService(&repository)
		expectedResult := "+5511999999999"

		// when
		result := service.SearchPhoneByName("")

		// then
		assert.True(t, repository.SearchPhoneByNameCalled)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("test with fake", func(t *testing.T) {
		// given
		testValues := map[string]string{"Nacho": "123456", "Nico": "234567"}
		repository := repository.LandlineFakeRepository{testValues}
		service := service.NewLandlineService(&repository)
		expectedResult := "123456"
		nameArg := "Nacho"

		// when
		result := service.SearchPhoneByName(nameArg)

		// then
		assert.Equal(t, expectedResult, result)
	})
}
