package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
)

/**
添加一个nail 【事务】
1. 判断是否已经存在
2. 如果已被删除，重新改为启用
3. 如果不存在，添加
*/
func (*dao) AddNailOne(name, color string) {
	result, err := common.DB.Exec("insert into yu_nail(name,color) value(?,?)", name, color)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	lastInsertId, err1 := result.LastInsertId()
	rowsAffected, err2 := result.RowsAffected()
	fmt.Printf("lastInsertId = %d,err1 = %v\n", lastInsertId, err1)
	fmt.Printf("rowsAffected = %d,err2 = %v\n", rowsAffected, err2)
}

/**
修改一个nail 【事务】
1. 判断是否已经存在
2. 修改
*/
func (*dao) UpdateNailOne(name, color string) {
	result, err := common.DB.Exec("insert into yu_nail(name,color) value(?,?)", name, color)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	lastInsertId, err1 := result.LastInsertId()
	rowsAffected, err2 := result.RowsAffected()
	fmt.Printf("lastInsertId = %d,err1 = %v\n", lastInsertId, err1)
	fmt.Printf("rowsAffected = %d,err2 = %v\n", rowsAffected, err2)
}

// 查询所有nail
func (*dao) QueryNail(name, color string) {
	result, err := common.DB.Exec("insert into yu_nail(name,color) value(?,?)", name, color)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	lastInsertId, err1 := result.LastInsertId()
	rowsAffected, err2 := result.RowsAffected()
	fmt.Printf("lastInsertId = %d,err1 = %v\n", lastInsertId, err1)
	fmt.Printf("rowsAffected = %d,err2 = %v\n", rowsAffected, err2)
}
