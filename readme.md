## Prerequisite

* Clone the repository
   - `git clone https://github.com/Gabrielncosta/RabbitMQ-Go-And-NodeJS`

*  Set the environment variables
   - Rename the `.env.example` file to` .env`.
   - Fill in the RabbitMQ connection string in both folders

* Have docker installed on your computer




## Running the producer

```bash
# development
$ npm install / yarn dev

# run producer to start sending messages
$ node index.js
```

## Running the consumer with docker

```bash
# build the container with docker
$ docker build -t golang/consumer .

# run the consumer to start receiving messages in your terminal
$ docker run golang/consumer
```