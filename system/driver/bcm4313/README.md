# Check wrong kernel module blacklist for bcm4313

Check wrong kernel module blacklist for bcm4313

# How to check?

1, check pci of bcm4313

2, check whether brcmsmac is in use

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

# How to fix?

1, blacklist kernel module b43 and wl
2, remove package bcmwl-kernel-source

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
