/*
	This file is part of TurTest.

	Copyright (C) 2017  DolphinCommode

	TurTest is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	TurTest is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with TurTest.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"

	"turcore"
)

func main() {
	var tmp turcore.TurMachine

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
