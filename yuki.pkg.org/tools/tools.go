package tools

/*
	Remove duplicated element from array in O(n)
	Use additional memory to store counts of element of array
	if the counts > 2 -> the element is duplicated then set the array = -1

	Additional memory is array, denoted tmpA;
	The index of tmpA is denoted the target array's value and value of tmpA is counts number.
	So, assumption
		Target array = [1, 2, 3, 3, 6, 6, 19], and
		tmpA = [0, 0, 0, 0, 0,......0(19th)] in initial phase.
		(The length of tmpA is '19' because I need to use index of tmpA to store 19 counts of Target array in this example)

		Let Counts(elementVal) as funtion that used elementVal to map tmpA index, and value is tmpA[elementVal]++,
		where elementVal is from Target array.
	so the algorithm:

	for x in Target array {
		valueCounts = Counts(x);
		if valueCounts > 2 { // duplicated
            RemoveFromTargetArray(x); 
		}
	}

	where RemoveFromTargetArray(index) is a function to remove element of Target array by index,
	for now, the I set Target array[index] = -1 if valueCounts > 2 instead of this function.
*/
import (
	"yuki.pkg.org/tools/logger"
)

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
	// get max length for countsArray
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
	logger.InfoLog("remove duplicated element done!", "in tools/tools.go")
	return targetArray
}