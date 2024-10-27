package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"razvan": {
		AuthToken: "123ABC",
		Username:  "Razvan",
	},
	"marius": {
		AuthToken: "456DEF",
		Username:  "Marius",
	},
	"azalea": {
		AuthToken: "789GHI",
		Username:  "Azalea",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"razvan": {
		Coins:    42,
		Username: "Razvan",
	},
	"marius": {
		Coins:    14,
		Username: "Marius",
	},
	"azalea": {
		Coins:    298,
		Username: "Azalea",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
