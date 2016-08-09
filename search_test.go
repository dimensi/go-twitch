package twitch_test

import (
	"os"
	"testing"

	twitch "github.com/knspriggs/go-twitch"
	"github.com/stretchr/testify/assert"
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
}

func TestSearchChannels(t *testing.T) {
	req := &twitch.SearchChannelsInputType{
		Query: "knspriggs",
	}
	session, err := twitch.NewSession(DefaultURL, twitch.APIV3Header, clientID)
	resp, err := session.SearchChannels(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Total, 1)
	}
}

func TestSearchChannelsEmpty(t *testing.T) {
	req := &twitch.SearchChannelsInputType{
		Query: "agjansa",
	}
	session, err := twitch.NewSession(DefaultURL, twitch.APIV3Header, clientID)
	resp, err := session.SearchChannels(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Total, 0)
	}
}

func TestSearchStreams(t *testing.T) {
	req := &twitch.SearchStreamsInputType{
		Query: "League",
	}
	session, err := twitch.NewSession(DefaultURL, twitch.APIV3Header, clientID)
	resp, err := session.SearchStreams(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Total, 0)
	}
}

func TestSearchStreamsEmpty(t *testing.T) {
	req := &twitch.SearchStreamsInputType{
		Query: "agjansa",
	}
	session, err := twitch.NewSession(DefaultURL, twitch.APIV3Header, clientID)
	resp, err := session.SearchStreams(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, resp.Total, 0)
	}
}

func TestSearchGames(t *testing.T) {
	req := &twitch.SearchGamesInputType{
		Query: "League",
		Type:  "suggest",
		Live:  true,
	}
	session, err := twitch.NewSession(DefaultURL, twitch.APIV3Header, clientID)
	resp, err := session.SearchGames(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, len(resp.Games), 0)
	}
}

func TestSearchGamesEmpty(t *testing.T) {
	req := &twitch.SearchGamesInputType{
		Query: "agjansa",
	}
	session, err := twitch.NewSession(DefaultURL, twitch.APIV3Header, clientID)
	resp, err := session.SearchGames(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, len(resp.Games), 0)
	}
}
