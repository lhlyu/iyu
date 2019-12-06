package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
	"strconv"
)

func (*dao) InsertPost(param *vo.PostVo) error {
	tx, err := common.DB.Beginx()
	if err != nil {
		common.Ylog.Debug(err)
		return err
	}
	sql := "SELECT COUNT(*) + 1 FROM yu_post WHERE comment_id = ?"
	var floor int
	if err := tx.Get(&floor, sql, param.CommentId); err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return err
	}
	sql = "SELECT `floor` FROM yu_comment WHERE id = ?"
	var cmntFloor string
	if err := tx.Get(&cmntFloor, sql, param.CommentId); err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return err
	}
	floors := cmntFloor + "-" + strconv.Itoa(floor)
	sql = "INSERT INTO yu_post(comment_id,user_id,`floor`,at_id,at_floor,content,is_check) VALUES(?,?,?,?,?,?,?)"
	if _, err := tx.Exec(sql, param.CommentId, param.UserId, floors, param.AtId, param.AtFloor, param.Content, param.IsCheck); err != nil {
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

func (*dao) UpdatePost(param *po.YuPost) error {
	sql := "UPDATE yu_post SET content = ?,is_check = ?,is_delete = ?,updated_at = ? WHERE id = ?"
	if _, err := common.DB.Exec(sql, param.Content, param.IsCheck, param.IsDelete, param.Id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

func (*dao) GetPost(id int) (*po.YuPost, error) {
	sql := "select * from yu_post where id = ?"
	value := &po.YuPost{}
	if err := common.DB.Get(value, sql, id); err != nil {
		common.Ylog.Debug(err)
		return nil, err
	}
	return value, nil
}

func (*dao) QueryPostCount(param *vo.PostVo) (int, error) {
	sql := "SELECT COUNT(*) FROM yu_post WHERE is_check = 2 AND is_delete = 1 AND comment_id = ? ORDER BY created_at"
	var total int
	if err := common.DB.Get(&total, sql, param.CommentId); err != nil {
		common.Ylog.Debug(err)
		return 0, err
	}
	return total, nil
}

func (*dao) QueryPostPage(page *common.Page, param *vo.PostVo) ([]*po.YuPost, error) {
	sql := "SELECT * FROM yu_post WHERE is_check = 2 AND is_delete = 1 AND article_id = ? ORDER BY created_at limit ?,?"
	var values []*po.YuPost
	if err := common.DB.Select(&values, sql, param.CommentId, page.StartRow, page.PageSize); err != nil {
		common.Ylog.Debug(err)
		return nil, err
	}
	return values, nil
}
