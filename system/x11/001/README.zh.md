# Xauthority 权限错误导致无法登陆系统

Xauthority 权限错误导致无法登陆系统

1. ~/.Xauthority 为每个用户的当前文件，部分程序使用sudo运行
   时可能导致该文件权限错误, 导致后续无法登陆

2. /var/lib/lightdm/.Xauthority 为display manager创建，若错误
  的安装模式可能导致该文件权限错误，导致后续无法登陆系统

[Xauthority](https://en.wikipedia.org/wiki/X_Window_authorization)

# 如何检查?

遍历所有Xauthority文件确保都能匹配到正确的用户以及权限

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

# 如何修复?

将错误的文件权限以及文件所属调整正确

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
