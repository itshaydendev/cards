![Docker](https://github.com/itshaydendev/cards/workflows/Docker/badge.svg) ![Go](https://github.com/itshaydendev/cards/workflows/Go/badge.svg)

# Cards

🃏 Cards Against Humanity, implemented using Go and React.

## Running

Cards is very easy to install. The recommended way is
[with Docker](#with-docker), but you can also run
[by building](#from-source).

### With Docker

```shell
$ docker-compose up -d db
$ docker-compose up migrate # Runs DB migrations in a quick utility container
```

### From Source
```shell
$ make
$ build/cards
```