package commd5


import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"log"
)



func ComMd5(s string)string{

	m := md5.New()
	_,err := io.WriteString(m,string(s))
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)		//已经输出，但是是编码
	// 将编码转换为字符串
	newArr := fmt.Sprintf("%x",arr)
	//输出字符串字母都是小写，转换为大写
	return strings.ToTitle(newArr)

}
