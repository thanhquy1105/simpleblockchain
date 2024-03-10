package main

import (
	"fmt"

	"github.com/thanhquy1105/simpleblockchain/utils"
)

func main() {
	fmt.Println(utils.FindNeightbors("127.0.0.1", 5000, 0, 3, 5000, 5003))
}
