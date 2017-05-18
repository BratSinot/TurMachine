/*
	This file is part of TurCore.

	Copyright (C) 2017  DolphinCommode

	TurCore is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	TurCore is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with TurCore.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"infytape"

	"fmt"
	"log"
	"os"
	"bufio"
	"regexp"
	"strconv"
)

const (
	tLeft = iota
	tCenter
	tRight
)

type slot struct {
	move uint8
	set uint8
	nextSost int
}

type TurMachine struct {
	//Tape infytape.InfyTape
	infytape.InfyTape

	cmd [][2]slot
	sost int
	add_sost int
}

func (tm *TurMachine)Add(q int, curr int, move uint8, set uint8, nextSost int) {
	for len(tm.cmd) <= q {
		tmp := [2]slot{{}}
		tm.cmd = append(tm.cmd, tmp)
	}

	tm.cmd[q][curr].move = move
	tm.cmd[q][curr].set = set
	tm.cmd[q][curr].nextSost = nextSost
}

func (tm *TurMachine)FromFile(fstr string) {
	tm.cmd = nil

	tm.sost = 1 //q1
// ---------------------------------- //
	_fin, err := os.Open(fstr)
	if err != nil {
		log.Fatal(err)
	}
	defer _fin.Close()

	fin := bufio.NewScanner(_fin)
	fin.Split(bufio.ScanLines)

	rgx := regexp.MustCompile(`([01])\s*q([0-9]*)\s*->\s*([01])\s*q([0-9])*\s*([LCR])`)
	for fin.Scan() {
		arr := rgx.FindAllStringSubmatch(fin.Text(), -1)

		//Add(q int, curr int, move uint8, set int8, nextSost int)
		q, _        := strconv.Atoi(arr[0][2])
		curr, _     := strconv.Atoi(arr[0][1])
		move        := 0
		switch arr[0][5] {
			case "L":
				move = tLeft
			case "C":
				move = tCenter
			case "R":
				move = tRight
		}
		set, _	    := strconv.Atoi(arr[0][3])
		nextSost, _ := strconv.Atoi(arr[0][4])
		tm.Add(q, curr, uint8(move), uint8(set), nextSost)
	}
}

func (tm *TurMachine)move(m uint8) {
	switch m {
		case tLeft:
			tm.Left()
		case tCenter:
			//tm.Center()
		case tRight:
			tm.Right()
	}
}

func (tm *TurMachine)Execute() {
	for tm.sost != 0 /* q0 */ {
		curr := tm.Curr()
		switch *curr {
			case 0:
				*curr = tm.cmd[tm.sost][0].set
				tm.move(tm.cmd[tm.sost][0].move)
				tm.sost = tm.cmd[tm.sost][0].nextSost
			case 1:
				*curr = tm.cmd[tm.sost][1].set
				tm.move(tm.cmd[tm.sost][1].move)
				tm.sost = tm.cmd[tm.sost][1].nextSost
		}
	}
}

func main() {
	var tmp TurMachine

	/*tmp.FromFile("in.txt")
	tmp.Set([]int8{1,1,1,0,1,0,1,1})
	tmp.Execute()
	tmp.Display()*/

	fmt.Println("INC:")
	tmp.FromFile("examples/inc.txt")
	tmp.Set([]uint8{1,1,1,1})
	tmp.Execute()
	tmp.Display()

	fmt.Println("\nSUM:")
	tmp.FromFile("examples/sum.txt")
	tmp.Set([]uint8{1,1,1,1,0,1,1})
	tmp.Execute()
	tmp.Display()

	fmt.Println("\nX1:")
	tmp.FromFile("examples/x1.txt")
	tmp.Set([]uint8{1,1,0,1,1,0,1,1,1,1})
	tmp.Execute()
	tmp.Display()

	fmt.Println("\nZERO:")
	tmp.FromFile("examples/zero.txt")
	tmp.Set([]uint8{1,1,1,1})
	tmp.Execute()
	tmp.Display()

	fmt.Println("\nZERON:")
	tmp.FromFile("examples/zeroN.txt")
	tmp.Set([]uint8{1,1,0,1,1,0,1,1,1,1})
	tmp.Execute()
	tmp.Display()
}