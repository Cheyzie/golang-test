# Golang test assignment

## Launch

copy .env.example to .env file

docker compose up

To check kafka queue:

docker exec -it broker kafka-console-consumer --bootstrap-server broker:9092 --topic feedbacks --from-beginning

## Endpoints:

- POST /api/v1/signup
	- Request
		{
			"username": str
			"email": str,
			"password": str
		}

  

- POST /api/v1/signin

	- Request
		{
			"email": str,
			"password": str
		}
	- Response
		{
			"access_token": str
		}

  

- POST /api/v1/feedbacks

	- Request
		{
			"customer_name": str,
			"email": str,
			"feedback_text": str,
			"source": str
		}
	- Response
		{
			"id": int
		}

- GET /api/v1/feedbacks/:id

- GET /api/v1/feedbacks

- GET /api/v1/feedbacks?offset=<int>&limit=<int>