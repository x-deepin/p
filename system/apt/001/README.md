# No ack for apt-get update

No ack for apt-get update

因为系统更新导致的apt-get卡主导致后续内容全部
无法进行．

# How to check?

Get the result by checking file existing in /var/log/lastore.

```
check () 
{ 
    start=$(ps -o start,cmd -C apt-get | grep /var/lib/lastore/source.d | awk '{print $1}');
    echo $start
}
```

# How to fix?

just remove the old files

```
fix () 
{ 
    pid=$(ps -o pid,cmd -C apt-get | grep /var/lib/lastore/source.d | awk '{print $1}');
    if $pid; then
        kill $pid;
    fi
}
```
