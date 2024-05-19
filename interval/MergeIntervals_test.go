package interval


func TestMergeIntervals(){
	intervals := [][]int{
		{1,3},
		{2,6},
		{8,10},
		{15,18},
	}

	res :=MergeInterval(intervals)

	println(res)
}