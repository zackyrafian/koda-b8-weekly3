package service

import "seindonesia/internal/domain"

var orderList []domain.OrderItem

func AddOrder(order domain.OrderItem) {
    orderList = append(orderList, order)
}

func GetOrders() []domain.OrderItem {
    return orderList
}

func ClearOrders() {
    orderList = []domain.OrderItem{}
}

