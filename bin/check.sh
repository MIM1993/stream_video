#!/bin/bash


pid=`netstat -tlnup | grep 8889 |awk '{printf $7}'|cut -d/ -f1`


echo $pid
kill -9 $pid