# Message Brokers

[Stream Processing (DDIA)](Message%20Brokers%2029b87a4c207a498ab0a8eb2735b15bdd/Stream%20Processing%20(DDIA)%20f4987f81092d4e0cb9f8f95cdd325250.md)

[What is a Message Queue and When should you use Messaging Queue Systems Like RabbitMQ and Kafka](https://www.youtube.com/watch?v=W4_aGb_MOls)

What is the difference between a queue and a pub/sub? 

- Once we dequeue the item from the queue it’s gone
- For a pub/sub, we can still find the item after dequeuing, by maintaining a pointer
    - You can optionally go back and forward

> At the core every message broker is either log or queue based.
> 
> 
> Queue based: RabbitMQ, ActiveMQ, MSMQ, AWS SQS, JMQ
> 
> Log based: Apache Kakfa, Apache Pulsar, AWS Kinesis, Azure Event Hubs
> 

## Queues are transient and logs are persistent

[Message Queues: RabbitMQ vs. Kafka](https://medium.com/@ramekris/message-queues-rabbitmq-vs-kafka-335076a5ff88)

- The log-based approach works best where each message is fast to process and message ordering is important
- Queue-based approach can be used when processing a message takes a long time and message ordering is not a factor.

![Untitled](Message%20Brokers%2029b87a4c207a498ab0a8eb2735b15bdd/Untitled.png)

Your web server should only spend time responding to the request, it shouldn’t be spending time processing your time-consuming tasks(CPU-Intensive). USE QUEUE

![Untitled](Message%20Brokers%2029b87a4c207a498ab0a8eb2735b15bdd/Untitled%201.png)

[Doordash moves their Backend to Apache Kafka from RabbitMQ, VERY interesting! Let us discuss it!](https://www.youtube.com/watch?v=sXjWTLMGmVY)

---

[Event-Driven Architectures - The Queue vs The Log — Jack Vanlightly](https://jack-vanlightly.com/blog/2018/5/20/event-driven-architectures-the-queue-vs-the-log)

[Kafka](Message%20Brokers%2029b87a4c207a498ab0a8eb2735b15bdd/Kafka%2060efa40da813472bafbd67e8f75131bf.md)