/*
 * @Description:
 * @Author: redxing96@163.com
 * @Date: 2025-06-26 19:21:13
 * @LastEditTime: 2025-06-26 19:21:22
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/model/crontab.go
 */
package model

import "time"

type CrontabEntity struct {
	Name        string
	Time        time.Time
	Status      int
	IsSingleton bool
}

type CrontabSearchOutput = CrontabEntity
