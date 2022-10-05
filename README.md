# Futbol Matches

## Running the project

**Requirements:** golang 1.18 docker 20.10.12

**Start (build included):**

start components
```shell
    make start
```

stop all components

```shell
    make stop
```
**Dev (build included):**

Run local instance

```shell
    make dev
```

**Run tests:**

Using the native go command:

```shell
    make test
```

**Test coverage:**

The test runner includes the package code coverage. To get a detailed code report, run the following runner to generate an html view of it 

```shell
    make coverage
```

## Live API

**Backend:**

The api can be tested on the following domain

```
    https://api.futbol.fast.ar
```

**Bot:**

To talk to the bot simply add it to telegram and send the next message:
```
    /comenzar
```
<p align="middle">
<img src="https://i.ibb.co/PFbmmFq/telegram-bot-qr-3.jpg" alt="Futbol Bot QR" width="200"/>
</p>