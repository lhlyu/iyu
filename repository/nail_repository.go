package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
)

// add a nail
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
