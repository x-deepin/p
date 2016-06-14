# System files maybe hacked or damaged

System files maybe hacked or damaged

It may be caused some weird things happen.

# How to check?

Using /var/lib/dpkg/info/*.md5sums to do the integrity test

```
check () 
{ 
    return 0;
    cat /var/lib/dpkg/info/lastore-daemon.md5sums | awk '{ $2="/"+$2; print $2, $1, (md5sum $2 | getline)}'
}
```

# How to fix?

1. use dpkg -s to find the package contains the damaged files.
2. use apt-get install --reinstall to fix the damged files.

```
fix () 
{ 
    return 0
}
```
