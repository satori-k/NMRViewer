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
)

// These variables are later read from procs file.

const ABSF1 float64 = 16.39466
const ABSF2 float64 = -4.090485
const FTSIZE int = 65536
const SF float64 = 400.13000916893

// read file from directory and return an array.
func readfile(dir string) ([]float64, []float64) {
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
	xaxis, yaxis := readfile("./1r")
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
	line.LineStyle.Color = color.Black
	p.Add(line)
	if err != nil {
		fmt.Println(err)
	}
	err = p.Save(800, 800, "graph.png")
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	drawPlot()
}
