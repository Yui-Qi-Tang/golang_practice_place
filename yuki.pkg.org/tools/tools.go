package tools

var arraySize int // private

func init() {
	arraySize = 10
}

func getArraySize() int {
	return arraySize
}

func makeArray() []int {
	return make([]int, arraySize, arraySize)
}

func SetArraySize(targetArray []int) {
	max := len(targetArray)
    for i:=0; i<len(targetArray); i++{
        if targetArray[i] > max {
            max = targetArray[i]
		}
	}
	arraySize = max + 1 // To prevent index out of range
}

func RemoveDupElement(targetArray []int) []int {
	countsArray := makeArray()
	for i:=0; i<len(targetArray); i++ {
		countsArray[targetArray[i]]++
		tmp := countsArray[targetArray[i]]
		if tmp >= 2 {
			targetArray[i] = -1
		}
	}
	return targetArray
}