# config
一个读取配置文件的工具，目前暂时只支持.conf类型的配置文件。

基本使用示例
```go
func main() {
	var CONF = config.CONF

	if err := CONF.LoadConfiguration(yourFilePath); err != nil {
		panic(err)
	}

	...

	// get values.
	str := CONF.GetString(group, key)
	boolean := CONF.GetBoo(group, key)
	integer := CONF.GetInt(group, keyy)
}
```

设置默认值示例，默认值的优先级比从配置文件里面读取的配置信息低。
```go
func init() {
	RegisterServer()
}

// set default config values.
func RegisterServer() {
	group :=  config.NewGroup("server")

	group.SetString("host", "127.0.0.1")
	group.SetString("port", "8080")
	config.CONF.Register(group)
}
```
