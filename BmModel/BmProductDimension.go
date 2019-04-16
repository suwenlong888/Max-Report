package BmModel

import (
	//"errors"
	//"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type ProductDimension struct {
	ID						string        `json:"-"`
	Id_						bson.ObjectId `json:"-" bson:"_id"`
	Company_ID				string	`json:"company-id" bson:"company-id"`
	Market					string	`json:"market" bson:"market"`
	Ym						string	`json:"ym" bson:"ym"`
	Sales					float64	`json:"sales" bson:"sales"`
	Units					float64	`json:"units" bson:"units"`
	Product_Name			string	`json:"product-name" bson:"product-name"`
	Province        		string  `json:"province" bson:"province"`
	City           			string  `json:"city" bson:"city"`
	Corp_Name				string  `json:"corp_name" bson:"corp_name"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a ProductDimension) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *ProductDimension) SetID(id string) error {
	a.ID = id
	return nil
}
func (a *ProductDimension) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
	/*
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "ids":
			r := make(map[string]interface{})
			var ids []bson.ObjectId
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		case "scenario-id":
			rst[k] = v[0]
		}
	}
	return rst
	*/
}
