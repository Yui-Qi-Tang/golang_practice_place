package tools

// import "fmt"

var arraySize int // private

func init() {
	arraySize = 10
}

func ShowArrayInfo() {
	// fmt.Println(arraySize)
	// array := makeArray()
	// fmt.Println(array)
}

func getArraySize() int {
	return arraySize
}

func makeArray() []int {
	return make([]int, arraySize, arraySize)
}

func SetArraySize(newArraySize int) {
	arraySize = newArraySize
}

func RemoveDupElement(targetArray []int) []int {
	countsArray := makeArray()
	for i:=0; i<len(targetArray); i++ {
		countsArray[targetArray[i]]++
		tmp := countsArray[targetArray[i]]
		if tmp >= 2{
			targetArray[i] = -1
		}
	}
	return targetArray
}