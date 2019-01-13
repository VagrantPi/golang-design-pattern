"github.com/magiconair/properties" 可以用來管理 key-value 集合

```
year=2019
month=1
day=13
```

請使用 Adapter Pattern 將 input.txt 中的 key-value 集合保存至文件中的 FileProperties class，FileProperties 為 FileIO(Target 角色) 的 implementation，然後經過一些修改後存成 newfile.txt

FileIO:

```go
type FileIO interface {
	ReadFile(fileName string)
	WriteToFile(fileName string)
	SetVale(key, value string)
	GetValue(key string)
}
```

main:

```go
func main() {
	var fp model.FileIO = &model.FileProperties{}
	fp.ReadFile("./input.txt")
	fp.SetVale("year", "2019")
	fp.SetVale("month", "1")
	fp.SetVale("day", "13")
	fp.WriteToFile("newfile.txt")
}
```

input.txt

```
year=1999
```

newfile.txt

```
# written by FileProperties
# Sun, 13 Jan 2019 17:03:09 CST1999
day=13
year=2019
month=1
```
