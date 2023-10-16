package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/montanaflynn/stats"
)

// a function read csv file and to df
func loadDF(filePath string) dataframe.DataFrame {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("cant open it")
	}
	defer file.Close()

	mydataframe := dataframe.ReadCSV(file)

	return mydataframe
}

// create stat summary
func statSummary(df dataframe.DataFrame) string {
	colNN := df.Names() // create a slice of column name
	//fmt.Println(colNN)
	var statSum string
	//df.Names()  = colName := []string{"value", "income", "age", "rooms", "bedrooms", "pop", "hh"}
	for _, colN := range colNN {
		colData := df.Col(colN)
		var data []float64
		switch colData.Type() {
		case series.Float:
			data = colData.Float()
		case series.Int:
			intData, err := colData.Int()
			if err != nil {
				fmt.Println("something wrong loading intgers")
				//continue
			}
			for _, val := range intData {
				data = append(data, float64(val))
			}
		}
		//can't use describe func percentiles := []float64{25.0, 50.0, 75.0}
		//describs, _ := stats.Describe(data, false, &percentiles)
		//summary of the data
		count := len(data)
		mean, _ := stats.Mean(data)
		median, _ := stats.Median(data)
		max, _ := stats.Max(data)
		min, _ := stats.Min(data)
		stddev, _ := stats.StandardDeviation(data)
		statSum += fmt.Sprintf("Column: %s\ncount: %v\nMean: %f\nMedian: %f\nMax: %f\nMin: %f\nStandard Deviation: %f\n\n\n ",
			colN, count, mean, median, max, min, stddev)

	}
	return statSum

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	//ask user input
	fmt.Println("Welcome to data summary application!\nNow, please enter the path of your CSV file")
	dataPath, _ := reader.ReadString('\n')
	dataPath = strings.TrimSpace(dataPath)

	//howmany interations?
	fmt.Println("How many time you want this summary to run?")
	dataR, _ := reader.ReadString('\n')
	dataR = strings.TrimSpace(dataR)
	iteRations, err := strconv.Atoi(dataR)
	if err != nil {
		fmt.Println("The input iteration is invalid, please reenter")
	}

	var statSum11 string //create new var to store all output data
	//create output
	file, err := os.Create("outputGo.txt")
	if err != nil {
		fmt.Println("cant output the txt")

	}
	defer file.Close()
	//iterate 100 times to read all data into file and find processing time
	startTime := time.Now()
	for i := 0; i < iteRations; i++ {
		df := loadDF(dataPath)
		//fmt.Println(df)
		statSum11 += statSummary(df)
		//fmt.Println(statSum11)
		file.WriteString(statSum11)
	}
	elaspedTime := time.Since(startTime)
	fmt.Printf("The Golang code processing time:%v", elaspedTime)

}
