/*
==========
Cariddi
==========

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/cariddi

	@Author:      edoardottt, https://www.edoardoottavianelli.it
*/

package output

import (
	"fmt"

	"github.com/fatih/color"
)

//EncapsulateGreen >
func EncapsulateGreen(inp string) {
	// Create a custom print function for convenience
	green := color.New(color.FgGreen).PrintfFunc()
	green("[ + ] ")
	fmt.Println(inp)
}

//EncapsulateRed >
func EncapsulateRed(inp string) {
	// Create a custom print function for convenience
	red := color.New(color.FgRed).PrintfFunc()
	red("[ - ] ")
	fmt.Println(inp)
}

//EncapsulateYellow >
func EncapsulateYellow(inp string) {
	// Create a custom print function for convenience
	yellow := color.New(color.FgYellow).PrintfFunc()
	yellow("[ ? ] ")
	fmt.Println(inp)
}

//EncapsulateCustomGreen >
func EncapsulateCustomGreen(alert string, inp string) {
	// Create a custom print function for convenience
	green := color.New(color.FgGreen).PrintfFunc()
	green("[ %s ] ", alert)
	fmt.Println(inp)
}

//EncapsulateCustomRed >
func EncapsulateCustomRed(alert string, inp string) {
	// Create a custom print function for convenience
	red := color.New(color.FgRed).PrintfFunc()
	red("[ %s ] ", alert)
	fmt.Println(inp)
}

//EncapsulateCustomYellow >
func EncapsulateCustomYellow(alert string, inp string) {
	// Create a custom print function for convenience
	yellow := color.New(color.FgYellow).PrintfFunc()
	yellow("[ %s ] ", alert)
	fmt.Println(inp)
}
