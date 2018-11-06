# Archer
Archer - HTTP REST Orchestrator

# Example
```json
{
	"url": "https://host/teste/v1/zipcodes/{zipcode}",
	"method": "GET",
	"path": [
		{
			"key": "zipcode",
			"value":"1234567"
		}
	],
	"header": [
		{
			"key": "x-ola-teste",
			"value":"ola"
		}
	],
	"query-string": [
		{
			"key": "app-key",
			"value":"teste"
		}
	]
}
```