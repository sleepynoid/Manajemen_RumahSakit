package pasien

import (
	"Manajemen_RumahSakit/entities"
	"fmt"
	// ui "github.com/gizak/termui/v3"
	// "github.com/gizak/termui/v3/widgets"
)

type NodePasien struct {
	Head *entities.Pasien
	Tail *entities.Pasien
}

func (pl *NodePasien) SearchPasien(target int) *entities.Pasien {
	fmt.Println("search")
	temp := pl.Head
	for temp != nil {
		fmt.Println("loop")
		if temp.ID == target {
			return temp
		}
		temp = temp.Next
	}
	return nil
}

func (pl *NodePasien) InsertPasien(pasien *entities.Pasien) {
	fmt.Println("InsertPasien start!!!")
	if pl.Head == nil {
		fmt.Println("first")
		pl.Head = pasien
		pl.Tail = pasien
	} else {
		fmt.Println("NOTfirst")
		pl.Tail.Next = pasien
		pl.Tail = pasien
		fmt.Println(pl.Head.ID, pl.Head.Next.ID)
	}
}

func (pl *NodePasien) DeletePasien(target int) bool {
	fmt.Println("DeletePasien")
	if pl.Head == nil {
		fmt.Println("1")
		return false
	}
	if pl.Head.ID == target {
		fmt.Println("2")
		fmt.Println(pl.Head.ID)
		pl.Head = pl.Head.Next
		fmt.Println(pl.Head.ID)
		return true
	}
	fmt.Println("3")
	temp := pl.Head
	for temp.Next != nil {
		fmt.Println("loop")
		if temp.Next.ID == target {
			fmt.Println("found")
			temp.Next = temp.Next.Next
			return true
		}
		temp = temp.Next
	}
	return false
}

func (pl *NodePasien) UpdatePasien(target int, update *entities.Pasien) bool {
	current := pl.Head
	for current != nil {
		if current.ID == target {
			current.Nama = update.Nama
			current.Tlp = update.Tlp
			current.Kondisi = update.Kondisi
			current.Riwayat = update.Riwayat
			return true
		}
	}
	fmt.Println("Not Found")
	return false
}

func (pl *NodePasien) ListPasien() {
	temp := pl.Head
	for temp != nil {
		print(temp.ID, "->")
		temp = temp.Next
	}
}
