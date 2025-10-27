# openlist-plugin-exa
测试 OpenList 插件系统的示例插件
## 安装
1. 需要clone github.com/dezhishen/OpenList 到 openlist-dezhishen 目录下
2. 切换分支到 plugin-rpc
3. clone 本仓库到 openlist-plugin-exa 目录下
4. 进入 openlist-plugin-exa 目录，执行 go mod tidy
5. 执行 `mkdir -p ./plugins && go build -o ./plugins/openlist-plugin-exa.exe examples/main.go`
6. 运行 本仓库 `go run main.go` 文件，测试插件是否加载成功
## 说明
本插件示例演示了如何在 OpenList 中创建一个插件，并注册多个驱动器（drivers）。插件系统允许用户通过插件扩展 OpenList 的功能，而无需修改核心代码。