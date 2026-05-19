package main

import "fmt"

type Item struct {
	ID    int
	Name  string
	Price float64
}

func findSliceIndex(menu []Item, id int) int {
	for i, item := range menu {
		if item.ID == id {
			return i
		}
	}
	return -1
}

func main() {
	var menu []Item

	menu = append(menu, Item{1, "Espresso", 30000}, Item{2, "Bac Xiu", 35000})

	for {
		fmt.Print("\n[SLICE]\n1.Xem\n2.Thêm\n3.Sửa\n4.Xóa\n5.Thoát\nChọn: ")
		var choice, id int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if len(menu) == 0 {
				fmt.Println("Menu trống!")
				continue
			}
			for _, item := range menu {
				fmt.Printf("ID: %d | %-12s | %.0fđ\n", item.ID, item.Name, item.Price)
			}
		case 2:
			var item Item
			fmt.Print("Nhập ID, Tên, Giá: ")
			fmt.Scanln(&item.ID, &item.Name, &item.Price)

			if findSliceIndex(menu, item.ID) != -1 {
				fmt.Println("ID đã tồn tại!")
				continue
			}
			menu = append(menu, item)
			fmt.Println("Thêm thành công!")
		case 3:
			fmt.Print("Nhập ID cần sửa: ")
			fmt.Scanln(&id)
			if idx := findSliceIndex(menu, id); idx != -1 {
				fmt.Print("Nhập Tên và Giá mới: ")
				fmt.Scanln(&menu[idx].Name, &menu[idx].Price)
				fmt.Println("Sửa thành công!")
			} else {
				fmt.Println("Không tìm thấy!")
			}
		case 4:
			fmt.Print("Nhập ID cần xóa: ")
			fmt.Scanln(&id)
			if idx := findSliceIndex(menu, id); idx != -1 {
				menu = append(menu[:idx], menu[idx+1:]...)
				fmt.Println("Xóa thành công!")
			} else {
				fmt.Println("Không tìm thấy!")
			}
		case 5:
			return
		}
	}
}
