松鼠 AI 一面算法

请完成下面算法题：写个方法，随机获取一个 name

已知：
name1 的权重是 10；name2 的权重是 20；name3 的权重是 50。
需要按照权重比例随机取出一个 name（name1 出现的概率是 10/80，name2 出现的概率是 20/80，name3 出现的概率是 50/80）

```
func main() {
	randName1 := getRandName(map[string]int{"name1": 10, "name2": 20, "name3": 50})
	randName2 := getRandName(map[string]int{"name1": 10, "name2": 20, "name3": 50, "name4": 20})
	fmt.Printf(randName1, randName2)
}

func getRandName(arr map[string]int) string {

}

```
