package BmFactory

import (
	"github.com/PharbersDeveloper/Max-Report/BmDataStorage"
	"github.com/PharbersDeveloper/Max-Report/BmHandler"
	"github.com/PharbersDeveloper/Max-Report/BmModel"
	"github.com/PharbersDeveloper/Max-Report/BmResource"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	//"github.com/PharbersDeveloper/Max-Report/BmMiddleware"
)

type BmTable struct{}

var BLACKMIRROR_MODEL_FACTORY = map[string]interface{}{
	"BmProductDimension":                BmModel.ProductDimension{},
	"BmMarketDimension":                 BmModel.MarketDimension{},
}

var BLACKMIRROR_RESOURCE_FACTORY = map[string]interface{}{
	"BmProductDimensionResource":               BmResource.BmProductDimensionResource{},
	"BmMarketDimensionResource":                BmResource.BmMarketDimensionResource{},
}

var BLACKMIRROR_STORAGE_FACTORY = map[string]interface{}{
	"BmProductDimensionStorage":               BmDataStorage.BmProductDimensionStorage{},
	"BmMarketDimensionStorage":                BmDataStorage.BmMarketDimensionStorage{},
}

var BLACKMIRROR_MIDDLEWARE_FACTORY = map[string]interface{}{
	//"BmCheckTokenMiddleware": BmMiddleware.BmCheckTokenMiddleware{},
}

var BLACKMIRROR_DAEMON_FACTORY = map[string]interface{}{
	 "BmMongodbDaemon": BmMongodb.BmMongodb{},
	 "BmRedisDaemon":   BmRedis.BmRedis{},
}

var BLACKMIRROR_FUNCTION_FACTORY = map[string]interface{}{
	"BmMarketScopeHandler":     	   BmHandler.MarketScopeHandler{},
	"BmProductCountHandler":     	   BmHandler.ProductCountHandler{},
	"BmMarketRatioHandler":     	   BmHandler.MarketRatioHandler{},
	"BmMaxPortionHandler":     	   	   BmHandler.MaxPortionHandler{},
	"BmMarketScopeTrendHandler":       BmHandler.MarketScopeTrendHandler{},
	"BmProductPortionHandler":     	   BmHandler.ProductPortionHandler{},
	"BmVolumeTrendHandler":     	   BmHandler.VolumeTrendHandler{},
	"BmVolumeCmpHandler":     	   	   BmHandler.VolumeCmpHandler{},
	"BmPortionCmpHandler":     	   	   BmHandler.PortionCmpHandler{},
	"BmPortionTrendHandler":     	   BmHandler.PortionTrendHandler{},
}


func (t BmTable) GetModelByName(name string) interface{} {
	return BLACKMIRROR_MODEL_FACTORY[name]
}

func (t BmTable) GetResourceByName(name string) interface{} {
	return BLACKMIRROR_RESOURCE_FACTORY[name]
}

func (t BmTable) GetStorageByName(name string) interface{} {
	return BLACKMIRROR_STORAGE_FACTORY[name]
}

func (t BmTable) GetDaemonByName(name string) interface{} {
	return BLACKMIRROR_DAEMON_FACTORY[name]
}

func (t BmTable) GetFunctionByName(name string) interface{} {
	return BLACKMIRROR_FUNCTION_FACTORY[name]
}

func (t BmTable) GetMiddlewareByName(name string) interface{} {
	return BLACKMIRROR_MIDDLEWARE_FACTORY[name]
}