/*
	This file is part of InfyTape.

	Copyright (C) 2017  DolphinCommode

	InfyTape is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	InfyTape is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with InfyTape.  If not, see <http://www.gnu.org/licenses/>.
*/

package infytape

import (
	"fmt"
)

type InfyTape struct {
	mem []uint8
	cur int
}

func (tp *InfyTape)Init(size int) {
	tp.mem = make([]uint8, size, int(float64(size) * 1.5))
}

func (tp *InfyTape)Set(a []uint8) {
	tp.mem = a
	tp.cur = 0
}

func (tp *InfyTape)Left() (*uint8) {
	if tp.cur == 0 {
		tp.mem = append([]uint8{0}, tp.mem...)
	} else {
		tp.cur--
	}

	return &tp.mem[tp.cur]
}

func (tp InfyTape)Curr() (*uint8) {
	return &tp.mem[tp.cur]
}

func (tp *InfyTape)Right() (*uint8) {
	if tp.cur + 1 >= len(tp.mem) {
		tp.mem = append(tp.mem, 0)
	}

	tp.cur++
	return &tp.mem[tp.cur]
}

func (tp InfyTape)Display() {
	fmt.Printf("size: %d\ncap: %d\ncurr: %d\n", len(tp.mem), cap(tp.mem), tp.cur)
	for i, c := range tp.mem {
		fmt.Printf("%3d: %d\n", i, c)
	}
}