package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/po"
)

// add a user
func (*dao) AddUserOne(user *po.YuUser) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("AddUserOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	var count int
	sql := "select count(*) from yu_user where third_id = ? and `from` = ?"
	if err = tx.Get(&count, sql, user.ThirdId, user.From); err != nil {
		return NewRepositoryError("AddUserOne", sql, errcode.ERROR, err)
	}
	if count > 0 {
		// exists
		return NewRepositoryError("AddUserOne", "", errcode.EXISTS_DATA)
	}
	// no exists
	sql = "INSERT INTO `lhlyu_blog`.`yu_user` (`third_id`, `is_admin`, `from`, `status` , `avatar_url`, `user_url`, `user_name`, `bio`, `ip`) VALUES (?,?,?,?,?,?,?,?,?);"
	_, err = tx.Exec(sql, user.ThirdId, user.IsAdmin, user.From, user.Status, user.AvatarUrl, user.UserUrl, user.UserName, user.Bio, user.Ip)
	if err != nil {
		rollerr := tx.Rollback()
		return NewRepositoryError("AddUserOne", sql, errcode.ERROR, err, rollerr)
	}
	return nil
}

// update user
func (*dao) UpdateUserOne(user *po.YuUser) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("UpdateUserOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newUsers := []*po.YuUser{}
	sql := "select * from yu_user where id = ? limit 1"
	if err = tx.Select(&newUsers, sql, user.Id); err != nil {
		return NewRepositoryError("UpdateUserOne", sql, errcode.ERROR, err)
	}
	if len(newUsers) == 0 {
		// 不存在
		return NewRepositoryError("UpdateUserOne", sql, errcode.NO_EXISTS_DATA)
	}
	newUser := newUsers[0]
	if user.Status > 0 && user.Status != newUser.Status {
		newUser.Status = user.Status
	}
	if user.IsAdmin > 0 && user.IsAdmin != newUser.IsAdmin {
		newUser.IsAdmin = user.IsAdmin
	}
	if user.AvatarUrl != "" && user.AvatarUrl != newUser.AvatarUrl {
		newUser.AvatarUrl = user.AvatarUrl
	}
	if user.UserUrl != "" && user.UserUrl != newUser.UserUrl {
		newUser.UserUrl = user.UserUrl
	}
	if user.UserName != "" && user.UserName != newUser.UserName {
		newUser.UserName = user.UserName
	}
	if user.Bio != "" && user.Bio != newUser.Bio {
		newUser.Bio = user.Bio
	}
	if user.Ip != "" && user.Ip != newUser.Ip {
		newUser.Ip = user.Ip
	}
	sql = "update yu_user set is_admin = ?,status = ?,avatar_url = ?,user_url = ?,user_name = ?,bio = ?,ip = ?,updated_at = now() where id = ?"
	if _, err = tx.Exec(sql, newUser.IsAdmin, newUser.Status, newUser.AvatarUrl, newUser.UserUrl, newUser.UserName, newUser.Bio, newUser.Ip, newUser.Id); err != nil {
		rollerr := tx.Rollback()
		return NewRepositoryError("UpdateUserOne", sql, errcode.ERROR, err, rollerr)
	}
	return nil
}

// conditional query users
func (*dao) QueryUser(user *po.YuUser, page *common.Page) ([]*po.YuUser, *repositoryError) {
	countSql := "select count(*) from yu_user where 1 = 1 "
	var params []interface{}
	whstr := ""
	if user == nil {
		user = &po.YuUser{}
	}
	if user.Id > 0 {
		whstr += "and id = ? "
		params = append(params, user.Id)
	}
	if user.IsAdmin > 0 {
		whstr += "and is_admin = ? "
		params = append(params, user.IsAdmin)
	}
	if user.Status > 0 {
		whstr += "and status = ? "
		params = append(params, user.Status)
	}
	if user.From > 0 {
		whstr += "and `from` = ? "
		params = append(params, user.From)
	}
	if user.UserName != "" {
		whstr += "and user_name like ? "
		params = append(params, "%"+user.UserName+"%")
	}
	if user.Bio != "" {
		whstr += "and bio like ? "
		params = append(params, "%"+user.Bio+"%")
	}
	var total int
	if err := common.DB.Get(&total, countSql+whstr, params...); err != nil {
		return nil, NewRepositoryError("QueryUser", countSql+whstr, errcode.ERROR, params, err)
	}
	if total == 0 {
		return nil, NewRepositoryError("QueryUser", "", errcode.NO_EXISTS_DATA)
	}
	page.SetTotal(total)
	sql := "select * from yu_user where 1 = 1 " + whstr + "order by is_admin desc,status,`from`,updated_at desc,created_at desc limit ?,?"
	params = append(params, page.StartRow, page.PageSize)
	newUsers := []*po.YuUser{}
	if err := common.DB.Select(&newUsers, sql, params...); err != nil {
		return nil, NewRepositoryError("QueryUser", sql, errcode.ERROR, params, err)
	}
	if len(newUsers) == 0 {
		// 不存在
		return nil, NewRepositoryError("QueryUser", "", errcode.NO_EXISTS_DATA)
	}
	return newUsers, nil
}

// get a user
func (*dao) GetUser(user *po.YuUser) (*po.YuUser, *repositoryError) {
	sql := "select * from yu_user where id = ? limit 1"
	newUsers := []*po.YuUser{}
	if err := common.DB.Select(&newUsers, sql, user.Id); err != nil {
		return nil, NewRepositoryError("GetUser", sql, errcode.ERROR, err)
	}
	if len(newUsers) == 0 {
		// 不存在
		return nil, NewRepositoryError("GetUser", sql, errcode.NO_EXISTS_DATA)
	}
	return newUsers[0], nil
}
