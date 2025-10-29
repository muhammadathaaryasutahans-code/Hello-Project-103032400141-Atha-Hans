// Ditambahkan oleh Muhammad Atha Aryasuta Hans 103032400141

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Struct untuk menyimpan data tim
type Team struct {
	Name    string
	Players []string
	Wins    int
	Losses  int
}

// Struct untuk menyimpan hasil pertandingan
type Match struct {
	Team1  string
	Team2  string
	Score1 int
	Score2 int
	Date   string
}

// Array untuk menyimpan data tim dan pertandingan
var teams []Team
var matches []Match

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== APLIKASI PENGELOLAAN TURNAMEN ESPORT ===")
		fmt.Println("1. Tambah Data Tim")
		fmt.Println("2. Tambah Hasil Pertandingan")
		fmt.Println("3. Tampilkan Daftar Tim")
		fmt.Println("4. Tampilkan Hasil Pertandingan")
		fmt.Println("5. Cari Tim")
		fmt.Println("6. Tampilkan Klasemen")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			tambahTim(scanner)
		case "2":
			tambahPertandingan(scanner)
		case "3":
			tampilkanTim()
		case "4":
			tampilkanPertandingan()
		case "5":
			cariTim(scanner)
		case "6":
			tampilkanKlasemen()
		case "0":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk menambahkan data tim
func tambahTim(scanner *bufio.Scanner) {
	fmt.Print("Nama tim: ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("Jumlah pemain: ")
	scanner.Scan()
	jumlahStr := scanner.Text()
	jumlah, _ := strconv.Atoi(jumlahStr)

	var pemain []string
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Nama pemain %d: ", i+1)
		scanner.Scan()
		pemain = append(pemain, scanner.Text())
	}

	tim := Team{nama, pemain, 0, 0}
	teams = append(teams, tim)
	fmt.Println("Tim berhasil ditambahkan!")
}

// Fungsi untuk menambahkan hasil pertandingan dan memperbarui jumlah menang/kalah
func tambahPertandingan(scanner *bufio.Scanner) {
	fmt.Print("Nama tim 1: ")
	scanner.Scan()
	t1 := scanner.Text()
	fmt.Print("Skor tim 1: ")
	scanner.Scan()
	s1, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Nama tim 2: ")
	scanner.Scan()
	t2 := scanner.Text()
	fmt.Print("Skor tim 2: ")
	scanner.Scan()
	s2, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Tanggal pertandingan (YYYY-MM-DD): ")
	scanner.Scan()
	tgl := scanner.Text()

	matches = append(matches, Match{t1, t2, s1, s2, tgl})

	// Update win/loss untuk masing-masing tim
	for i := range teams {
		if teams[i].Name == t1 {
			if s1 > s2 {
				teams[i].Wins++
			} else {
				teams[i].Losses++
			}
		} else if teams[i].Name == t2 {
			if s2 > s1 {
				teams[i].Wins++
			} else {
				teams[i].Losses++
			}
		}
	}

	fmt.Println("Hasil pertandingan berhasil ditambahkan!")
}

// Fungsi untuk menampilkan semua tim yang sudah ditambahkan
func tampilkanTim() {
	fmt.Println("\n--- DAFTAR TIM ---")
	for _, t := range teams {
		fmt.Printf("Nama: %s\n", t.Name)
		fmt.Printf("Pemain: %s\n", strings.Join(t.Players, ", "))
		fmt.Printf("Menang: %d, Kalah: %d\n\n", t.Wins, t.Losses)
	}
}

// Fungsi untuk menampilkan semua hasil pertandingan yang telah dicatat
func tampilkanPertandingan() {
	fmt.Println("\n--- HASIL PERTANDINGAN ---")
	for _, m := range matches {
		fmt.Printf("%s vs %s | %d - %d | %s\n", m.Team1, m.Team2, m.Score1, m.Score2, m.Date)
	}
}

// Fungsi untuk mencari data tim berdasarkan nama
func cariTim(scanner *bufio.Scanner) {
	fmt.Print("Masukkan nama tim yang dicari: ")
	scanner.Scan()
	nama := scanner.Text()

	found := false
	for _, t := range teams {
		if strings.ToLower(t.Name) == strings.ToLower(nama) {
			fmt.Printf("\nNama: %s\nPemain: %s\nMenang: %d, Kalah: %d\n", t.Name, strings.Join(t.Players, ", "), t.Wins, t.Losses)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Tim tidak ditemukan.")
	}
}

// Fungsi untuk menampilkan klasemen berdasarkan jumlah kemenangan
func tampilkanKlasemen() {
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Wins > teams[j].Wins
	})
	fmt.Println("\n--- KLASEMEN ---")
	for i, t := range teams {
		fmt.Printf("%d. %s - Menang: %d | Kalah: %d\n", i+1, t.Name, t.Wins, t.Losses)
	}
}
