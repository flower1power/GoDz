package main

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	list []Bin
}

func newBin(
	id string,
	private bool,
	createdAt time.Time,
	name string) *Bin {

	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
}

func main() {

}
