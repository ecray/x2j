# x2j
A commad-line stream converter for XML to JSON
... because working with XML sucks

Run Tests and Build
``` 
make
```

```cat example/test | ./x2j | jq '.'
{
  "customers": {
    "customer": [
      {
        "name": "blah",
        "transaction": "19.95",
        "id": "1"
      },
      {
        "id": "2",
        "name": "meh",
        "transaction": "6.50"
      },
      {
        "id": "4",
        "name": "yep",
        "transaction": "61.08"
      }
    ]
  }
}
```
