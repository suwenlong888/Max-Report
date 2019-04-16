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
type ProductPortionHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}

func (h ProductPortionHandler) NewBmProductPortionHandler(args ...interface{}) ProductPortionHandler {
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
	return ProductPortionHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}

func (h ProductPortionHandler) ProductPortion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	proin :=  BmModel.ProductDimension{}
	var proout []BmModel.ProductDimension
	jso := jsonapiobj.JsResult{}
	var sum float64
	response := map[string]interface{}{
		"status": "",
		"first": nil,
		"second":  nil,
		"third":  nil,
		"forth":  nil,
		"fifth":  nil,
		"others":  nil,
		"error":  nil,
	}

	n,_ := strconv.Atoi(r.Header["Ym"][0][:4])
	y,_:= strconv.Atoi(r.Header["Ym"][0][6:8])

	//同年同月多个市场
	ps := fmt.Sprintf("%d-%02d", n,y)
	condtmp := bson.M{"ym": ps}
	err := h.db.FindMultiByCondition(&proin,&proout,condtmp,"-sales",-1,-1)
	if err != nil{
		return 0
	}
	for _,mark:=range proout{
		sum+=mark.Sales
	}
	response["first"] = fmt.Sprintf("%f", proout[0].Sales/sum)
	response["second"] = fmt.Sprintf("%f", proout[1].Sales/sum)
	response["third"] = fmt.Sprintf("%f", proout[2].Sales/sum)
	response["forth"] = fmt.Sprintf("%f", proout[3].Sales/sum)
	response["fifth"] = fmt.Sprintf("%f", proout[4].Sales/sum)
	other := 1-proout[0].Sales/sum-proout[1].Sales/sum-proout[2].Sales/sum-proout[3].Sales/sum-proout[4].Sales/sum
	response["others"] = fmt.Sprintf("%f", other)
	response["status"] = "ok"
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
	return 0
}

func (h ProductPortionHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h ProductPortionHandler) GetHandlerMethod() string {
	return h.Method
}

