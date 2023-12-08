# Go Shopify CSV Maker
This creates a CSV filled with dummy product data for Shopify. 

It can generate any number of products, but be aware that Shopify has a 15 MB limit. You can use another tool (like this project https://github.com/eric-koyanagi/go-shopify-csv-helper) to split up larger CSVs -- I allow it to create large CSVs so I can easily test a tool like this. 

# How to Use
Update the constants in main.go to control how many products to generate and the name of the output file. 

Look into the "getDummyRows" package and the "get.go" file inside that package -- there are a large variety of constants you can change to control how it generates files. 

Most of these are fairly self-explanitory. 

- SALE_CHANCE: sets a "compare at" price for an item, by default an on-sale item is 50% off
- SALE_AMOUNT: how much more expensive the original item is, defaults to 1.5 
- COST_PER_ITEM: the percent of the product price that is "cost", defaulting to 40%

# Need more fields?
Add them to the NewRow method. You don't need to worry about the order of the row -- the map will automaticlly be re-ordered as it's turned into a slice, based on the keys as defined in the header. 

Your field does need to match the header key; these header keys were taken from Shopify's sample CSV file exactly. 