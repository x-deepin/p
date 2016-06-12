# Wrong package version in local system

Wrong package version in local system

Local system may had installed some packages
which has different version with http://packages.deepin.com/deepin.
It may be caused local system can't correctly upgrade.

# How to check?

Using apt-show-versions to detect packages which has
different version in local between http://packages.deepin.com/deepin

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

# How to fix?

1. reinstall packages with corrected version
2. remove the unknown packages

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
