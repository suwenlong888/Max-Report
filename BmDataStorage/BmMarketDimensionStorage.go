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

// BmMarketDimensionStorage stores all MarketDimensiones
type BmMarketDimensionStorage struct {
	db *BmMongodb.BmMongodb
}

func (s BmMarketDimensionStorage) NewMarketDimensionStorage(args []BmDaemons.BmDaemon) *BmMarketDimensionStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &BmMarketDimensionStorage{mdb}
}

// GetAll returns the model map (because we need the ID as key too)
func (s BmMarketDimensionStorage) GetAll(r api2go.Request, skip int, take int) []*BmModel.MarketDimension {
	in := BmModel.MarketDimension{}
	var out []BmModel.MarketDimension
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*BmModel.MarketDimension
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.MarketDimension)
	}
}

// GetOne model
func (s BmMarketDimensionStorage) GetOne(id string) (BmModel.MarketDimension, error) {
	in := BmModel.MarketDimension{ID: id}
	model := BmModel.MarketDimension{ID: id}
	err := s.db.FindOne(&in, &model)
	if err == nil {
		return model, nil
	}
	errMessage := fmt.Sprintf("MarketDimension for id %s not found", id)
	return BmModel.MarketDimension{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a model
func (s *BmMarketDimensionStorage) Insert(c BmModel.MarketDimension) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *BmMarketDimensionStorage) Delete(id string) error {
	in := BmModel.MarketDimension{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("MarketDimension with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *BmMarketDimensionStorage) Update(c BmModel.MarketDimension) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("MarketDimension with id does not exist")
	}

	return nil
}

func (s *BmMarketDimensionStorage) Count(req api2go.Request, c BmModel.MarketDimension) int {
	r, _ := s.db.Count(req, &c)
	return r
}
