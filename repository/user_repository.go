package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
)

func (d *dao) QueryUser(id ...int) []*po.YuUser {
	sql := "SELECT * FROM yu_user"
	var params []interface{}
	if len(id) > 0 {
		marks := d.createQuestionMarks(len(id))
		params = d.intConvertToInterface(id)
		sql += fmt.Sprintf(" where id in (%s)", marks)
	}
	var values []*po.YuUser
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err)
		return nil
	}
	return values
}

func (d *dao) InsertUser(user *vo.UserEditParam) (int, error) {
	sql := "INSERT INTO yu_user (third_id,`from`,avatar_url,user_url,user_name,bio,ip) VALUES(?,?,?,?,?,?,?)"
	result, err := common.DB.Exec(sql, user.ThirdId, user.From, user.AvatarUrl, user.UserUrl, user.UserName, user.Ip)
	if err != nil {
		d.Error(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		d.Error(err)
		return 0, err
	}
	return int(id), nil
}

func (d *dao) GetUserById(id int) (*po.YuUser, error) {
	sql := "select * from yu_user where id = ? limit 1"
	result := &po.YuUser{}
	if err := common.DB.Get(result, sql, id); err != nil {
		d.Error(err)
		return nil, err
	}
	return result, nil
}

func (d *dao) UpdateUser(user *po.YuUser) error {
	sql := "UPDATE yu_user SET role = ?,`status` = ?,avatar_url = ?,user_name = ?,bio = ?,ip = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, user.Role, user.Status, user.AvatarUrl, user.UserName, user.Bio, user.Ip, user.Id); err != nil {
		d.Error(err)
		return err
	}
	return nil
}

func (d *dao) GetUsersCount(param *vo.UserParam) (int, error) {
	sql := "SELECT count(*) FROM yu_user WHERE 1 = 1"
	var params []interface{}
	if param.KeyWord != "" {
		keyWord := "%" + param.KeyWord + "%"
		sql += " and (user_name like ? or bio like ?)"
		params = append(params, keyWord, keyWord)
	}
	if param.Id > 0 {
		sql += " and id = ?"
		params = append(params, param.Id)
	}
	var result int
	if err := common.DB.Get(&result, sql, params...); err != nil {
		d.Error(err)
		return 0, err
	}
	return result, nil
}

func (d *dao) QueryUserPage(param *vo.UserParam, page *common.Page) ([]int, error) {
	sql := "SELECT id FROM yu_user WHERE 1 = 1"
	var params []interface{}
	if param.KeyWord != "" {
		keyWord := "%" + param.KeyWord + "%"
		sql += " and (user_name like ? or bio like ?)"
		params = append(params, keyWord, keyWord)
	}
	if param.Id > 0 {
		sql += " and id = ?"
		params = append(params, param.Id)
	}
	sql += " order by `status`,created_at desc limit ?,?"
	params = append(params, page.StartRow, page.PageSize)
	var result []int
	if err := common.DB.Select(&result, sql, params...); err != nil {
		d.Error(err)
		return nil, err
	}
	return result, nil
}
