#!/bin/bash


echo "开始推送代码到gitee"
git push gitee main
echo "开始推送代码到github"
git push github main
echo "代码推送完成"