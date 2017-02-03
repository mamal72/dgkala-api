# dgkala-api
A dead simple HTTP API for [dgkala](https://github.com/mamal72/dgkala).


## Installation

```bash
go get github.com/mamal72/dgkala-api
```


## Usage

```bash
dgkala-api
```

*You can also set ENV variables or run the app like this to change host or port:*

```bash
HOST="localhost" PORT="5000" dgkala-api
```

Now you can send simple `GET` http requests to following addresses:

- `http://HOST:PORT/search/:keyword`
- `http://HOST:PORT/incredible-offers`.


## Development

1. Clone the repository.

1. Edit the example config ENV config in `.env.example` file and add `PORT` and `HOST`.

1. Rename `.env.example` to `.env`.

1. Build or run the app.


## Ideas || Issues
Just create an issue and describe it. I'll check it ASAP!


## Contribution

You can fork the repository, improve or fix some part of it and then send the pull requests back if you want to see them here. I really appreciate that. :heart:


## License

Licensed under the [MIT License](https://github.com/mamal72/dgkala-api/blob/master/LICENSE).