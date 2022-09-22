### Temporal Go Demo

```sh
curl --request POST \
  --url http://localhost:3000/api/transaction \
  --header 'Content-Type: application/json' \
  --data '{
	"from_account": 1,
	"to_account": 2,
	"amount": 100
}'
```
