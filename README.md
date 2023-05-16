# Piadocas

**Piadocas** it's a project created to help people that are learning how to
consume an API. It's a place to share jokes through HTTP requests.

Consume it in any language, doing any project that you want and if you feel
comfortable, can share with us. We can upload it to this repository, on the
"Consume Examples" folder to inspire other people.

## Summary

- [How it works](#how-it-works)
  - [How to post a joke](#how-to-post-an-joke)
  - [How to list jokes by username](#how-to-list-jokes-by-username)
  - [How to list jokes by time range](#how-to-list-jokes-by-timerange)
  - [How to list jokes by language](#how-to-list-jokes-by-language)
  - [How to delete a joke](#how-to-delete-an-joke)
- [Contribute](#contribute)
  - [Run on local using Docker and Docker compose](#run-on-local-using-docker-and-docker-compose)
  - [TODO](#todo)

## How it works

In **Piadocas** you can do three actions, create a joke, list a user's jokes and
delete one joke. Those three actions use the same path `/joke/`, but it depends
on the HTTP verb and the parameters passed in URL or the request body.

### How to post a Joke

To create a Joke we use the `POST` method at route `/joke/` passing through the
request body the fields:

```
{
   "call": "Joke body",
   "finish": "Punch line",
   "language": "en",
   "owner": "Your github username"
}
```

If it's succeed should return 201 (Created) status code and an JSON containing
the joke data in our database. Example with curl:

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

To list all jokes created by a specific user we have to use the `GET` method at
route `/joke/user/{username}` where `{username}` field must be replaced by the
owner(GitHub username).

If it's succeed should return 200 (Created) status code and an JSON containing
an array with all the jokes in our database where the owner is equal the
`{username}`. Example with curl:

```
curl --request GET \
  --url http://localhost:8080/joke/user/kalogs-c
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
        "language": "en",
		"created_at": "2022-09-10T18:33:28Z"
	}
]
```

Check that the field `{username}` on the URL is the same at `owner`, in this
example `"kalogs-c"`.

### How to list jokes by time range

# TODO

### How to list jokes by language

# TODO

### How to delete a joke

curl --request DELETE\
--url http://localhost:8080/joke/1

This request should return:

```
HTTP Status 200 (OK)
"Deleted sucessfully"
```

## Contribute

### Run on local using Docker and Docker compose

Just runs the command `docker-compose up -d` to build and run the following
containers:

    - PgAdmin binded on port 5050
    - Postgres binded on port 5432
    - Piadocas Api binded on port 8080

To actually run the API we will need to enter into our running container called
`app`and run the Go code, using the following commands:

```
docker exec -it app sh
go run main.go --dev
```

### TODO

- Get random joke
