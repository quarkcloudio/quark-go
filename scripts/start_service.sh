#!/bin/sh
PROJECT_PATH=./
PROJECT_NAME=main
PROJECT_LOG_NAME=storage/logs/app.log

# 判断logs文件夹下是否有日志文件，没有则创建，用于记录日志
if [ ! -f $PROJECT_PATH$PROJECT_LOG_NAME ]; then
  touch $PROJECT_PATH$PROJECT_LOG_NAME
fi

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

# start process
tpid=`ps -ef|grep $PROJECT_PATH$PROJECT_NAME|grep -v grep|grep -v kill|awk '{print $2}'`
if [[ ${tpid} ]]; then
    echo 'App is already running.'
else
    echo 'App is NOT running.'

    nohup $PROJECT_PATH$PROJECT_NAME >$PROJECT_PATH$PROJECT_LOG_NAME 2>&1 &
    
    echo Start Success!
    
    if [ $1 ]; then
        sleep 2
        tail -f $PROJECT_PATH$PROJECT_LOG_NAME
    fi
fi
