# x2j
A commad-line stream converter for XML to JSON
... because working with XML sucks

Run Tests and Build
``` 
$ make
```

```
$ cat example/test.xml
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
</customers>

$ cat example/test.xml | ./x2j | jq '.'
{
  "customers": {
    "customer": [
      {
        "order": {
          "id": "666eefff",
          "product": {
            "id": "34223",
            "name": "vin"
          }
        },
        "transaction": "19.95",
        "id": "1",
        "name": "blah"
      },
      {
        "id": "2",
        "name": "meh",
        "order": {
          "id": "edf7f1ad",
          "product": {
            "id": "2222",
            "name": "øl"
          }
        },
        "transaction": "6.50"
      },
      {
        "id": "4",
        "name": "yep",
        "order": {
          "id": "dd27f6f8",
          "product": {
            "id": "11111",
            "name": "log"
          }
        },
        "transaction": "61.08"
      }
    ]
  }
}
```
