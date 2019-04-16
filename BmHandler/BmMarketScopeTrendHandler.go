package BmHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
	"reflect"
	"strconv"
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/julienschmidt/httprouter"
)
type MarketScopeTrendHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}

func (h MarketScopeTrendHandler) NewBmMarketScopeTrendHandler(args ...interface{}) MarketScopeTrendHandler {
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
	return MarketScopeTrendHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}

func (h MarketScopeTrendHandler) MarketScopeTrend(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	in := BmModel.MarketDimension{}
	var results []string
	var out []BmModel.MarketDimension
	var oneout BmModel.MarketDimension
	jso := jsonapiobj.JsResult{}
	var sum float64
	response := map[string]interface{}{
		"status": "",
		"result": nil,
		"error":  nil,
	}

	n,_ := strconv.Atoi(r.Header["Ym"][0][:4])
	y,_:= strconv.Atoi(r.Header["Ym"][0][5:7])

	for i := 0;i<13;i++{
		//同年同月多个市场
		ps := fmt.Sprintf("%d-%02d", n,y)
		condtmp := bson.M{"ym": ps}
		err := h.db.FindMultiByCondition(&in,&out,condtmp,"-sales",0,10)
		if err != nil{
			return 0
		}
		for _,mark:=range out{
			sum+=mark.Sales
		}
		cond := bson.M{"ym": ps,"market":r.Header["Market"][0]}
		err = h.db.FindOneByCondition(&in,&oneout,cond)
		if err != nil{
			return 0
		}
		sale := oneout.Sales
		this := sale/sum
		sum=0
		tmpresult:=fmt.Sprintf("%f", this)
		results=append(results,tmpresult)
		y--
	}

	response["status"] = "ok"
	response["result"] = results
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
	return 0
}

func (h MarketScopeTrendHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h MarketScopeTrendHandler) GetHandlerMethod() string {
	return h.Method
}


