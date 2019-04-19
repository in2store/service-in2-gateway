package modules

import "github.com/in2store/service-in2-gateway/clients/client_in2_auth"

func GetTokens(req client_in2_auth.GetTokensRequest, client *client_in2_auth.ClientIn2Auth) (*client_in2_auth.GetTokensResult, error) {
	tokens, err := client.GetTokens(req)
	if err != nil {
		return nil, err
	}
	return &tokens.Body, nil
}
