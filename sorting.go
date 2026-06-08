package main 

func insertionSortAsc(data *[10]float64) {
	var idx int
	var temp float64
	for i := 1; i < 10; i++ {
		idx = i
		temp = data[i]
		for idx > 0 && data[idx-1] > temp {
			data[idx] = data[idx-1]
			idx = idx - 1
		}
		data[idx] = temp
	}
}

func selectionSort(data *dataBase, totalUser, id int) {
	//asc
	var idx int
	var temp dataComponent

	for i := 1; i < totalUser; i++ {
		idx = i
		temp = data[i]
		for y := i + 1; y < totalUser; y++ {
			if data[y].indexFloat(id) < data[idx].indexFloat(id) {
				idx = y
			}
		}
		data[i] = data[idx]
		data[idx] = temp
	}
}

func selectionSortDesc(data *dataBase, totalUser, id int) {
	//desc
	var idx int
	var temp dataComponent
	for i := 1; i < totalUser; i++ {
		idx = i
		temp = data[i]
		for y := i + 1; y < totalUser; y++ {
			if data[y].indexFloat(id) > data[idx].indexFloat(id) {
				idx = y
			}
		}
		data[i] = data[idx]
		data[idx] = temp
	}

}

func insertionSortData(data *dataBase, totalUser, id int) {
	var idx int
	var temp dataComponent

	for i := 2; i < totalUser; i++ {
		idx = i
		temp = data[i]
		for idx > 1 && data[idx-1].indexFloat(id) > temp.indexFloat(id) {
			data[idx] = data[idx-1]
			idx = idx - 1
		}
		data[idx] = temp
	}
}

func insertionSortDataDesc(data *dataBase, totalUser, id int) {
	var idx int
	var temp dataComponent
	for i := 2; i < totalUser; i++ {
		idx = i
		temp = data[i]
		for idx > 1 && data[idx-1].indexFloat(id) < temp.indexFloat(id) {
			data[idx] = data[idx-1]
			idx = idx - 1
		}
		data[idx] = temp
	}
}

func insertionSortDataString(data *dataBase, totalUser, id int) {
	var idx int
	var temp dataComponent
	for i := 2; i < totalUser; i++ {
		idx = i
		temp = data[i]
		for idx > 1 && data[idx-1].indexString(id) > temp.indexString(id) {
			data[idx] = data[idx-1]
			idx = idx - 1
		}
		data[idx] = temp
	}
}

func insertionSortDataStringDesc(data *dataBase, totalUser, id int) {
	var idx int
	var temp dataComponent
	for i := 2; i < totalUser; i++ {
		idx = i
		temp = data[i]
		for idx > 1 && data[idx-1].indexString(id) < temp.indexString(id) {
			data[idx] = data[idx-1]
			idx = idx - 1
		}
		data[idx] = temp
	}
}