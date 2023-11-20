# 搜索平台

## 项目结构

```shell
search_engine/
├── analyzer            // 解析器，分词作用，与索引平台的解析不一样
├── cmd                 // 启动文件
├── data                // 存放boltdb(但目前存放在根目录下) 🫓 again
│   └── db
├── ranking             // 排序器
├── repository          // 存储介质
│   ├── db              // OLTP:mysql
│   │   └── dao
│   ├── spark           // replace mapreduce in the future, 🫓 again
│   ├── starrock        // OLAP:starrocks
│   │   ├── bi_dao
│   │   └── initdb.d
│   └── storage         // 存储db的函数
├── rpc                 // rpc调用
├── service             // 服务层
└── test                // 单词文件
```
