package bins

import "time"

type Bin struct {
	Id        string `json:"id"`
	Private   bool `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string `json:"name"`
}

type BinList struct {
	Lists []Bin `json:"lists"`
}

func NewBinList() *BinList {
	binList := BinList{}

	return &binList
}

func NewBin(
	id string,
	private bool,
	createdAt time.Time,
	name string) *Bin {

	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}

func (bin *Bin) save(){
	
}


func AddBinList(binList *BinList, newBin *Bin) {
	binList.Lists = append(binList.Lists, *newBin)
}

func DeleteBin(list *BinList, delBin *Bin){
}