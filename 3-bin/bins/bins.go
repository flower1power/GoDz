package bins

import (
	"github.com/google/uuid"
	"time"
)

type Bin struct {
	Id        uuid.UUID `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Lists []Bin `json:"lists"`
}

func NewBinList() *BinList {
	binList := BinList{}

	return &binList
}

func NewBin(
	private bool,
	name string) *Bin {

	return &Bin{
		Id:        uuid.New(),
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func AddBinList(binList *BinList, newBin *Bin) {
	binList.Lists = append(binList.Lists, *newBin)
}

func DeleteBin(list *BinList, delBin *Bin) {
}
