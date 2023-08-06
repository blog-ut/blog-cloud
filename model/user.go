// Package model
// Time : 2023/8/6 12:56
// Author : PTJ
// File : user
// Software: GoLand
package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	Username string `gorm:"not null;uniqueIndex;comment:用户名"`
	NickName string `gorm:"not null;comment:昵称"`
	Password string `gorm:"not null;comment:密码"`
	Intro    string `gorm:"comment:简介"`
	Avatar   string `gorm:"not null;comment:头像"`
	Phone    string `gorm:"comment:电话号码"`
	Email    string `gorm:"comment:邮箱"`
	Address  string `gorm:"comment:地址"`
	IsEnable int    `gorm:"not null;comment:是否禁用 0无效 1有效"`
	IsAdmin  int    `gorm:"not null;comment:是否是管理员 0无效 1有效"`
}

const PasswordCost = 12 //密码加密难度

// SetPassword 加密密码
func (user *SysUser) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *SysUser) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
