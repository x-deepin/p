- [system.001](https://github.com/x-deepin/p/blob/master//system/001)

    System files maybe hacked or damaged
    
    It may be caused some weird things happen.
    
    
- [system.apt.001](https://github.com/x-deepin/p/blob/master//system/apt/001)

    apt-get update长时间无响应
    
    因为系统更新导致的apt-get卡主导致后续内容全部
    无法进行．
    
    
- [system.apt.002](https://github.com/x-deepin/p/blob/master//system/apt/002)

    本地安装包与仓库不一致
    
    用户本地安装的部分包可能与仓库不一致或安装了仓库中不存
    的包．　由此可能导致系统无法正常更新到合适的版本．
    
    
- [system.apt.003](https://github.com/x-deepin/p/blob/master//system/apt/003)

    意外的终止操作导致异常状态
    
    apt以及dpkg都有内部状态如果强制或意外终止正在
    进行的任务很有可能导致包管理系统进入不一致状态．
    
    
- [system.x11.001](https://github.com/x-deepin/p/blob/master//system/x11/001)

    Xauthority 权限错误导致无法登陆系统
    
    1. ~/.Xauthority 为每个用户的当前文件，部分程序使用sudo运行
       时可能导致该文件权限错误, 导致后续无法登陆
    
    2. /var/lib/lightdm/.Xauthority 为display manager创建，若错误
      的安装模式可能导致该文件权限错误，导致后续无法登陆系统
    
    [Xauthority](https://en.wikipedia.org/wiki/X_Window_authorization)
    
    
