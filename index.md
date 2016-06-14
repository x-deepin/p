- [system.001](https://github.com/x-deepin/p/blob/master//system/001)

    System files maybe hacked or damaged
    
    It may be caused some weird things happen.
    
    
- [system.apt.001](https://github.com/x-deepin/p/blob/master//system/apt/001)

    No ack for apt-get update
    
    因为系统更新导致的apt-get卡主导致后续内容全部
    无法进行．
    
    
- [system.apt.002](https://github.com/x-deepin/p/blob/master//system/apt/002)

    Wrong package version in local system
    
    Local system may had installed some packages
    which has different version with http://packages.deepin.com/deepin.
    It may be caused local system can't correctly upgrade.
    
    
- [system.apt.003](https://github.com/x-deepin/p/blob/master//system/apt/003)

    Broken package system by unexpected interrupt operation
    
    apt and dpkg have their internal status. It can't work
    properly if the running operation was unexpected interrupt,
    e.g., force reboot.
    
    
