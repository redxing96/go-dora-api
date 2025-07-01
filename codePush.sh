#!/bin/bash
###
 # @Description: 推送代码到gitee
 # @Author: redxing96@163.com
 # @Date: 2025-06-26 19:12:22
 # @LastEditTime: 2025-07-01 20:08:55
 # @LastEditors: front end cabbage
 # @FilePath: /go-dora-api/codePush.sh
### 


echo "开始推送代码到gitee"
git push gitee develop
echo "开始推送代码到github"
git push github develop
echo "代码推送完成"

