package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kyokomi/lottery"
)

type Ball struct {
	Color string
	Count int
}

func (b *Ball) Prob() int {
	return b.Count
}

func (b *Ball) String() string {
	return fmt.Sprintf("%s(%d)", b.Color, b.Count)
}

// 袋の中に球が3種類入っていてそれを順次取り出して行くサンプル
func main() {
	lot := lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))
	red := &Ball{"Red", 1}
	green := &Ball{"Green", 10}
	blue := &Ball{"Blue", 3}

	bag := []lottery.Interface{red, green, blue}

	for {
		i := lot.Lots(bag...)
		if i < 0 {
			break
		}
		b, _ := bag[i].(*Ball)
		b.Count = b.Count - 1
		fmt.Println(b.Color, "--", red, green, blue)
	}
}
