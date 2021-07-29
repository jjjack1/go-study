//이건  차세대 언어인 go로 짠건 아님
package main

import "fmt"

type Bread struct {
	val string
}

type StrawberryJam struct {
	opened bool
}

type SpoonOfStrawberry struct {
}
type SpoonOfOrangeJam struct {
}

type Sandwitch struct {
	val string
}

type OrangeJam struct {
	opened bool
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

func OpenOrangeJam(jam *OrangeJam) {
	jam.opened = true
}

func GetOneSpoon(_ *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}
}

func GetOneOrangeJamSpoon(_ *OrangeJam) *SpoonOfOrangeJam {
	return &SpoonOfOrangeJam{}
}
func GetOneStrawberry(_ *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}
}

func PutJamOnBread(bread *Bread, Jam *SpoonOfStrawberry) {
	bread.val += " + Strawberry Jam"
}

func PutOrangeJamOnBread(bread *Bread, Jam *SpoonOfOrangeJam) {
	bread.val += " + Orange Jam"
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}
	return sandwitch
}

func main() {

	breads := GetBreads(2)
	jam := &OrangeJam{}
	OpenOrangeJam(jam)
	spoon := GetOneOrangeJamSpoon(jam)
	PutOrangeJamOnBread(breads[0], spoon)
	sandwitch := MakeSandwitch(breads)

	fmt.Println(sandwitch.val)

}
