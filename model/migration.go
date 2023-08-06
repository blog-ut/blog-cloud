// Package model
// Time : 2023/8/6 12:59
// Author : PTJ
// File : migration
// Software: GoLand
package model

import (
	"fmt"
)

// 数据库迁移
func migration() {
	err := DB.Set("options.table_operation", "charset=utf8mb4").
		AutoMigrate(
			&SysUser{},
		)
	if err != nil {
		fmt.Println("migration err:", err)
	}
}
