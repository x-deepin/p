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



