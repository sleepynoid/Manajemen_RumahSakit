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

func (ll *NodePasien) SearchPasien(target int) *entities.Pasien {
	fmt.Println("search")
	temp := ll.Head
	for temp != nil {
		fmt.Println("loop")
		if temp.ID == target {
			return temp
		}
		temp = temp.Next
	}
	return nil
}

func (ll *NodePasien) InsertPasien(ID int) {
	fmt.Println("InsertPasien start!!!")
	temp := new(entities.Pasien)
	temp.ID = ID
	temp.Next = nil
	// fmt.Println(temp)
	if ll.Head == nil {
		fmt.Println("first")
		ll.Head = temp
		ll.Tail = temp
	} else {
		fmt.Println("NOTfirst")
		ll.Tail.Next = temp
		ll.Tail = temp
		fmt.Println(ll.Head.ID, ll.Head.Next.ID)
	}
}

func (ll *NodePasien) DeletePasien(target int) {
	fmt.Println("DeletePasien")
	if ll.Head == nil {
		fmt.Println("1")
		return
	}
	if ll.Head.ID == target {
		fmt.Println("2")
		fmt.Println(ll.Head.ID)
		ll.Head = ll.Head.Next
		fmt.Println(ll.Head.ID)
		return
	}
	fmt.Println("3")
	temp := ll.Head
	for temp.Next != nil {
		fmt.Println("loop")
		if temp.Next.ID == target {
			fmt.Println("found")
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}
}

func (ll *NodePasien) EditPasien(target int, a int) {
	temp := ll.SearchPasien(target)
	if temp != nil {
		temp.ID = a
		return
	}
	fmt.Println("Not Found")
}

func (ll *NodePasien) ListPasien() {
	temp := ll.Head
	for temp != nil {
		print(temp.ID, "->")
		temp = temp.Next
	}
}
