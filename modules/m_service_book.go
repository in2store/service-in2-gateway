package modules

import (
	"github.com/in2store/service-in2-gateway/clients/client_in2_book"
)

func CreateBook(req client_in2_book.CreateBookBody, client *client_in2_book.ClientIn2Book) (*client_in2_book.CreateBookResult, error) {
	request := client_in2_book.CreateBookRequest{
		Body: req,
	}
	resp, err := client.CreateBook(request)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}
