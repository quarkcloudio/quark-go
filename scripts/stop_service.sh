#!/bin/sh
PROJECT_PATH=$(cd $(dirname $0)|cd ..|pwd)
PROJECT_NAME=/main

# stop process
tpid=`ps -ef|grep $PROJECT_PATH$PROJECT_NAME|grep -v grep|grep -v kill|awk '{print $2}'`
if [[ ${tpid} ]]; then
    echo 'Stop Process...'
    # 是先关闭和其有关的程序,再将其关闭
    kill -15 $tpid
fi

sleep 5

tpid=`ps -ef|grep $PROJECT_PATH$PROJECT_NAME|grep -v grep|grep -v kill|awk '{print $2}'`
if [[ ${tpid} ]]; then
    echo 'Kill Process!'
    kill -9 $tpid
else
    echo 'Stop Success!'
fi