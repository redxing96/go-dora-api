/*
 * @Description: 密码工具类
 * @Author: redxing96@163.com
 * @Date: 2025-06-24 20:45:51
 * @LastEditTime: 2025-06-24 20:45:58
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/common/password.go
 */
package common

import "golang.org/x/crypto/bcrypt"

// GenerateFromPassword 函数用于生成密码的哈希值
func GenerateFromPassword(password string) string {
	// 使用bcrypt库生成密码的哈希值
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 返回哈希值
	return string(hashedPassword)
}

// 函数用于比较哈希密码和密码是否匹配
// @params hashedPassword 保存密码
// @params password 待校验密码
func CompareHashAndPassword(hashedPassword, password string) bool {
	// 将哈希密码和密码转换为字节数组
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
