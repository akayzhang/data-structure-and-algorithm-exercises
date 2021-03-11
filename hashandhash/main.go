package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func GenSignature(unixSec int64, appid uint32, echo string, rand string, token string) string {
	//2.签名校验
	basic := strconv.FormatUint(uint64(unixSec), 10) + "|" +
		echo + "|" +
		strconv.FormatUint(uint64(appid), 10) + "|" +
		token
	tool := md5.New()
	tool.Write([]byte(basic))
	firstHash := tool.Sum(nil)  //先将基准数据md5

	fmt.Println("basic is ", basic)
	fmt.Println("[]byte(basic):", []byte(basic));

	fmt.Println("firstHash is ", firstHash)
	fmt.Println("firstHash is string:", hex.EncodeToString(firstHash[0:]))

	tool.Reset()
	tool.Write(firstHash) //----
	tool.Write([]byte(rand))
	saltMD5 := tool.Sum(nil)  //再将rand加盐md5, 还是md5再加盐更加安全, 而且md5算法相比sha256, 对cpu的消耗更低(30%)
	fmt.Println("saltMD5:===", saltMD5);
	return hex.EncodeToString(saltMD5[0:])
}

func main() {
	fmt.Println("hello， world")
	signa := GenSignature(1615365252 , 20032, "2345245", "rand", "zOOlS80Xz0BHng7l")
	fmt.Println("hello， signa:", signa)

}
