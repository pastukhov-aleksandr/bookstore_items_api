package services

import (
	"github.com/pastukhov-aleksandr/bookstore_items_api/src/domain/items"
	"github.com/pastukhov-aleksandr/bookstore_utils-go/rest_errors"
	"net/http"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	//Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	//item := items.Item{Id: id}
	//
	//if err := item.Get(); err != nil {
	//	return nil, err
	//}
	//return &item, nil
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implement", nil)
}

//func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
//	dao := items.Item{}
//	return dao.Search(query)
//}
