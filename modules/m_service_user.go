package modules

import "github.com/in2store/service-in2-gateway/clients/client_in2_user"

func GetUserByUserID(userID uint64, client *client_in2_user.ClientIn2User) (*client_in2_user.User, error) {
	resp, err := client.GetUserByUserID(client_in2_user.GetUserByUserIDRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}
