# bcm4313网卡中内核模块被错误屏蔽

bcm4313网卡中内核模块被错误屏蔽

# 如何检查?

1, 通过pci检查是否有bcm4313网卡

2, 检查brcmsmac模块是否在使用

```
check () 
{ 
    local pci_info="$(lspci -d "14e4:4727" 2>/dev/null)";
    if [ "${pci_info}" ]; then
        if ! lsmod | grep brcmsmac; then
            return 1;
        fi;
    fi
}
```

# 如何修复?

1, 屏蔽掉b43和wl模块

2, 删除软件包bcmwl-kernel-source

```
fix () 
{ 
    echo "fix bcm4313 --------------------------------";
    echo "blacklist b43" | sudo tee -a /etc/modprobe.d/bcm.conf;
    echo "blacklist wl" | sudo tee -a /etc/modprobe.d/bcm.conf;
    echo "brcmsmac" | sudo tee -a /etc/modules;
    sudo apt-get remove -y bcmwl-kernel-source;
    exit 0
}
```
