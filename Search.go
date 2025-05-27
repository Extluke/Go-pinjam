package main

import "fmt"

// untuk func sequential search
func sequentialSearch(data []Nasabah, target string) int {
	for i, n := range data {
		if n.Nama == target {
      return i
		}
	}

  fmt.Println("❌ Data tidak ditemukan")
	return -1 
}
