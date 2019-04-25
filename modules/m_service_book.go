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

func GetBookMetaByBookID(bookID uint64, client *client_in2_book.ClientIn2Book) (*client_in2_book.BookMeta, error) {
	request := client_in2_book.GetBookMetaByBookIDRequest{
		BookID: bookID,
	}
	resp, err := client.GetBookMetaByBookID(request)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}

func GetBookRepoByBookID(bookID uint64, client *client_in2_book.ClientIn2Book) (*client_in2_book.BookRepo, error) {
	request := client_in2_book.GetBookRepoByBookIDRequest{
		BookID: bookID,
	}
	resp, err := client.GetBookRepoByBookID(request)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}

func GetBooksMeta(req client_in2_book.GetBooksMetaRequest, client *client_in2_book.ClientIn2Book) (*client_in2_book.GetBooksMetaResult, error) {
	resp, err := client.GetBooksMeta(req)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}
