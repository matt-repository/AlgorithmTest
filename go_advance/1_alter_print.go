package go_advance

import (
	"fmt"
	"strings"
)

//1. 交替打印数字和字⺟
//问题描述
//使⽤两个 goroutine 交替打印序列，⼀个 goroutine 打印数字， 另外⼀个 goroutine 打印字⺟， 最终效果
//如下：
//源码参考
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

//解题思路
//问题很简单，使⽤ channel 来控制打印的进度。使⽤两个 channel ，来分别控制数字和字⺟的打印序列， 数字打
//印完成后通过 channel 通知字⺟打印, 字⺟打印完成后通知数字打印，然后周⽽复始的⼯作。

func AlterPrint() string {
	retStr := strings.Builder{}
	numChan, letterChan := make(chan bool), make(chan bool)
	go func() {
		i := 0
		for {
			<-numChan
			i++
			retStr.WriteString(fmt.Sprintf("%d", i))
			i++
			retStr.WriteString(fmt.Sprintf("%d", i))
			if i == 28 {
				close(numChan)
				break
			}
			letterChan <- true
		}
	}()

	numChan <- true
	letter := 'A' - 1
	for {
		<-letterChan
		letter++
		retStr.WriteString(string(letter))
		letter++
		retStr.WriteString(string(letter))
		numChan <- true
		if letter == 'Z' {
			close(letterChan)
			break
		}
	}
	fmt.Print(retStr.String())
	return retStr.String()
}
