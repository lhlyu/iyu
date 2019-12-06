package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
	"strconv"
)

func (*dao) InsertCmnt(param *vo.CmntVo) error {
	tx, err := common.DB.Beginx()
	if err != nil {
		common.Ylog.Debug(err)
		return err
	}
	sql := "SELECT COUNT(*) + 1 FROM yu_comment WHERE article_id = ?"
	var floor int
	if err := tx.Get(&floor, sql, param.ArticleId); err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return err
	}
	floors := "#" + strconv.Itoa(floor)
	sql = "INSERT INTO yu_comment(article_id,user_id,`floor`,content,is_check) VALUES(?,?,?,?,?)"
	if _, err := tx.Exec(sql, param.ArticleId, param.UserId, floors, param.Content, param.IsCheck); err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (*dao) UpdateCmnt(param *po.YuComment) error {
	sql := "UPDATE yu_comment SET content = ?,is_check = ?,is_delete = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, param.Content, param.IsCheck, param.IsDelete, param.Id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

func (*dao) GetCmnt(id int) (*po.YuComment, error) {
	sql := "select * from yu_comment where id = ?"
	value := &po.YuComment{}
	if err := common.DB.Get(value, sql, id); err != nil {
		common.Ylog.Debug(err)
		return nil, err
	}
	return value, nil
}

func (*dao) QueryCmntCount(param *vo.CmntVo) (int, error) {
	sql := "SELECT COUNT(*) FROM yu_comment WHERE article_id = ?"
	params := []interface{}{param.ArticleId}
	if param.IsCheck > 0 {
		sql += " and is_check = ?"
		params = append(params, param.IsCheck)
	}
	if param.IsDelete > 0 {
		sql += " and is_delete = ?"
		params = append(params, param.IsDelete)
	}
	sql += " ORDER BY created_at"
	var total int
	if err := common.DB.Get(&total, sql, params...); err != nil {
		common.Ylog.Debug(err)
		return 0, err
	}
	return total, nil
}

func (*dao) QueryCmntPage(page *common.Page, param *vo.CmntVo) ([]*po.YuComment, error) {
	sql := "SELECT * FROM yu_comment WHERE article_id = ?"
	params := []interface{}{param.ArticleId}
	if param.IsCheck > 0 {
		sql += " and is_check = ?"
		params = append(params, param.IsCheck)
	}
	if param.IsDelete > 0 {
		sql += " and is_delete = ?"
		params = append(params, param.IsDelete)
	}
	sql += " ORDER BY created_at limit ?,?"
	params = append(params, page.StartRow, page.PageSize)
	var values []*po.YuComment
	if err := common.DB.Select(&values, sql, params...); err != nil {
		common.Ylog.Debug(err)
		return nil, err
	}
	return values, nil
}
