package services

import (
	"database/sql"
	config "mongo-server/Config"
	models "mongo-server/Models"
)

func CreateOrder(order models.Order) (int64, error) {
	result, err := config.MySQLDB.Exec(
		"INSERT INTO orders (user_id, product_name, quantity, price) VALUES (?, ?, ?, ?)",
		order.UserID, order.ProductName, order.Quantity, order.Price,
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetAllOrders() ([]models.Order, error) {
	rows, err := config.MySQLDB.Query("SELECT id, user_id, product_name, quantity, price, created_at FROM orders ORDER BY id DESC")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.UserID, &order.ProductName, &order.Quantity, &order.Price, &order.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func GetOrderById(id string) (models.Order, error) {
	var order models.Order

	err := config.MySQLDB.QueryRow("SELECT id, user_id, product_name, quantity, price, created_at FROM orders WHERE id = ?", id).Scan(
		&order.ID, &order.UserID, &order.ProductName, &order.Quantity, &order.Price, &order.CreatedAt,
	)

	if err != nil {
		return order, err
	}

	return order, nil
}

func GetOrdersByUserId(userId string) ([]models.Order, error) {
	rows, err := config.MySQLDB.Query("SELECT id, user_id, product_name, quantity, price, created_at FROM orders WHERE user_id = ?", userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []models.Order

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.UserID, &order.ProductName, &order.Quantity, &order.Price, &order.CreatedAt)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func UpdateOrderById(id string, updatedOrder models.Order) (models.Order, error) {
	result, err := config.MySQLDB.Exec(
		"UPDATE orders SET user_id = ?, product_name = ?, quantity = ?, price = ? WHERE id = ?",
		updatedOrder.UserID, updatedOrder.ProductName, updatedOrder.Quantity, updatedOrder.Price, id,
	)

	if err != nil {
		return models.Order{}, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return models.Order{}, err
	}
	if affected == 0 {
		return models.Order{}, sql.ErrNoRows
	}

	return GetOrderById(id)
}

func DeleteOrderById(id string) error {
	result, err := config.MySQLDB.Exec("DELETE FROM orders WHERE id = ?", id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}

	return err
}
