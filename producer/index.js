require('dotenv/config');
const amqp = require('amqp-connection-manager')


(async () => {
  const connection = await amqp.connect(process.env.CONNECTION_STRING)
})()

const connection = amqp.connect()