# Wrong permission of Xauthority cause failed login DE

Wrong permission of Xauthority cause failed login DE

1. ~/.Xauthority owner by every user，There have some programs run
   with sudo may cause this file to be modified unexpected.

2. The file of /var/lib/lightdm/.Xauthority owner by display manager.
  It will cause login failed if the permission of the file is wrong.

[Xauthority](https://en.wikipedia.org/wiki/X_Window_authorization)

# How to check?

Ensure the owner of the file is matching the user and the
permission of the file is 0600

```
check () 
{ 
    for f in $(ls /home/*/.Xauthority /var/lib/lightdm/.Xauthority);
    do
        echo "Checking $f";
        if [[ $(stat -c '%a' $f) != "600" ]]; then
            echo "Wrong permission for $f";
            exit 1;
        fi;
        local dir_owner_name=$(stat -c '%U' $(dirname $f));
        if [[ $(stat -c '%U' $f) != $dir_owner_name ]]; then
            echo "Wrong owner for $f";
            exit 2;
        fi;
    done;
    exit 0
}
```

# How to fix?

Set the right permission and owner for the problem files

```
fix () 
{ 
    for f in $(ls /home/*/.Xauthority /var/lib/lightdm/.Xauthority);
    do
        if [[ $(stat -c '%a' $f) != "600" ]]; then
            echo "Fixing permission for $f";
            chmod 0600 $f && exit 1;
        fi;
        local dir_owner_name=$(stat -c '%U' $(dirname $f));
        if [[ $(stat -c '%U' $f) != $dir_owner_name ]]; then
            echo "Fixing permission for $f";
            chown $dir_owner_name $f && exit 2;
        fi;
    done;
    exit 0
}
```
