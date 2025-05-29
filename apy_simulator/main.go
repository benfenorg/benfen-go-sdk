package main

import (
	"fmt"
	"github.com/wcharczuk/go-chart"
	"os"
)

//stake_subsidy_start_epoch: 0
//stake_subsidy_initial_distribution_amount: 100000*1000000000
//stake_subsidy_period_length: 10
//stake_subsidy_decrease_rate: 1000

var epoch_time = 6 * 60 * 60 // 6 hours
var simulator_time = 365 * 24 * 60 * 60 * 7

var total_epoch_count = simulator_time / epoch_time
var epoch_counts_per_month = (int)(30 * 24 * 60 * 60 / epoch_time)
var epoch_counts_per_year float64 = (float64)(365 * 24 * 60 * 60 / epoch_time)

var validator_count float64 = 18
var validator_stake_bfc_count float64 = 200000

var start_subsidy_amount float64 = 4500
var subsidy_period_length = 10
var subsidy_decrease_rate = 0.0 // 0.5% change to 0.0, subsidy will not decrease

var total_subsidy_amount float64 = 0

var bfc_generate_by_nft = 0
var bfc_generate_rate_per_second float64 = 3
var two_years_seconds float64 = 2 * 365 * 24 * 60 * 60
var nfc_bfc_convert_stake_rate = 0.05

func OutputNftBfcAmount(index int) float64 {
	local_rate_per_second := bfc_generate_rate_per_second
	total_time := (float64)(index) * (float64)(epoch_time)
	day := (int)(total_time / (24 * 60 * 60))
	year := day / 365
	year_left := day % 365
	if year_left > 0 {
		year = year + 1
	}

	//nft reduce 50% every 2 years.
	var total_amount float64
	if year <= 2 {
		total_amount = (float64)(local_rate_per_second) * total_time
		return total_amount
	} else if year <= 4 {
		//first 2 year
		total_amount = (float64)(local_rate_per_second * two_years_seconds)
		total_amount = total_amount + (float64)(local_rate_per_second/2.0)*(total_time-(float64)(two_years_seconds))
		return total_amount
	} else if year <= 6 {
		//first 4 year
		total_amount = (float64)(local_rate_per_second * two_years_seconds)
		total_amount = total_amount + (float64)((local_rate_per_second/2.0)*two_years_seconds)
		total_amount = total_amount + (float64)(
			(local_rate_per_second/4.0)*(total_time-(float64)(
				two_years_seconds*2,
			)),
		)
		return total_amount
	} else if year <= 8 {
		//first 6 year
		total_amount = (float64)(local_rate_per_second * two_years_seconds)
		total_amount = total_amount + (float64)((local_rate_per_second/2.0)*two_years_seconds)
		total_amount = total_amount + (float64)((local_rate_per_second/4.0)*two_years_seconds)
		total_amount = total_amount + (float64)(local_rate_per_second/8.0)*(total_time-(float64)(two_years_seconds*3))
		return total_amount
	} else if year <= 10 {
		//first 8 year
		total_amount = (float64)(local_rate_per_second * two_years_seconds)
		total_amount = total_amount + (float64)((local_rate_per_second/2)*two_years_seconds)
		total_amount = total_amount + (float64)((local_rate_per_second/4)*two_years_seconds)
		total_amount = total_amount + (float64)((local_rate_per_second/8)*two_years_seconds)
		total_amount = total_amount + (float64)(local_rate_per_second/16)*(total_time-(float64)(two_years_seconds*4))
		return total_amount
	} else {
		//first 10 year
		total_amount = (float64)(local_rate_per_second * two_years_seconds)
		total_amount = total_amount + (float64)((local_rate_per_second/2)*two_years_seconds)
		total_amount = total_amount + (float64)((local_rate_per_second/4)*two_years_seconds)
		total_amount = total_amount + (float64)((local_rate_per_second/8)*two_years_seconds)
		total_amount = total_amount + (float64)((local_rate_per_second/16)*two_years_seconds)
		total_amount = total_amount + (float64)(local_rate_per_second/32)*(total_time-(float64)(two_years_seconds*5))
		return total_amount
	}
}
func main() {
	println("Hello, World!")

	println(total_epoch_count)
	fmt.Printf("epoch_counts_per_year: %.2f \n", epoch_counts_per_year)

	apy_data := make([]float64, total_epoch_count)

	index := 0
	for index < total_epoch_count {

		index = index + 1
		if index%subsidy_period_length == 0 {
			start_subsidy_amount = start_subsidy_amount - start_subsidy_amount*subsidy_decrease_rate
		}

		total_stake_bfc := validator_count * validator_stake_bfc_count

		//assume all the subsidy amount is staked
		total_stake_bfc = total_stake_bfc + total_subsidy_amount

		//assume 5% th nft bfc is staked
		//todo
		nft_generate_bfc := OutputNftBfcAmount(index)
		//fmt.Printf("index: %d nft_generate_bfc: %.2f \n", index, nft_generate_bfc)
		convert_stake_nft_generate_bfc := nft_generate_bfc * nfc_bfc_convert_stake_rate
		total_stake_bfc = total_stake_bfc + convert_stake_nft_generate_bfc

		current_apy := (start_subsidy_amount * epoch_counts_per_year) / total_stake_bfc

		apy_data[index-1] = float64(current_apy)

		total_subsidy_amount = total_subsidy_amount + start_subsidy_amount
	}

	outputApy(apy_data)
	outputApyMonth(apy_data)
	fmt.Printf("total_subsidy_amount: %.2f", total_subsidy_amount)

}

func outputApy(data []float64) {
	// 创建一个新的图表对象
	graph := chart.Chart{
		Width:  800,
		Height: 600,
		XAxis: chart.XAxis{
			Name:      "Time Axis",
			NameStyle: chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "APY Axis",
			NameStyle: chart.StyleShow(),
		},
	}
	x_value := make([]float64, len(data))
	y_value := make([]float64, len(data))
	//value_labels := make([]chart.Value2, len(data)/50)

	for index, value := range data {
		total_time := (float64)(index) * (float64)(epoch_time)
		day := (int)(total_time / (24 * 60 * 60))
		year := (float64)((float64)(day) / 365.0)

		//day = day % 30
		if index%10 == 0 {
			fmt.Print("index: ", index)
			fmt.Print(" day: ", day)
			fmt.Printf(" year: %.2f", year)

			fmt.Printf(" APY: %.2f%%\n", value*100)
		}

		x_value[index] = (float64)(index)
		y_value[index] = value * 100

	}

	// 添加数据点
	series := chart.ContinuousSeries{
		Name:    "Data",
		XValues: x_value,
		YValues: y_value,
	}

	// 将数据系列添加到图表中
	graph.Series = []chart.Series{series}

	// 保存图表为PNG文件
	file, _ := os.Create("chart.png")
	defer file.Close()
	graph.Render(chart.PNG, file)
}

func outputApyMonth(data []float64) {
	// 创建一个新的图表对象
	graph := chart.Chart{
		Width:  800,
		Height: 600,
		XAxis: chart.XAxis{
			Name:      "Time Axis",
			NameStyle: chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "APY Axis",
			NameStyle: chart.StyleShow(),
		},
	}

	monthCount := len(data) / (int)(epoch_counts_per_month)
	x_value := make([]float64, monthCount+1)
	y_value := make([]float64, monthCount+1)

	count := 0
	array_index := 0

	for index, value := range data {
		total_time := (float64)(index) * (float64)(epoch_time)
		day := (int)(total_time / (24 * 60 * 60))
		month := day / 30
		//year := month / 12

		day = day % 30
		month = month % 12

		if count%120 == 0 {
			x_value[array_index] = (float64)(array_index)
			y_value[array_index] = value * 100
			array_index = array_index + 1
		}

		count = count + 1

	}

	// 添加数据点
	series := chart.ContinuousSeries{
		Name:    "Data",
		XValues: x_value,
		YValues: y_value,
	}

	// 将数据系列添加到图表中
	graph.Series = []chart.Series{series}

	// 保存图表为PNG文件
	file, _ := os.Create("chartMonth.png")
	defer file.Close()
	graph.Render(chart.PNG, file)
}
