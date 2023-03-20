package keywordsbidratio

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/keywordsbidratio"
)

// Get 查询优词提量系数信息
func Get(clt *core.SDKClient, accessToken string, req *keywordsbidratio.GetRequest) ([]keywordsbidratio.Keyword, error) {
	var resp keywordsbidratio.GetResponse
	if err := clt.Get("v3.0/tools/keywords_bid_ratio/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
