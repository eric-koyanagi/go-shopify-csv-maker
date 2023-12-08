package getDummyRows

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

const MIN_PRICE float64 = 0.5
const MAX_PRICE float64 = 5000
const IMAGE_WIDTH int = 500
const IMAGE_HEIGHT int = 500
const MIN_VARIANT_GRAMS float64 = 0.5
const MAX_VARIANT_GRAMS float64 = 50000
const INVENTORY_MIN int = 0
const INVENTORY_MAX int = 1000000000
const SALE_CHANCE float64 = 0.3
const SALE_AMOUNT float64 = 0.5
const COST_PER_ITEM float64 = 0.4

var VARIANT_FULLFILMENT_SERVICES []string = []string{"manual", "manual", "manual"}

var Headers = []string{
	"Handle",
	"Title",
	"Body (HTML)",
	"Vendor",
	"Product Category",
	"Type",
	"Tags",
	"Published",
	"Option1 Name",
	"Option1 Value",
	"Option2 Name",
	"Option2 Value",
	"Option3 Name",
	"Option3 Value",
	"Variant SKU",
	"Variant Grams",
	"Variant Inventory Tracker",
	"Variant Inventory Qty",
	"Variant Inventory Policy",
	"Variant Fulfillment Service",
	"Variant Price",
	"Variant Compare At Price",
	"Variant Requires Shipping",
	"Variant Taxable",
	"Variant Barcode",
	"Image Src",
	"Image Position",
	"Image Alt Text",
	"Gift Card",
	"SEO Title",
	"SEO Description",
	"Google Shopping / Google Product Category",
	"Google Shopping / Gender",
	"Google Shopping / Age Group",
	"Google Shopping / MPN",
	"Google Shopping / AdWords Grouping",
	"Google Shopping / AdWords Labels",
	"Google Shopping / Condition",
	"Google Shopping / Custom Product",
	"Google Shopping / Custom Label 0",
	"Google Shopping / Custom Label 1",
	"Google Shopping / Custom Label 2",
	"Google Shopping / Custom Label 3",
	"Google Shopping / Custom Label 4",
	"Variant Image",
	"Variant Weight Unit",
	"Variant Tax Code",
	"Cost per item",
	"Price / International",
	"Compare At Price / International",
	"Status",
}

// Creates and returns a random product row as a map (order of the map keys doesn't matter)
func NewRow(handle string) map[string]string {
	isVariant := true
	if handle == "" {
		handle = gofakeit.BS() + "-" + gofakeit.BS()
		isVariant = false
	}

	var price float64 = gofakeit.Price(MIN_PRICE, MAX_PRICE)

	row := make(map[string]string)
	row["Handle"] = handle
	row["Title"] = gofakeit.ProductName()
	row["Body (HTML)"] = gofakeit.ProductDescription()
	row["Vendor"] = gofakeit.Company()
	row["Product Category"] = strconv.Itoa(gofakeit.Number(1, 5000))
	row["Type"] = gofakeit.ProductFeature()
	row["Tags"] = gofakeit.ProductFeature()
	row["Published"] = gofakeit.RandomString([]string{"true", "true", "false"})
	row["Variant SKU"] = gofakeit.UUID()
	row["Variant Price"] = strconv.FormatFloat(price, 'f', 2, 64)
	if gofakeit.Float64() <= SALE_CHANCE {
		row["Variant Compare At Price"] = strconv.FormatFloat(price*SALE_AMOUNT, 'f', 2, 64)
	}
	row["Variant Inventory Qty"] = strconv.Itoa(gofakeit.IntRange(INVENTORY_MIN, INVENTORY_MAX))
	row["Variant Inventory Policy"] = gofakeit.RandomString([]string{"deny", "deny", "continue"})
	row["Variant Fulfillment Service"] = gofakeit.RandomString(VARIANT_FULLFILMENT_SERVICES)
	row["Variant Barcode"] = gofakeit.ProductUPC()
	row["Image Src"] = gofakeit.ImageURL(IMAGE_WIDTH, IMAGE_HEIGHT)

	if isVariant {
		row["Option1 Name"] = "color"
		row["Option1 Value"] = gofakeit.Color()
		row["Variant Grams"] = strconv.FormatFloat(gofakeit.Float64Range(MIN_VARIANT_GRAMS, MAX_VARIANT_GRAMS), 'f', 2, 64)
		row["Variant Image"] = gofakeit.ImageURL(IMAGE_WIDTH, IMAGE_HEIGHT)
	}

	row["Cost per item"] = strconv.FormatFloat(price*COST_PER_ITEM, 'f', 2, 64)
	row["Status"] = gofakeit.RandomString([]string{"active", "active", "active", "draft", "archived"})

	return row
}

// Creates a new row, but returns it as a slice, ordered the same as the headers to ensure consistent output
func NewRowAsSlice(handle string) []string {
	orderedSlice := make([]string, len(Headers))
	row := NewRow(handle)

	for i, key := range Headers {
		orderedSlice[i] = row[key]
	}

	return orderedSlice
}
