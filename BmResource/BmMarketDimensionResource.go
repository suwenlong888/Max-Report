package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmMarketDimensionResource struct {
	BmMarketDimensionStorage *BmDataStorage.BmMarketDimensionStorage
}

func (c BmMarketDimensionResource) NewMarketDimensionResource(args []BmDataStorage.BmStorage) BmMarketDimensionResource {
	var cs *BmDataStorage.BmMarketDimensionStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmMarketDimensionStorage" {
			cs = arg.(*BmDataStorage.BmMarketDimensionStorage)
		}	
	}
	return BmMarketDimensionResource{BmMarketDimensionStorage: cs}
}

// FindAll MarketDimensions
func (c BmMarketDimensionResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.BmMarketDimensionStorage.GetAll(r,-1,-1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmMarketDimensionResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmMarketDimensionStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmMarketDimensionResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.MarketDimension)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmMarketDimensionStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmMarketDimensionResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmMarketDimensionStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmMarketDimensionResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.MarketDimension)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmMarketDimensionStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
