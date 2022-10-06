# Piadocas

**Piadocas** it's a project created to help people that are learning how to consume an API. It's a place to share jokes through HTTP requests.

Consume it in any language, doing any project that you want and if you feel confortable, can share with us. We can upload it to this repository, on the "Consume Examples" folder to inspire other people.

## Summary

- [How it works](#how-it-works)
  - [How to post an joke](#how-to-post-an-joke)
  - [How to list jokes by username](#how-to-list-jokes-by-username)
  - [How to delete an joke](#how-to-delete-an-joke)
- [Contribute](#contribute)
  - [Run on local using Docker and Docker compose](#run-on-local-using-docker-and-docker-compose)

## How it works

In **Piadocas** you can do three actions, create an joke, list an user's jokes and delete one joke. Those three actions use the same path `/joke/` but it depends of the HTTP verb and the parameters passed in URL or the request body.

### How to post an Joke

To create an Joke we use the `POST` method at route `/joke/` passing through the request body the fields:

```
 {
    "call": "Joke body",
    "finish": "Punch line",
    "owner": "Your github username"
 }
```

If it's succeed should return 201 (Created) status code and an JSON containing the joke data in our database. Example with curl:

```
curl --request POST \
  --url http://localhost:8080/joke \
  --header 'Content-Type: application/json' \
  --data '{
	"call": "That is my test joke",
	"finish": "Very funny, right?",
	"owner": "kalogs-c"
}'
```

This request should return:

```
HTTP Status 201 (Created)
{
	"id": 1,
	"call": "That is my test joke",
	"finish": "Very funny, right?",
	"owner": "kalogs-c",
	"created_at": "2022-09-10T18:33:28.127455634Z"
}
```

### How to list jokes by username

To list all jokes created by an specific user we have to use the `GET` method at route `/joke/{username}` where `{username}` field must be replaced by the owner(github username).

If it's succeed should return 200 (Created) status code and an JSON containing an array with all the jokes in our database where the owner is equal the `{username}`. Example with curl:

```
curl --request GET \
  --url http://localhost:8080/joke/kalogs-c
```

This request should return:

```
HTTP Status 200 (OK)
[
	{
		"id": 1,
		"call": "That is my test joke",
		"finish": "Very funny, right?",
		"owner": "kalogs-c",
		"created_at": "2022-09-10T18:33:28Z"
	}
]
```

Check that the field `{username}` on the url is the same at `owner`, in this example `"kalogs-c"`.

### How to delete an joke

curl --request DELETE \
  --url http://localhost:8080/joke/1

This request should return:
```
HTTP Status 200 (OK)
"Deleted sucessfully"
```

## Contribute

### Run on local using Docker and Docker compose

Just runs the command `docker-compose up -d` to build and run the following containers:

	- PhpMyAdmin binded on port 9090
	- MySql binded on port 3306
	- Piadocas Api

To actually run the API we will need to enter into our running container called `app`and run the golang code, using the following commands: 
```
docker exec -it app sh
go run main.go
```