package main

import (
  "fmt"
  "example/hello/utilities/pinUtils"
)


func main() {
 
	data, dataError, dataResponseCode := pinUtils.PinFromUrl("https://ar.pinterest.com/pin/1046383294679022220/")
        fmt.Println(data, dataError, dataResponseCode)
}
