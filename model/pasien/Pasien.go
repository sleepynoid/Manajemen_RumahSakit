package pasien

import (
	"Manajemen_RumahSakit/entities"
)

func search(target int, head *entities.Pasien) *entities.Pasien {
	temp := head
	var found *entities.Pasien = nil
	for temp != nil {
		if target == temp.ID {
			found = temp
		}
		temp = temp.Next
	}
	return found
}
func InsertPasien(ID int, head *entities.Pasien) {
	temp := new(entities.Pasien)
	temp.ID = ID
	temp.Next = nil
	if head == nil {
		head = temp
	} else {
		current := head
		for current.Next != nil {
			current = current.Next
		}
	}
}
func InsertDataPasien(ID int, head *entities.Pasien) {

}
