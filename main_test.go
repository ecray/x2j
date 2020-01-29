package main

import (
	"bytes"
	"encoding/json"
	"testing"

	xj "github.com/basgys/goxml2json"
	"github.com/stretchr/testify/assert"
)

type Customers struct {
	customer []Customer
}

type Customer struct {
	ID          int
	name        string
	order       []Order
	transaction float64
}

type Order struct {
	ID      int
	product int64
	name    string
}

func TestConvStream(t *testing.T) {
	assert := assert.New(t)
	var customers []Customers
	r := []byte(`
	<?xml version="1.0" encoding="UTF-8"?>
	<customers>
		<customer id="1">
			<name>blah</name>
			<order id="666eefff">
				<product id="34223">
				<name>vin</name>
				</product>
			</order>
			<transaction>19.95</transaction>
		</customer>
		<customer id="2">
			<name>meh</name>
			<order id="edf7f1ad">
				<product id="2222">
				<name>øl</name>
				</product>
			</order>
			<transaction>6.50</transaction>
		</customer>
		<customer id="4">
			<name>yep</name>
			<order id="dd27f6f8">
				<product id="11111">
				<name>log</name>
				</product>
			</order>
			<transaction>61.08</transaction>
		</customer>
	</customers>`)

	c, err := xj.Convert(bytes.NewReader(r), xj.WithAttrPrefix(""))
	if err != nil {
		t.Error("Failed to convert")
	}
	json.Unmarshal(c.Bytes(), &customers)
	assert.Nil(customers)

}
