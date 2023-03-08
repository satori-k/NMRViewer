package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"image/color"
	"io"
	"os"
	"strconv"
	"strings"
)

// These variables are later read from procs file (for 1r of 1a file)
// or acqus file (for fid file)

var ABSF1 float64

var ABSF2 float64
var FTSIZE int
var SF float64

// DTYPP or DTYPA in these files.
var DTYPE int

func readconfig(dir string) {
	f, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		if err != nil {
			print(err)
		}
		var value string
		if len(items) == 2 {
			switch items[0] {
			case "##$ABSF1=":
				value, _ := strconv.ParseFloat(items[1], 64)
				ABSF1 = value
			case "##$ABSF2=":
				value, _ := strconv.ParseFloat(items[1], 64)
				ABSF2 = value
			case "##$FTSIZE=":
				value, _ := strconv.ParseFloat(items[1], 64)
				FTSIZE = int(value)
			case "##$SF=":
				value, _ := strconv.ParseFloat(items[1], 64)
				SF = value
			case "##$DTYPP=":
				value, _ := strconv.ParseFloat(items[1], 64)
				DTYPE = int(value)
			}
			fmt.Println(value)
		}
	}

}

// read file from directory and return an array.
func readfile(dir string) ([]float64, []float64) {
	readconfig("procs")
	f, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReader(f)
	data := []float64{}
	ppm := []float64{}

	for i := 0; i < FTSIZE; i++ {
		var num int32
		err = binary.Read(r, binary.LittleEndian, &num)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		// Generate the Y value (strength).
		data = append(data, float64(num))

		// Generate X axis data (chemical shift)
		x_value := ABSF1 - float64(i)*(ABSF1-ABSF2)/float64(FTSIZE)
		ppm = append(ppm, x_value)
	}
	return ppm, data
}

func generatePoints() plotter.XYs {
	points := make(plotter.XYs, FTSIZE)
	xaxis, yaxis := readfile("1r")
	for i := 0; i < FTSIZE; i++ {
		points[i].X = xaxis[i]
		points[i].Y = yaxis[i]
	}
	return points
}

func drawPlot() {
	//fmt.Println(xaxis)
	p := plot.New()
	p.Title.Text = "NMRViewer"
	p.X.Label.Text = "Chemical Shift (ppm)"
	p.Y.Label.Text = "Strength"
	line, err := plotter.NewLine(generatePoints())
	if err != nil {
		fmt.Println(err)
	}

	line.LineStyle.Color = color.Black
	p.Add(line)
	p.X.Scale = plot.InvertedScale{Normalizer: p.X.Scale}

	err = p.Save(800, 800, "graph.png")
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	drawPlot()
}
