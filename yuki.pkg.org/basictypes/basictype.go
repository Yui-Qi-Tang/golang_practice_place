package basictype // package的名稱，跟檔案明不同沒關係，但建議該功能的主要程式碼中的package name，盡量跟檔案名稱相同，然後檔案名稱與資料夾名稱相同。

var integer int = 5

var (
	MaxInt uint64 = 1 << 64 -1
	Default string = "Hi, this is default string!!"
	BooleanVar bool = true
)

type person struct {
	name string
	ages int
}

func (p person) GetPersonName() string {
	return p.name
}

func (p person) GetPersonAges() int {
	return p.ages
}

func (p *person) SetPersonName(newName string) {
	/* 
		想要修改struct成員的值，建議是用指標；否則go會傳入一個struct copy，
		那麼，響修改的數值，會更新在struct copy上，而非該struct instance。

		雖然可以在這邊將傳入的struct copy的指標取出，然後再寫入新的數值後回傳(回傳指標位置):	
		p.name = newName
		return &p

		這樣一來，萬一結構的成員過大，會導致copy的量也會遽增，使得程式會吃掉不少記憶體。

		golang的FAQ也有建議，若是struct的方法想修改其成員，請傳遞一個指向該結構的指標進來。
		https://golang.org/doc/faq#methods_on_values_or_pointers

	*/
	p.name = newName
}

func NewPerson(name string, ages int) *person {
	newPerson := new(person)
	newPerson.name = name
	newPerson.ages = ages
	return newPerson
}

