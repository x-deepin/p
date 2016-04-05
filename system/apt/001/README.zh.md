# apt-get update长时间无响应
apt-get update长时间无响应

因为系统更新导致的apt-get卡主导致后续内容全部
无法进行．

# 如何检查?
通过检测/var/log/lastore/下面是否包含旧有的日志文件目录判断结果

```
check () 
{ 
    start=$(ps -o start,cmd -C apt-get | grep /var/lib/lastore/source.d | awk '{print $1}');
    echo $start
}
```

# 如何修复?
删掉旧有文件即可

```
fix () 
{ 
    pid=$(ps -o pid,cmd -C apt-get | grep /var/lib/lastore/source.d | awk '{print $1}');
    if $pid; then
        kill $pid;
    fi
}
```

