# shut-init

zhashut 的脚手架

## 使用指引

1. 进入到根目录下的 `config.yaml` 修改 `MySQL` 和 `redis` 的配置
2. 如果不需要 `redis`，可以在根目录的 `main.go` 里面将 `initialize.InitCache()` 给注释掉即可
3. `initialize.Router()` 函数里只做了初始化配置，根据自己情况添加路由
4. `docs/project.sql` 是用来放置项目用到的所以建表 SQL

> 暂时就这么多，后续待补充