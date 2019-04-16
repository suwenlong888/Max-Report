package BmDataStorage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/manyminds/api2go"
)

// BmProductDimensionStorage stores all ProductDimensiones
type BmProductDimensionStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmProductDimensionStorage) NewProductDimensionStorage(args []BmDaemons.BmDaemon) *BmProductDimensionStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmProductDimensionStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmProductDimensionStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.ProductDimension {
	in := BmModel.ProductDimension{}
	var out []BmModel.ProductDimension
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.ProductDimension
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.ProductDimension)
	}
}

// GetOne model
func (s BmProductDimensionStorage) GetOne(id string) (BmModel.ProductDimension, error) {
	in := BmModel.ProductDimension{ID: id}
	model := BmModel.ProductDimension{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("ProductDimension for id %s not found", id)
	return BmModel.ProductDimension{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmProductDimensionStorage) Insert(c BmModel.ProductDimension) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmProductDimensionStorage) Delete(id string) error {
	in := BmModel.ProductDimension{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ProductDimension with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmProductDimensionStorage) Update(c BmModel.ProductDimension) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ProductDimension with id does not exist")
	}

	return nil
}

func (s *BmProductDimensionStorage) Count(req api2go.Request, c BmModel.ProductDimension) int {
	r, _ := s.db.Count(req, &c)
	return r
}
