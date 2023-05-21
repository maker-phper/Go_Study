//this exercise will be improved later
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

func main() {
	if len(os.Args[1:]) > 0 {
		//get command-line arguments
		for _, arg := range os.Args[1:] {
			printValue(arg)
		}
	} else {
		//get the standard input
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			printValue(scan.Text())
		}
	}
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
func (f Feet) String() string {
	return fmt.Sprintf("%gfeets", f)
}
func (m Meter) String() string {
	return fmt.Sprintf("%gmeters", m)
}
func (p Pound) String() string {
	return fmt.Sprintf("%gPounds", p)
}
func (k Kilogram) String() string {
	return fmt.Sprintf("%gKilograms", k)
}

func printValue(arg string) {
	v, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	c := Celsius(v)
	fa := Fahrenheit(v)
	fe := Feet(v)
	m := Meter(v)
	p := Pound(v)
	k := Kilogram(v)
	fmt.Printf("%g, %s, %s, %s, %s, %s, %s\n", v, c, fa, fe, m, p, k)
}
