package BmHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
	"reflect"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/julienschmidt/httprouter"
)
type MaxPortionHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}

func (h MaxPortionHandler) NewBmMaxPortionHandler(args ...interface{}) MaxPortionHandler {
	var m *BmMongodb.BmMongodb
	var hm string
	var md string
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmMongodb" {
					m = dm.(*BmMongodb.BmMongodb)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}
	return MaxPortionHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}

func (h MaxPortionHandler) MaxPortion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	in := BmModel.MarketDimension{}
	proin :=  BmModel.ProductDimension{}
	var out    []BmModel.MarketDimension
	var proout []BmModel.ProductDimension
	//var oneout BmModel.MarketDimension
	jso := jsonapiobj.JsResult{}
	response := map[string]interface{}{
		"status": "",
		"sum": nil,
		"same":  nil,
		"ring":  nil,
		"error":  nil,
	}

	n,_ := strconv.Atoi(r.Header["Ym"][0][:4])
	y,_:= strconv.Atoi(r.Header["Ym"][0][6:8])
	//本年
	ps := fmt.Sprintf("%d-%02d", n,y)
	condtmp := bson.M{"ym": ps}
	err := h.db.FindMultiByCondition(&in,&out,condtmp,"-sales",0,1)
	if err != nil{
		return 0
	}
	this := out[0].Sales
	err = h.db.FindMultiByCondition(&proin,&proout,condtmp,"-sales",0,1)
	this = this/proout[0].Sales
	response["sum"] = fmt.Sprintf("%f", this)
	
	//同比 
	ln:=n-1
	lps := fmt.Sprintf("%d-%02d", ln,y)
	if len(r.Header["Market"][0])<=0{
		return 0
	}
	condtmp = bson.M{"ym": lps}
	err = h.db.FindMultiByCondition(&in,&out,condtmp,"-sales",0,1)
	if err != nil{
		return 0
	}
	same := out[0].Sales
	err = h.db.FindMultiByCondition(&proin,&proout,condtmp,"-sales",0,1)
	same = same/proout[0].Sales
	same = this/same
	response["same"] = fmt.Sprintf("%f", same)

	//环比
	ly := y-1
	lps = fmt.Sprintf("%d-%02d", n,ly)
	condtmp = bson.M{"ym": lps}
	err = h.db.FindMultiByCondition(&in,&out,condtmp,"-sales",-1,-1)
	if err != nil{
		return 0
	}
	ring := out[0].Sales
	err = h.db.FindMultiByCondition(&proin,&proout,condtmp,"-sales",0,1)
	ring = ring/proout[0].Sales
	ring = this/ring
	response["ring"] = fmt.Sprintf("%f", ring)

	response["status"] = "ok"
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
	return 0
}

func (h MaxPortionHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h MaxPortionHandler) GetHandlerMethod() string {
	return h.Method
}

