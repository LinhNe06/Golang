package main

import "fmt"

type Item struct {
	ID    int
	Name  string
	Price float64
}

const Max = 10

var menu [Max]Item
var size = 0

func findIndex(id int) int {
	for i := 0; i < size; i++ {
		if menu[i].ID == id {
			return i
		}
	}
	return -1
}

func main() {
	menu[0] = Item{1, "Espresso", 30000}
	menu[1] = Item{2, "Bac Xiu", 35000}
	size = 2

	for {
		fmt.Print("\n1.Xem\n2.Thêm\n3.Sửa\n4.Xóa\n5.Thoát\nChọn: ")
		var choice, id int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if size == 0 {
				fmt.Println("Menu trống!")
				continue
			}
			for i := 0; i < size; i++ {
				fmt.Printf("ID: %d | %-12s | %.0fđ\n", menu[i].ID, menu[i].Name, menu[i].Price)
			}
		case 2:
			if size >= Max {
				fmt.Println("Menu đầy!")
				continue
			}
			fmt.Print("Nhập ID, Tên, Giá: ")
			var item Item
			fmt.Scanln(&item.ID, &item.Name, &item.Price)

			if findIndex(item.ID) != -1 {
				fmt.Println("ID đã tồn tại!")
				continue
			}
			menu[size] = item
			size++
		case 3:
			fmt.Print("Nhập ID cần sửa: ")
			fmt.Scanln(&id)
			if idx := findIndex(id); idx != -1 {
				fmt.Print("Nhập Tên mới và Giá mới: ")
				fmt.Scanln(&menu[idx].Name, &menu[idx].Price)
			} else {
				fmt.Println("Không tìm thấy!")
			}
		case 4:
			fmt.Print("Nhập ID cần xóa: ")
			fmt.Scanln(&id)
			if idx := findIndex(id); idx != -1 {
				menu[idx] = menu[size-1]
				menu[size-1] = Item{}
				size--
				fmt.Println("Xóa thành công!")
			} else {
				fmt.Println("Không tìm thấy!")
			}
		case 5:
			return
		}
	}
}
