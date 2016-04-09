# 意外的终止操作导致异常状态

意外的终止操作导致异常状态

apt以及dpkg都有内部状态如果强制或意外终止正在
进行的任务很有可能导致包管理系统进入不一致状态．

# 如何检查?

1. 快速通过/var/lib/dpkg/updates目录情况检测
   是否存在dpkg异常终止．
2. 通过apt-get check检测是否存在apt中断问题．



# 如何修复?

1. 使用dpkg --configre -a修复dpkg问题
2. 使用apt-get install -f修复apt问题


