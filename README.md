# SCG Nexter assignment


## Deploy on local docker
```sh
docker build -t scg-nexter-assignment . --build-arg CMD=assignment-server
docker run --rm -p 8080:8080 scg-nexter-assignment
```

## API Specs
+ ***[Search] Multiple search from datasets***
#### Example request
```sh
curl --location --request POST 'http://localhost:8080/search/multiple' \
--header 'Content-Type: application/json' \
--data-raw '{
    "keys": [
        "x",
        "X"
    ],
    "dataset": [
        "x",
        "X"
    ],
    "caseSensitive": false
}'
```
#### Example success response
HTTP Status 200
```json
    [
        { "key":"x", "positions":[0,1] },
        { "key":"X", "positions":[0,1] }
    ]
```

&nbsp;
+ ***[Cashier] Get currently money in cashier***
#### Example request
```sh
curl --location --request GET 'http://localhost:8080/cashier/money'
```
#### Example success response
HTTP Status 200
```json
{
    "0.25":50,
    "1":20,
    "10":20,
    "100":15,
    "1000":10,
    "20":30,
    "5":20,
    "50":20,
    "500":20
}
```

&nbsp;
+ ***[Cashier] Add money to cashier***
#### Example request
```sh
curl --location --request POST 'http://localhost:8080/cashier/money' \
--header 'Content-Type: application/json' \
--data-raw '{
    "value": 500,
    "amount": 2
}'
```
#### Example success response
HTTP Status 200 <br/>
No content

&nbsp;
+ ***[Cashier] Change money***
#### Example request
```sh
curl --location --request POST 'http://localhost:8080/cashier/change' \
--header 'Content-Type: application/json' \
--data-raw '{
    "price": 300.50,
    "moneyForBuy": {
        "100.00": 4
    }
}'
```
#### Example success response
HTTP Status 200
```json
{
    "0.25": 2,
    "1": 4,
    "20": 2,
    "5": 1,
    "50": 1
}
```