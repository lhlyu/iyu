package repository

import (
    "github.com/lhlyu/iyu/common"
    "github.com/lhlyu/iyu/errcode"
    "github.com/lhlyu/iyu/repository/po"
)

// add a user
func (*dao) AddUserOne(user *po.YuUser,userInfo *po.YuUserInfo) *repositoryError{
    tx, err := common.DB.Beginx()
    if err != nil {
        return NewRepositoryError("AddUserOne", "", errcode.ERROR, err)
    }
    defer tx.Commit()
    var count int
    sql := "select count(*) from yu_user where third_id = ? and from = ?"
    if err = tx.Select(&count, sql, user.ThirdId,user.From); err != nil {
        return NewRepositoryError("AddUserOne", sql, errcode.ERROR, err)
    }
    if count > 0 {
        // exists
        return NewRepositoryError("AddUserOne", "", errcode.EXISTS_DATA)
    }
    // no exists
    sql = "insert into yu_user(third_id,is_admin,`from`,status) value(?,?,?,?)"
    result, err := tx.Exec(sql, user.ThirdId,user.IsAdmin,user.From,user.Status);
    if err != nil {
        rollerr := tx.Rollback()
        return NewRepositoryError("AddUserOne", sql, errcode.ERROR, err, rollerr)
    }
    sql = "insert into yu_user_info(user_id,avatar_url,user_url,user_name,bio,ip) value(?,?,?,?,?,?)"
    if _,err := tx.Exec(sql,result.LastInsertId(),userInfo.AvatarUrl,userInfo.UserUrl,userInfo.UserName,userInfo.Bio,userInfo.Ip);err != nil{
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
        return NewRepositoryError("UpdateUserOne", sql, errcode.NO_EXISTS_DATA, err)
    }
    newUser := newUsers[0]
    if user.Status > 0 && user.Status != newUser.Status{
        newUser.Status = user.Status
    }
    if user.IsAdmin > 0 && user.IsAdmin != newUser.IsAdmin{
        newUser.IsAdmin = user.IsAdmin
    }
    sql = "update yu_user set is_admin = ?,status = ?,updated_at = now() where id = ?"
    if _, err = tx.Exec(sql, newUser.IsAdmin,newUser.Status,newUser.Id); err != nil {
        rollerr := tx.Rollback()
        return NewRepositoryError("UpdateUserOne", sql, errcode.ERROR, err, rollerr)
    }
    return nil
}

// conditional query users

// get a user
func (*dao) UpdateUserOne(user *po.YuUser) (*po.YuUser,*repositoryError) {

}
