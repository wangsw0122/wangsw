package main

import "fmt"

////类型重命名
//type Math = int
//type English = int

func main() {
	vg := &voteGame{students: []*student{
		&student{name:fmt.Sprintf("%s","小李")},
		&student{name:fmt.Sprintf("%s","小王")},
		&student{name:fmt.Sprintf("%s","小吴")},
		&student{name:fmt.Sprintf("%s","小张")},
		&student{name:fmt.Sprintf("%s","小康")},
	}}
	leader := vg.goRun()
	fmt.Println(leader)
}

type voteGame struct {
	students []*student
}

func (g *voteGame) goRun() *Leader {
	for _, item := range g.students {
		item.voteA(g.students[0])
	}
	maxScore := -1
	maxScoreIndex := -1
	for i, item := range g.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 {
		return (*Leader)(g.students[maxScoreIndex])
	}
	return nil
}

type Leader student

//投票选班长小游戏
type student struct {
	name string
	agree    int
	disagree int
}

func (std *student) voteA(target *student) {
	target.agree++
}
func (std *student) voteD(target *student) {
	target.disagree++
}
