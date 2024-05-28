package campaign

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "Campaign x"
	content  = "Body"
	contacts = []string{"email@e.com", "email@gmail.com"}
	now      = time.Now().Add(-time.Minute)
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert2.New(t)

	campaing, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))

}

func Test_NewCampaign_IDIsNotNull(t *testing.T) {

	assert := assert2.New(t)

	campaing, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaing.ID)

}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert2.New(t)

	campaing, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaing.CreatedOn, now)

}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert2.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert2.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert2.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("emails is required", err.Error())
}
