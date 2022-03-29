# config
一个读取配置文件的工具

```go
package main

func main() {
  err := config.LoadConfiguration(filePath)
  if err != nil {
    panic(err)
  }
  
  ...
  
  <!--  use CONF  -->
  _ := CONF.GetString(group, key)
}
```
