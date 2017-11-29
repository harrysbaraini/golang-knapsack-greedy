package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Define o tipo Item, que receberá
// os valores.

type Item struct {
	Name string
	Value int
	Weight int
	Ratio int
}

// Definimos o tipo ItemsSlice que será utilizado
// para criar um mapa de itens.
// Também já definimos as funções da interface de Sort,
// para organizarmos o array por taxa de valor / peso.

type ItemsSlice []*Item

func (d ItemsSlice) Len() int {
	return len(d)
}

func (d ItemsSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d ItemsSlice) Less(i, j int) bool {
	return d[i].Ratio > d[j].Ratio
}

// Função que calcula o Knapsack

func knapsack(maxWeight int, items ItemsSlice) (totalValue int, totalWeight int, selectedItems []string) {
	// criamos um array multidimensional do tamanho da quantidade de valores
	totalWeight = 0
	totalValue = 0

	for i := range items {
		if totalWeight + items[i].Weight <= maxWeight {
			totalWeight += items[i].Weight
			totalValue += items[i].Value
			selectedItems = append(selectedItems, items[i].Name + " (" + strconv.Itoa(items[i].Value) + " de " + strconv.Itoa(items[i].Value) + ") [ratio: " + strconv.Itoa(items[i].Ratio) + "]")
		} else {
			remains := maxWeight - totalWeight
			amount := int(float64(items[i].Value) * (float64(remains) / float64(items[i].Weight)))
			totalValue += amount
			selectedItems = append(selectedItems, items[i].Name + " (" + strconv.Itoa(amount) + " de " + strconv.Itoa(items[i].Value) + ") [ratio: " + strconv.Itoa(items[i].Ratio) + "]")
			break
		}
	}

	return
}

func main() {
	names := []string{"Londres", "Paris", "Dublin", "Amsterdam", "Lisboa", "Madrid", "Barcelona", "Ibiza"}
	values := []int{40, 25, 17, 35, 38, 34, 54, 48}
	weights := []int{49, 35, 15, 43, 38, 37, 21, 37}
	maxWeight := 137

	items := make(ItemsSlice, 0, len(values))

	for i := range values {
		items = append(items, &Item{
			Name: names[i],
			Value: values[i],
			Weight: weights[i],
			Ratio: int(float32(values[i]) / float32(weights[i])),
		})
	}

	sort.Sort(items)

	totalValue, totalWeight, selectedItems := knapsack(maxWeight, items)

	fmt.Printf("A lista de valores é %v\n", values)
	fmt.Printf("A lista de peso é %v\n", weights)
	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
	fmt.Printf("Os itens selecionados são: \n\n")

	for i := range selectedItems {
		fmt.Printf("\t%s\n", selectedItems[i])
	}

	fmt.Printf("\nPara capacidade %d, O peso total é %d e o valor máximo é %d\n", maxWeight, totalWeight, totalValue)
}