package oplog

import (
	"caoguo/internal/api"
	"caoguo/internal/core"
	"caoguo/internal/model"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/tools"
)

// Search 搜索
func Search(c *api.Context, req SearchParam) {
	// 验证token，并获取用户名
	_, bfind := c.GetUserInfo()
	if !bfind {
		c.WriteJSON(message.GetErrorMsg(message.TokenFailure, "please login"))
		return
	}

	if req.Page_num == 0 {
		req.Page_num = 1 // 默认第一页
	}

	startCount := (req.Page_num - 1) * contentCount // 每页10行数据

	var items []model.OpLogTbl
	var result []OpLogInfo

	orm := core.Dao.GetDBr()
	sql := orm.Model(model.OpLogTbl{}).Order("id desc").Offset(startCount).Limit(contentCount)
	sql2 := orm.Model(model.OpLogTbl{})

	if len(req.Text) > 0 {
		sql = sql.Where("attach0 like ?", "%"+req.Text+"%").Or("operator like ?", "%"+req.Text+"%").
			Or("topic like ?", "%"+req.Text+"%").Or("bundle like ?", "%"+req.Text+"%")

		sql2 = sql2.Where("attach0 like ?", "%"+req.Text+"%").Or("operator like ?", "%"+req.Text+"%").
			Or("topic like ?", "%"+req.Text+"%").Or("bundle like ?", "%"+req.Text+"%")
	}
	if !req.Start_time.Time.IsZero() && !req.End_time.Time.IsZero() {
		req.End_time.Time = req.End_time.Time.AddDate(0, 0, 1)
		sql = sql.Where("create_time between ? and ?", req.Start_time.Time, req.End_time.Time)

		sql2 = sql2.Where("create_time between ? and ?", req.Start_time.Time, req.End_time.Time)
	}
	sql.Order("create_time desc").Find(&items)

	var count int64
	sql2.Count(&count)

	if len(items) > 0 {
		for _, v := range items {
			var tmp OpLogInfo
			tmp.Bundle = v.Bundle
			if len(v.Pid) > 0 {
				tmp.Bundle += "(" + v.Pid + ")"
			}
			tmp.Create_time = v.CreateTime
			tmp.Operate_id = v.Attach0
			tmp.Operator = v.Operator
			tmp.Topic = v.Topic
			result = append(result, tmp)
		}
	}

	totalPage := tools.GetTotalPageNum(startCount, int(count)) //总页数
	mk := make(map[string]interface{})
	mk["total_page"] = totalPage
	mk["total_count"] = count
	mk["current_page"] = req.Page_num
	mk["current_data"] = result
	mess := message.GetSuccessMsg(message.NormalMessageID)
	mess.Data = mk
	c.WriteJSON(mess)
}
