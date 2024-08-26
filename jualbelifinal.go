package main

import "fmt"

const NMAX = 10

type barang struct {
	nama    string
	harga   int
	stok    int
	terjual int
}

type transaksi struct {
	namaBarang string
	jumlah     int
	hargaTotal int
}

type arrBarang [NMAX]barang
type arrTransaksi [NMAX]transaksi

var jumlahBarang int = 0
var jumlahTransaksi int = 0

func main() {
	var brg arrBarang
	var trs arrTransaksi

	for {
		var choice int
		fmt.Println("===========APLIKASI JUAL BELI==============")
		fmt.Println("1. Tambah data barang")
		fmt.Println("2. Ubah data barang")
		fmt.Println("3. Hapus data barang")
		fmt.Println("4. Tambah data transaksi")
		fmt.Println("5. Ubah data transaksi")
		fmt.Println("6. Hapus data transaksi")
		fmt.Println("7. Urutkan barang berdasar harga")
		fmt.Println("8. Urutkan barang berdasar stok")
		fmt.Println("9. Cari barang berdasar nama")
		fmt.Println("10. Cari barang berdasar harga")
		fmt.Println("11. Cari barang berdasar stok")
		fmt.Println("12. Tampilkan data transaksi")
		fmt.Println("13. Tampilkan data barang")
		fmt.Println("14. Tampilkan 5 barang terlaris")
		fmt.Println("15. Keluar")

		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&choice)
		fmt.Println(" ")

		switch choice {
		case 1:
			tambahdatabarang(&brg)
			fmt.Println(" ")
		case 2:
			ubahdatabarang(&brg)
			fmt.Println(" ")
		case 3:
			hapusdatabarang(&brg)
			fmt.Println(" ")
		case 4:
			tambahdatatransaksi(&brg, &trs)
			fmt.Println(" ")
		case 5:
			ubahdatatransaksi(&trs, &brg)
			fmt.Println(" ")
		case 6:
			hapusdatatransaksi(&trs, &brg)
			fmt.Println(" ")
		case 7:
			selectionsortHarga(&brg)
			fmt.Println(" ")
		case 8:
			insertionSortStok(&brg)
			fmt.Println(" ")
		case 9:
			sequentialsearchNama(&brg)
			fmt.Println(" ")
		case 10:
			binarySearchHarga(&brg)
			fmt.Println(" ")
		case 11:
			binarySearchStok(&brg)
			fmt.Println(" ")
		case 12:
			showAllDataTransaksi(&trs)
			fmt.Println(" ")
		case 13:
			showAllData(&brg)
			fmt.Println(" ")
		case 14:
			showTop5BarangTerlaris(&brg)
			fmt.Println(" ")
		case 15:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func tambahdatabarang(B *arrBarang) {
	var namabarang string
	var hargabarang, stokbarang, jml int

	fmt.Print("Berapa jenis barang yang ingin ditambahkan? ")
	fmt.Scanln(&jml)

	for i := 0; i < jml && jumlahBarang < NMAX; i++ {
		fmt.Println("Masukkan nama barang:")
		fmt.Scanln(&namabarang)
		fmt.Println("Masukkan harga barang:")
		fmt.Scanln(&hargabarang)
		fmt.Println("Masukkan stok barang:")
		fmt.Scanln(&stokbarang)

		B[jumlahBarang].nama = namabarang
		B[jumlahBarang].harga = hargabarang
		B[jumlahBarang].stok = stokbarang
		jumlahBarang++
	}
}

func ubahdatabarang(B *arrBarang) {
	var choicenama, newnama string
	var newharga, newstok int

	fmt.Println("Masukkan nama barang yang ingin diubah:")
	fmt.Scanln(&choicenama)

	for i := 0; i < jumlahBarang; i++ {
		if B[i].nama == choicenama {
			fmt.Print("Masukkan nama baru: ")
			fmt.Scanln(&newnama)
			fmt.Print("Masukkan harga baru: ")
			fmt.Scanln(&newharga)
			fmt.Print("Masukkan stok baru: ")
			fmt.Scanln(&newstok)

			B[i].nama = newnama
			B[i].harga = newharga
			B[i].stok = newstok

			fmt.Println("Data barang berhasil diubah.")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

func hapusdatabarang(B *arrBarang) {
	var choicenama string

	fmt.Println("Masukkan nama barang yang ingin dihapus:")
	fmt.Scanln(&choicenama)

	for i := 0; i < jumlahBarang; i++ {
		if B[i].nama == choicenama {
			for j := i; j < jumlahBarang-1; j++ {
				B[j] = B[j+1]
			}
			jumlahBarang--
			fmt.Println("Data barang berhasil dihapus.")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

func tambahdatatransaksi(B *arrBarang, T *arrTransaksi) {
	var namaBarang string
	var jumlah int

	fmt.Println("Masukkan nama barang yang ingin dibeli:")
	fmt.Scanln(&namaBarang)

	for i := 0; i < jumlahBarang; i++ {
		if B[i].nama == namaBarang {
			fmt.Print("Masukkan jumlah yang ingin dibeli: ")
			fmt.Scanln(&jumlah)

			if B[i].stok >= jumlah {
				B[i].stok -= jumlah
				T[jumlahTransaksi].namaBarang = namaBarang
				T[jumlahTransaksi].jumlah = jumlah
				T[jumlahTransaksi].hargaTotal = B[i].harga * jumlah
				B[i].terjual += jumlah
				jumlahTransaksi++

				fmt.Println("Transaksi berhasil.")
			} else {
				fmt.Println("Stok tidak mencukupi.")
			}
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

func ubahdatatransaksi(T *arrTransaksi, B *arrBarang) {
	var namaBarang string
	var jumlahBaru int

	fmt.Println("Masukkan nama barang dari transaksi yang ingin diubah:")
	fmt.Scanln(&namaBarang)

	for i := 0; i < jumlahTransaksi; i++ {
		if T[i].namaBarang == namaBarang {
			for j := 0; j < jumlahBarang; j++ {
				if B[j].nama == T[i].namaBarang {
					B[j].stok += T[i].jumlah
					B[j].terjual -= T[i].jumlah
				}
			}

			fmt.Print("Masukkan jumlah baru: ")
			fmt.Scanln(&jumlahBaru)

			for j := 0; j < jumlahBarang; j++ {
				if B[j].nama == T[i].namaBarang {
					if B[j].stok >= jumlahBaru {
						B[j].stok -= jumlahBaru
						B[j].terjual += jumlahBaru
						T[i].jumlah = jumlahBaru
						T[i].hargaTotal = jumlahBaru * B[j].harga

						fmt.Println("Transaksi berhasil diubah.")
					} else {
						fmt.Println("Stok tidak mencukupi.")
					}
				}
			}
			return
		}
	}
	fmt.Println("Transaksi tidak ditemukan.")
}

func hapusdatatransaksi(T *arrTransaksi, B *arrBarang) {
	var namaBarang string

	fmt.Println("Masukkan nama barang dari transaksi yang ingin dihapus:")
	fmt.Scanln(&namaBarang)

	for i := 0; i < jumlahTransaksi; i++ {
		if T[i].namaBarang == namaBarang {
			for j := 0; j < jumlahBarang; j++ {
				if B[j].nama == T[i].namaBarang {
					B[j].stok += T[i].jumlah
					B[j].terjual -= T[i].jumlah
				}
			}

			for j := i; j < jumlahTransaksi-1; j++ {
				T[j] = T[j+1]
			}
			jumlahTransaksi--
			fmt.Println("Transaksi berhasil dihapus.")
			return
		}
	}
	fmt.Println("Transaksi tidak ditemukan.")
}

func selectionsortHarga(B *arrBarang) {
	for pass := 0; pass < jumlahBarang-1; pass++ {
		idx := pass
		for i := pass + 1; i < jumlahBarang; i++ {
			if B[i].harga < B[idx].harga {
				idx = i
			}
		}
		B[pass], B[idx] = B[idx], B[pass]
	}

	fmt.Println("Data barang terurut berdasarkan harga:")
	showAllData(B)
}

func insertionSortStok(B *arrBarang) {
	for pass := 1; pass < jumlahBarang; pass++ {
		temp := B[pass]
		i := pass - 1

		for i >= 0 && B[i].stok > temp.stok {
			B[i+1] = B[i]
			i--
		}
		B[i+1] = temp
	}

	fmt.Println("Data barang terurut berdasarkan stok:")
	showAllData(B)
}

func sequentialsearchNama(B *arrBarang) {
	var nama string
	fmt.Println("Masukkan nama barang yang dicari:")
	fmt.Scanln(&nama)

	for i := 0; i < jumlahBarang; i++ {
		if B[i].nama == nama {
			fmt.Println("Data ditemukan:")
			fmt.Println("Nama: ", B[i].nama)
			fmt.Println("Harga: ", B[i].harga)
			fmt.Println("Stok: ", B[i].stok)
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func binarySearchHarga(B *arrBarang) {
	var harga int
	fmt.Println("Masukkan harga barang yang dicari:")
	fmt.Scanln(&harga)

	left, right := 0, jumlahBarang-1

	for left <= right {
		mid := (left + right) / 2

		if B[mid].harga == harga {
			fmt.Println("Data ditemukan:")
			fmt.Println("Nama: ", B[mid].nama)
			fmt.Println("Harga: ", B[mid].harga)
			fmt.Println("Stok: ", B[mid].stok)
			return
		} else if B[mid].harga < harga {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func binarySearchStok(B *arrBarang) {
	var stok int
	fmt.Println("Masukkan stok barang yang dicari:")
	fmt.Scanln(&stok)

	left, right := 0, jumlahBarang-1

	for left <= right {
		mid := (left + right) / 2

		if B[mid].stok == stok {
			fmt.Println("Data ditemukan:")
			fmt.Println("Nama: ", B[mid].nama)
			fmt.Println("Harga: ", B[mid].harga)
			fmt.Println("Stok: ", B[mid].stok)
			return
		} else if B[mid].stok < stok {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func showAllData(B *arrBarang) {
	fmt.Println("Data barang:")
	for i := 0; i < jumlahBarang; i++ {
		fmt.Printf("Nama: %s, Harga: %d, Stok: %d, Terjual: %d\n", B[i].nama, B[i].harga, B[i].stok, B[i].terjual)
	}
}

func showAllDataTransaksi(T *arrTransaksi) {
	fmt.Println("Data transaksi:")
	for i := 0; i < jumlahTransaksi; i++ {
		fmt.Printf("Barang: %s, Jumlah: %d, Total: %d\n", T[i].namaBarang, T[i].jumlah, T[i].hargaTotal)
	}
}

func showTop5BarangTerlaris(B *arrBarang) {
	selectionsortTerjual(B)
	fmt.Println("5 Barang Terlaris:")
	for i := 0; i < 5 && i < jumlahBarang; i++ {
		fmt.Printf("Nama: %s, Terjual: %d\n", B[i].nama, B[i].terjual)
	}
}

func selectionsortTerjual(B *arrBarang) {
	for pass := 0; pass < jumlahBarang-1; pass++ {
		idx := pass
		for i := pass + 1; i < jumlahBarang; i++ {
			if B[i].terjual > B[idx].terjual {
				idx = i
			}
		}
		B[pass], B[idx] = B[idx], B[pass]
	}
}
