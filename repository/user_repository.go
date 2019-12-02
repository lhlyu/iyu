package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
)

func (*dao) InsertUser(user *po.YuUser) (int, error) {
	sql := "INSERT INTO yu_user (third_id,`from`,avatar_url,user_url,user_name,bio,ip) VALUES(?,?,?,?,?,?,?)"
	result, err := common.DB.Exec(sql, user.ThirdId, user.From, user.AvatarUrl, user.UserUrl, user.UserName, user.Ip)
	if err != nil {
		common.Ylog.Debug(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		common.Ylog.Debug(err)
		return 0, err
	}
	return int(id), nil
}

func (*dao) GetUserById(id int) (*po.YuUser, error) {
	sql := "select * from user where id = ? limit 1"
	result := &po.YuUser{}
	if err := common.DB.Get(result, sql, id); err != nil {
		common.Ylog.Debug(err)
		return nil, err
	}
	return result, nil
}

func (*dao) UpdateUser(user *po.YuUser) error {
	sql := "UPDATE yu_user SET role = ?,`status` = ?,avatar_url = ?,user_name = ?,bio = ?,ip = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, user.Role, user.Status, user.AvatarUrl, user.UserName, user.Bio, user.Ip, user.Id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

func (*dao) GetUsersCount(param *vo.UserParam) (int, error) {
	sql := "SELECT count(*) FROM yu_user WHERE 1 = 1"
	var params []interface{}
	if param.KeyWord != "" {
		keyWord := "%" + param.KeyWord + "%"
		sql += " and (user_name like ? or bio like ?)"
		params = append(params, keyWord, keyWord)
	}
	var result int
	if err := common.DB.Get(&result, sql, params...); err != nil {
		common.Ylog.Debug(err)
		return 0, err
	}
	return result, nil
}

func (*dao) QueryUser(param *vo.UserParam, page *common.Page) ([]*po.YuUser, error) {
	sql := "SELECT * FROM yu_user WHERE 1 = 1"
	var params []interface{}
	if param.KeyWord != "" {
		keyWord := "%" + param.KeyWord + "%"
		sql += " and (user_name like ? or bio like ?)"
		params = append(params, keyWord, keyWord)
	}
	sql += " limit ?,?"
	params = append(params, page.StartRow, page.PageSize)
	var result []*po.YuUser
	if err := common.DB.Select(&result, sql, params...); err != nil {
		common.Ylog.Debug(err)
		return nil, err
	}
	return result, nil
}
