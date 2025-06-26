#!/bin/bash
###
 # @Description: 推送代码到gitee
 # @Author: redxing96@163.com
 # @Date: 2025-06-26 19:12:22
 # @LastEditTime: 2025-06-26 21:12:35
 # @LastEditors: front end cabbage
 # @FilePath: /go-dora-api/codePush.sh
### 

# 定义颜色
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # 无色

echo -e "${YELLOW}开始推送代码到gitee${NC}"
git push gitee develop
# echo -e "${YELLOW}开始推送代码到github${NC}"
# git push github main
echo -e "${GREEN}代码推送完成${NC}"

