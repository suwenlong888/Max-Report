package BmResource

import (
	"errors"
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type BmProductDimensionResource struct {
	BmProductDimensionStorage *BmDataStorage.BmProductDimensionStorage
}

func (c BmProductDimensionResource) NewProductDimensionResource(args []BmDataStorage.BmStorage) BmProductDimensionResource {
	var cs *BmDataStorage.BmProductDimensionStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "BmProductDimensionStorage" {
			cs = arg.(*BmDataStorage.BmProductDimensionStorage)
		}	
	}
	return BmProductDimensionResource{BmProductDimensionStorage: cs}
}

// FindAll ProductDimensions
func (c BmProductDimensionResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	result := c.BmProductDimensionStorage.GetAll(r,-1,-1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c BmProductDimensionResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.BmProductDimensionStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c BmProductDimensionResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.ProductDimension)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := c.BmProductDimensionStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c BmProductDimensionResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.BmProductDimensionStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c BmProductDimensionResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(BmModel.ProductDimension)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := c.BmProductDimensionStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
