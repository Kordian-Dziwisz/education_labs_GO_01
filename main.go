package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func get_box_num(boxes_amount uint32) uint32 {
	return rand.Uint32() % boxes_amount
}

func chooseBox(box_num uint32, boxes_amount uint32, max_open_boxes uint32) func(bool) bool {
	real_box_num := get_box_num(boxes_amount)
	empty_boxes := make([]uint32, 0, max_open_boxes)
	for opened_boxes := uint32(0); opened_boxes < max_open_boxes; opened_boxes++ {
		empty_box_num := get_box_num(boxes_amount)
		for ; empty_box_num == real_box_num || empty_box_num == box_num || slices.Contains(empty_boxes, empty_box_num); empty_box_num = get_box_num(boxes_amount) {
		}
		empty_boxes = append(empty_boxes, empty_box_num)
	}
	return func(change_choice bool) bool {
		if change_choice {
			old_box_num := box_num
			box_num = get_box_num(boxes_amount)
			for ; old_box_num == box_num || slices.Contains(empty_boxes, box_num); box_num = get_box_num(boxes_amount) {
			}
		}
		return box_num == real_box_num
	}
}

func main() {
	liczba_rund := uint32(1e4)
	strategia_gracza := true
	boxes_amount := uint32(1e2)
	max_opened_boxes := uint32(98)

	wons, tries := 0.0, 0.0
	for i := liczba_rund; i > 0; i-- {
		change_choice := chooseBox(get_box_num(boxes_amount), boxes_amount, max_opened_boxes)
		if change_choice(strategia_gracza) {
			wons += 1
		}
		tries += 1
	}
	result := wons / tries
	fmt.Println(result)
}
