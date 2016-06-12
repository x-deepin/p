# 本地安装包与仓库不一致

本地安装包与仓库不一致

用户本地安装的部分包可能与仓库不一致或安装了仓库中不存
的包．　由此可能导致系统无法正常更新到合适的版本．

# 如何检查?

通过使用apt-show-versions工具检测出本地包版本与仓库不一致的情况．

```
check () 
{ 
    helper_ensure_installed_apt_show_versions;
    IFS='
' GLOBIGNORE='*';
    declare -i newer=0;
    declare -i other=0;
    for line in $(apt-show-versions);
    do
        if [[ ! -n $(echo $line | grep -E 'newer than version|No available') ]]; then
            continue;
        fi;
        if [[ -n $(echo $line | grep -F 'newer than version') ]]; then
            newer+=1;
            echo $line;
        fi;
        if [[ -n $(echo $line | grep -F 'No available') ]]; then
            other+=1;
            echo $line;
        fi;
    done;
    [[ $other == 0 ]] || printf "\nYou have %d packages doesn't in repository\n" $other;
    [[ $newer == 0 ]] || printf "\nYou have %d packages newer than version in repository\n" $newer;
    if [[ $other > 0 || $newer > 0 ]]; then
        exit 1;
    fi
}
```

# 如何修复?

1. 重新安装仓库中对应的版本．
2. 删除仓库中不存在的包

```
fix () 
{ 
    helper_ensure_installed_apt_show_versions;
    local unknown_packages=();
    IFS='
' GLOBIGNORE='*';
    for line in $(apt-show-versions);
    do
        if [[ ! -n $(echo $line | grep -E 'newer than version|No available') ]]; then
            continue;
        fi;
        if [[ -n $(echo $line | grep -F 'newer than version') ]]; then
            helper_fix_wrong_version_package $(echo $line | cut -d ':' -f1);
        fi;
        if [[ -n $(echo $line | grep -F 'No available') ]]; then
            unknown_packages+=("$(echo $line | cut -d ':' -f1)");
        fi;
    done;
    if [[ ${#unknown_packages[@]} > 0 ]]; then
        echo "There has some unknown packages wouldn't remove.";
        echo "You can execute below command to remove this packages";
        echo -e "\t sudo apt-get remove ${unknown_packages[@]}";
    fi;
    exit 0
}
```
