package campaign

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)
import "emailn/internal/contract"

func test_Create_Campaign(t *testing.T) {
	assert := assert2.New(t)
	service := Service{}

	newCampaign := contract.NewCampaignDTO{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"teste1@gmail.com"},
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)

}
