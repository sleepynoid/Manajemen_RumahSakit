package controllers

import (
	"Manajemen_RumahSakit/entities"
	"Manajemen_RumahSakit/model/pasien"
	"fmt"
)

type PasienController struct {
	PasienList *pasien.NodePasien
}

func (pc *PasienController) searchPasien(ID int) *entities.Pasien {
	return pc.PasienList.SearchPasien(ID)
}

func (pc *PasienController) tambahPasien(input entities.Pasien) {
	temp := &entities.Pasien{
		ID:      input.ID,
		Nama:    input.Nama,
		Tlp:     input.Tlp,
		Kondisi: input.Kondisi,
		Riwayat: input.Riwayat,
		Next:    nil,
	}
	pc.PasienList.InsertPasien(temp)
}

func (pc *PasienController) DeletePasien(ID int) bool {
	return pc.PasienList.DeletePasien(ID)
}

func (pc *PasienController) UpdatePasien(ID int, input entities.Pasien) bool {
	update := &entities.Pasien{
		ID:      input.ID,
		Nama:    input.Nama,
		Tlp:     input.Tlp,
		Kondisi: input.Kondisi,
		Riwayat: input.Riwayat,
		Next:    nil,
	}
	return pc.PasienList.UpdatePasien(ID, update)
}

func (pc *PasienController) ListPasien() {
	current := pc.PasienList.Head
	for current != nil {
		fmt.Println("ID Pasien:", current.ID)
		fmt.Println("Nama Pasien:", current.Nama)
		fmt.Println("Nomor Telepon Pasien:", current.Tlp)
		fmt.Println("Kondisi Pasien:", current.Kondisi)
		fmt.Println("Riwayat Pasien:", current.Riwayat)
		fmt.Println()
		current = current.Next
	}
}
