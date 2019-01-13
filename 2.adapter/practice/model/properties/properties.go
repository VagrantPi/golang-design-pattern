package properties

import (
	"os"
	"time"

	"github.com/magiconair/properties"
)

// StoreToFile - properties 缺少 store method ，所以這邊我們實作一個
func StoreToFile(OutputStream *properties.Properties, header string) {
	f, err := os.Create("./newfile.txt")
	check(err)
	defer f.Close()

	f.WriteString("# " + header + "\n")
	f.Sync()
	f.WriteString("# " + time.Now().Format(time.RFC1123) + "\n")
	f.Sync()

	keys := OutputStream.Keys()
	for i := 0; i < len(keys); i++ {
		v, ok := OutputStream.Get(keys[i])
		if !ok {
			continue
		}
		f.WriteString(keys[i] + "=" + v + "\n")
		f.Sync()
	}

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
