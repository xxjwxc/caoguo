package product

import (
	"caoguo/internal/api"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/internal/service/weixin"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

// GetProductByGpids 通过gpid获取商品信息
func GetProductByGpids(gpids []string) (map[string]*model.ProductTbl, error) {
	orm := core.Dao.GetDBr()
	pds, err := model.ProductTblMgr(orm.DB).GetBatchFromGpid(gpids)
	if err != nil {
		return nil, err
	}

	mp := make(map[string]*model.ProductTbl, len(pds))
	for _, v := range pds {
		mp[v.Gpid] = v
	}

	return mp, nil
}

// GetImages image output
func GetImages(src string) []string {
	var out []string
	tools.JSONEncode(src, &out)
	return out
}

// GetContext 获取图文详情
func GetContext(src string) (img []string, video []string) {
	var out []PContext
	tools.JSONEncode(src, &out)
	for _, v := range out {
		if v.Type == 2 { // 图片
			img = append(img, v.Context)
		} else if v.Type == 3 { // 视频
			video = append(video, v.Context)
		}
	}

	return
}

func getIsFavorite(c *api.Context, gpid string) bool {
	sessionID := c.GetSessionToken()
	mylog.Infof("sessionid:%v", sessionID)
	session := weixin.GetSessionkey(sessionID)
	if len(session.Openid) == 0 {
		return false
	}

	mgr := model.FavoriteTblMgr(core.Dao.GetDBw().DB)
	resp, err := mgr.FetchUniqueIndexByUserGpid(gpid, session.Openid)
	if err != nil { // 没找到
		return false
	}

	return (resp.ID > 0)
}
