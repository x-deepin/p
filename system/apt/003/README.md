# Broken package system by unexpected interrupt operation

Broken package system by unexpected interrupt operation

apt and dpkg have their internal status. It can't work
properly if the running operation was unexpected interrupt,
e.g., force reboot.

# How to check?

1. quickly check the status of /var/lib/dpkg/lock.
2. use "apt-get check" to see whether the apt broken.

```
check () 
{ 
    pgrep dpkg > /dev/null 2>&1;
    if [[ $? == 0 && -e /var/lib/dpkg/updates ]]; then
        exit 1;
    fi;
    exit 0
}
```

# How to fix?

1. dpkg --configure -a to fix dpkg broken
2. apt-get install -f to fix apt broken

```
fix () 
{ 
    echo "Fixing dpkg...";
    dpkg --force-confold --configure -a;
    echo "Done";
    echo "Fixing apt...";
    apt-get install -f;
    echo "Done"
}
```
