package main

import "fmt"

type CafeItem struct {
	Name  string
	Price float64
}

func main() {
	menu := make(map[int]CafeItem)

	menu[1] = CafeItem{"Espresso", 30000}
	menu[2] = CafeItem{"Bac Xiu", 35000}

	for {
		fmt.Print("\n[MAP]\n1.Xem\n2.Thêm\n3.Sửa\n4.Xóa\n5.Thoát\nChọn: ")
		var choice, id int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if len(menu) == 0 {
				fmt.Println("Menu trống!")
				continue
			}
			for id, item := range menu {
				fmt.Printf("ID: %d | %-12s | %.0fđ\n", id, item.Name, item.Price)
			}
		case 2:
			var name string
			var price float64
			fmt.Print("Nhập ID, Tên, Giá: ")
			fmt.Scanln(&id, &name, &price)

			if _, exists := menu[id]; exists {
				fmt.Println("Lỗi: ID này đã tồn tại!")
				continue
			}
			menu[id] = CafeItem{Name: name, Price: price}
			fmt.Println("Thêm thành công!")
		case 3:
			fmt.Print("Nhập ID cần sửa: ")
			fmt.Scanln(&id)

			if _, exists := menu[id]; exists {
				var item CafeItem
				fmt.Print("Nhập Tên và Giá mới: ")
				fmt.Scanln(&item.Name, &item.Price)
				menu[id] = item
				fmt.Println("Sửa thành công!")
			} else {
				fmt.Println("Không tìm thấy!")
			}
		case 4:
			fmt.Print("Nhập ID cần xóa: ")
			fmt.Scanln(&id)

			if _, exists := menu[id]; exists {
				delete(menu, id)
				fmt.Println("Xóa thành công!")
			} else {
				fmt.Println("Không tìm thấy!")
			}
		case 5:
			return
		}
	}
}
