package server

import "code.game.com/api"

var apiSevice *api.APIClient

func GetAPIInstance() *api.APIClient {
	if apiSevice == nil {
		apiSevice = api.NewAPIClient("proxy")
	}

	return apiSevice
}
