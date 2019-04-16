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
type VolumeTrendHandler struct {  //销售额趋势
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}

func (h VolumeTrendHandler) NewBmVolumeTrendHandler(args ...interface{}) VolumeTrendHandler {
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
	return VolumeTrendHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}

func (h VolumeTrendHandler) VolumeTrend(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	proin :=  BmModel.ProductDimension{}
	//var out    []BmModel.MarketDimension
	var proout []BmModel.ProductDimension
	jso := jsonapiobj.JsResult{}
	var sum float64
	response := map[string]interface{}{
		"status": "",
		"sum": nil,
		"ring":  nil,
		"error":  nil,
	}

	n,_ := strconv.Atoi(r.Header["Ym"][0][:4])
	y,_:= strconv.Atoi(r.Header["Ym"][0][6:8])
	var thisresult []string
	var ringresult []string
	for i:=0;i<13;i++{
		//本年
		//同年同月多个市场
		ps := fmt.Sprintf("%d-%02d", n,y)
		condtmp := bson.M{"ym": ps}

		err := h.db.FindMultiByCondition(&proin,&proout,condtmp,"-sales",0,10)
		if err != nil{
			return 0
		}
		for _,mark:=range proout{
			sum+=mark.Sales
		}
		this := sum
		thisresult = append(thisresult,fmt.Sprintf("%f", this))
		sum=0

		//环比
		ly := y-1
		lps := fmt.Sprintf("%d-%02d", n,ly)
		condtmp = bson.M{"ym": lps}
		err = h.db.FindMultiByCondition(&proin,&proout,condtmp,"-sales",0,10)
		if err != nil{
			return 0
		}
		for _,mark:=range proout{
			sum+=mark.Sales
		}
		ring := sum
		ring = this/ring
		ringresult = append(ringresult,fmt.Sprintf("%f", ring))
		sum=0
		y--
	}
	response["sum"] = thisresult
	response["ring"] = ringresult
	response["status"] = "ok"
	jso.Obj = response
	enc := json.NewEncoder(w)
	enc.Encode(jso.Obj)
	return 0
}

func (h VolumeTrendHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h VolumeTrendHandler) GetHandlerMethod() string {
	return h.Method
}

