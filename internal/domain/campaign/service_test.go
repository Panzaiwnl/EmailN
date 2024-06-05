package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalErrors"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() []Campaign {
	//args := r.Called()
	return nil
}

var (
	newCampaign = contract.NewCampaignDTO{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"teste1@gmail.com"},
	}
	repository = new(repositoryMock)
	service    = Service{Repository: repository}
)

func Test_Create_campaign(t *testing.T) {
	assert2 := assert.New(t)

	newCampaign := contract.NewCampaignDTO{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"teste1@gmail.com"},
	}

	id, err := service.Create(newCampaign)

	assert2.NotNil(id)
	assert2.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {

	newCampaign := contract.NewCampaignDTO{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"teste1@gmail.com"},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		} else if campaign.Content != newCampaign.Content {
			return false
		} else if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)

}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert2 := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert2.NotNil(err)
	assert2.Equal("Name is required", err.Error())

}

func Test_create_ValidadeRepositorySave(t *testing.T) {
	assert2 := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert2.True(errors.Is(internalErrors.ErrInternal, err))
}
