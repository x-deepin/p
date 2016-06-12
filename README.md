# 问题集 (PSet)
PSet即Problem的集合. 每个Problme描述了一个问题的基本情况.
以及通过对应的fix脚本进行检测/修复.

PSet对应的使用者为两类

1. 直接在线游览生成出来的markdown文档. 作为wiki式的知识积累与完善站点.
   访问对应fix脚本目录的页面即可https://github.com/x-deepin/p/tree/master/system/apt/002

   - [中文索引](https://github.com/x-deepin/p/blob/master/index.zh.md)
   - [Index](https://github.com/x-deepin/p/blob/master/index.md)

2. 通过fixme客户端自动拉取PSet检测与执行(自动或手动)

```
$ ./bin/fixme show
+-----------------+---------------------------------------------------------+----------+
| ID              | Title                                                   | EffectMe |
+-----------------+---------------------------------------------------------+----------+
| driver.test.001 | rm /etc/X11/Xorg.conf                                   | ?        |
| system.apt.001  | No ack for apt-get update                               | no       |
| system.apt.002  | Wrong package version in local system                   | yes      |
| system.apt.003  | Broken package system by unexpected interrupt operation | no       |
| test.001        | I can't fix this.                                       | ?        |
+-----------------+---------------------------------------------------------+----------+
```

# 如何编写fix脚本?

fix脚本由Meta,Comment和Common Functions 3个基本元素组成,
并通过公用代码来实现外围函数与文档的自动生成.

fix通过相对于根目录的路径隐含了一个Problem ID(pid)字段.
即 system/apt/002/fix 这个P的id为system.apt.002.fix

在使用fixme客户端时需要使用这种格式的pid进行大部分操作.

具体可以参考 https://github.com/x-deepin/p/blob/master/system/apt/002/fix 脚本进行编写.

## Common Functions

fix脚本第一行有效代码必须引入本项目根目录下的functions文件一般固定形式

``` source ../../../functions ```

fix脚本最后一行有效代码必须调用```base_main $*```以便完成自动生成
基本功能如

```
% ./fix -h
Options:
    -f|--fix) dry run the fix function.
    -c|--check) dry run the check function.
    -d|--description) print the description of the problem.
    -t|--title) show the title of the problem.
    -l|--lang) specify display language.
    -m|--meta) print the meta data in json format.
    -v|--verbose) show code in render document.
    -h|--help) print this usage.
    --force) force run the function.

    Use fixme client instead use this script directly.

    The default action is print the document of
    the problem in markdown foramt.

    ./fix -v > README.md
    LANG=zh_CN ./fix -v > README.zh.md
```

## META

基本信息目前,使用时直接赋值即可(一定要先source进functions文件)

- META["AUTHOR"]="姓名 <邮箱地址>"

- META["AUTO_CHECK"]=false|false

  默认false. 若为true则自动进行定期检测.若检测有风险或较耗时请不要设置
  为true

- META["AUTO_FIX"]=true|false

  默认为false 若为true则当检测出有此p的时候自动执行修复代码

- META["Expiry_Date"]=`date`

  格式为date -d 可以识别的字符串. 当"本地系统时间"次日期则此脚本自动被
  fixme客户端屏蔽
- META["OVERDUE"]=true|false

  若为false则会被fixme客户端屏蔽. (但生成的文档以及脚本编号依旧保留)

## Documents

使用COMMENT函数编写文档.
```
COMMENT $lang $target <<EOF
文档正文.
这里的内容会保持原始缩进. 但其中的变量以及函数都会被实际执行.
请小心使用.
EOF
```
其中
 - $lang要与环境变量LANG的前两个字符一致.如zh,en,it等. (缺省为en)
 - $target字段目前有description, check, fix 3个有效值. (小写)

文档的定义建议放到对应的函数体前,即可支持markdown功能,又能直接
支持对函数本身的注释.


- Description

问题描述文档.
格式类似git的commit规范.即
第一行作为标题 首字母大写,不需要句尾符号,
第二留空.第三行之后作为详细描述.

- Check

建议紧挨着function check之前进行定义. 描述如何进行问题检测.

- Fix

建议紧挨着function fix之前进行定义. 描述如何进行问题修复

## Code

- Check问题检测代码.

结果通过$?返回. 即使用return语句.(若无则使用最后一条语句的执行结果)当
$? == 0 时候表示Check执行成功,系统*无*此问题. 其他值表示发现此问题.

- Fix

问题修复代码同Check方法. 结果通过$?返回,为0时候表示修复成功. (此
时执行Check应该也返回0)


# 测试脚本内容

使用自动生成的帮助文档即可大概了解如何进行测试.

```
./fix -h
```
- -f 尝试执行fix函数
- -c 尝试执行check函数
- --force 配合-f/-c参数执行真实操作.
- -l 测试不同语言的效果. 默认从$LANG进行解析.

# 生成文档

遵循以上几个简单的规则后实现了Comment以及check && fix两个函数即可
运行该脚本了. 默认会自动生成markdown格式的文档.

fixme客户端本身只依赖一个fix脚本文件. 但在线游览功能需要脚本编写者手动
生成文档. 使用以下命令即可.
```
./fix -v
```
根据支持的文档语言不同(请
至少支持zh和en)分别执行
```
LANG=en_US ./fix -v > README.md # 默认语言使用en
LANG=zh_CN ./fix -v > README.zh.md #其他语言在.md前加上对应的2字符代码
```

# 生成索引
在根目录下执行go run index.go 即可自动更新索引文件(index.md)

## 注意事项

1. Check以及Fix函数会被*自动*插入到文档中.但不包含函数内部调用的其他函
   数实现.
2. 一定要记得第一行加入```source functions```最后一行执行
   ```base_main "$*"```
3. fix脚本不可以从相对路径下引入其他脚本. 即fix脚本所属目录下的其他文
   件是不会被拷贝到执行目录里面的.
4. description文档的第一行会作为实际的标题进行显示(生成的markdown以及
   fixme客户端)
5. check以及fix函数是小写.


# 其他改进设想

1. 加个index文件，方便客户端之类的显示
2. 使用github issuse建立评论/评分系统．比如 driver.test.001 可以开一个
   issue进行讨论
3. 如果使用github接口，那么普通用户无法登陆．取舍:普通用户不需要反馈?
   github更容易聚集人气?
4. github作为纯数据存储地方．提供更友好的索引页面．　将issue以及README
   二次展示．与发型版的知识库集成起来．
5. 一个简单的server提供API避免依赖github API
