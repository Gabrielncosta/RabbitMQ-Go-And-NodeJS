require('dotenv').config()
const amqp = require('amqp-connection-manager');

(async () => {
  try {
    const connection = await amqp.connect([process.env.CONNECTION_STRING])

    const queueName = 'test-queue';
  
    const channel = await connection.createChannel({
      json: true,
      setup: function(channel) {
        return channel.assertQueue(queueName, { durable: true })
      }
    });

    setInterval(async () => {
      await channel.sendToQueue(queueName, { hello: 'world '});

      console.log('message sent');
    }, 5000)
  } catch (error) {
    console.log(error)
  }
})();

// const connection = amqp.connect()