# 错误状态码

## 示例
```go
package main

import (
    "fmt"
    "github.com/yzha5/merr"
)

func demo1() error {
    return merr.New("recordNotFound","no rows in result set")
}

func demo2() error {
    return merr.New(1000,"error message")
}

func main() {
    err1:=demo1()
    if err1!=nil{
        fmt.Printf("error code: %v",err1.(*merr.M).Code) //recordNotFound
        fmt.Printf("error text: %s",err1.(*merr.M).Text) //no rows in result set
    }

    if err2:=demo2();err2!=nil{
        fmt.Printf("error code: %v",err2.(*merr.M).Code) //1000
        fmt.Printf("error text: %s",err2.(*merr.M).Text) //error message
    }
}
```