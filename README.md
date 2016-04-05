# 什么鬼
高级版How to Fix
1. 基于github的问题描述与跟踪完善
   - 点赞
   - 踩
   - 注释
   - 完善脚本

2. https://github.com/x-deepin/fixme 客户端自动检测当前系统已知的问题

3. 通过github api与各发型版的wiki/bug tracking system关联起来

# 如何编写修复脚本
1. 参考 https://github.com/x-deepin/p/blob/master/system/apt/001 下的fix脚本进行编写
  - 首先载入根目录下的functions文件
    ``` source ../../../functions ```
  - 然后使用COMMENT指令分别定义description, check, fix三者的描述
  - 分别定义实际操作的check, fix函数
2. 运行编写的脚本并生成支持语言的README.
```
./fix > README.md
LANG=zh_CN ./fix > README.zh.md
```

# 怎么玩?
```
├── driver
│   └── test
│       └── 001
│           ├── check
│           ├── fix
│           └── README.md
├── README.md
├── system
│   └── apt
│       └── 001
│           ├── detect
│           ├── fix
│           └── README.md
└── test
    └── 001
        ├── check
        ├── detect
        ├── fix
        └── README.md
```
有3个p
1. driver.test.001
2. system.apt.001
3. test.001

每个p又包含
1. 问题描述 (README.md)
2. 问题检测 (detect)
3. 问题修复 (fix)

fixme客户端拉取此git项目内容缓存来显示给用户

# 设想
1. 加个index文件，方便客户端之类的显示
2. 使用github issuse建立评论/评分系统．比如 driver.test.001 可以开一个issue进行讨论
3. 如果使用github接口，那么普通用户无法登陆．取舍:普通用户不需要反馈? github更容易聚集人气?
4. github作为纯数据存储地方．提供更友好的索引页面．　将issue以及README二次展示．与发型版的知识库集成起来．
5. 自己写个简单的server提供API避免依赖github

