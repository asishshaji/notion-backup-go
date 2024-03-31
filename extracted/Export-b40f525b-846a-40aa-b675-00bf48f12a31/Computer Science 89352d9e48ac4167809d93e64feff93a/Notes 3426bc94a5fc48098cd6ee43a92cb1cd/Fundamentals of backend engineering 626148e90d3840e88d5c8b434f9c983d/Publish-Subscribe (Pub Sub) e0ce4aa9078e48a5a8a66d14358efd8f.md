# Publish-Subscribe (Pub/Sub)

Pub/Sub is an asynchronous and scalable messaging service that decouples services producing messages from services processing those messages.

![https://miro.medium.com/v2/resize:fit:750/0*2pjoU9OlhVGanYpy.gif](https://miro.medium.com/v2/resize:fit:750/0*2pjoU9OlhVGanYpy.gif)

Basically, the Pub/Sub model involves:

1. A publisher sends a message.
2. A subscriber receives the message via a message broker.

Publisher/subscriber messaging, or pub/sub messaging, is a form of asynchronous service-to-service communication used in serverless and microservices architectures. 

In a pub/sub model, any message published to a topic is immediately received by all of the subscribers to the topic.

![https://miro.medium.com/v2/resize:fit:1100/format:webp/1*n3WAkUW-UtI9_pRSsoUcmA.png](https://miro.medium.com/v2/resize:fit:1100/format:webp/1*n3WAkUW-UtI9_pRSsoUcmA.png)

# **Problems with Request-Response Architecture**

Request-Response architecture is the most popular and common messaging system employed when making network calls. REST is one of the most common forms of messaging protocols used in distributed systems. The client makes a *Request* which is serviced by the server to provide a *Response*.

The problem arises as the distributed architecture gets complex. The introduction of multiple services creates a chain of strongly connected components. This increases the latency and waits time, and failure in one component can lead to cascading failures.

For example, let’s say we have a video uploader service for Instagram reels. You have an upload service, a compressor service, a formatting service and a notification service, all of which make network calls using Request-Response architecture. The user uploads a video which is handled by the uploader service. The uploader service then triggers a compressor service that’s responsible for compressing the raw video. The formatting service then formats the video using the filters, music etc. selected by the user. The notification service then provides a notification to the end user that the upload is successful.

![https://miro.medium.com/v2/resize:fit:1100/format:webp/1*aO5VQYvzqv8IolXgiSPfhQ.png](https://miro.medium.com/v2/resize:fit:1100/format:webp/1*aO5VQYvzqv8IolXgiSPfhQ.png)

If these systems interact using request-response messaging, it would lead to a complex graph of connected components as the system evolves. The user has to wait till each service processes the request and provides a response. This would increase latency. The system would also not be fault tolerant as a failure in any one of the services would lead to failure in the entire chain.

# **Event-Driven Architecture**

Event-Driven Architecture was designed to solve the above problem. We could make the system loosely coupled by using events to deal with the exchange of information. An **Event** is a change in state, or an update, like an item being placed in a shopping cart on an e-commerce website. Events can either carry the state (the item purchased, its price, and a delivery address) or events can be identifiers (a notification that an order was shipped).

Event-driven architectures have three key components: event producers, routers, and consumers. A producer publishes an event to the message router, which filters and pushes the events to consumers. Producer services and consumer services are decoupled, which allows them to be scaled, updated, and deployed independently.

![https://miro.medium.com/v2/resize:fit:1155/1*i1ZvUBq8mG4TW_C0VchfRA.png](https://miro.medium.com/v2/resize:fit:1155/1*i1ZvUBq8mG4TW_C0VchfRA.png)

Using this architecture we can scale each component individually and make them fault tolerant. Since the messages are async we don’t have to wait for response from downstream services.

A Message Router, or a Message Broker(also called Topics in Kafka), is software that enables applications, systems, and services to communicate with each other and exchange information. The message broker does this by translating messages between formal messaging protocols. This allows interdependent services to “talk” with one another directly, even if they were written in different languages or implemented on different platforms.

# **Publisher-Subscriber Messaging**

Publisher Subscriber Messaging, also known as the Publisher Subscriber pattern, is a messaging pattern that provides a framework for exchanging messages. It also allows for loose coupling and scaling between the sender of messages (publishers) and receivers (subscribers) on messaging brokers they subscribe to. Some common technologies using Pub/Sub Messaging are RabbitMQ, Kafka, Redis etc.

Messages are sent from a publisher to subscribers as they become available. Publishers are services that send messages to the message broker. The subscribers then subscribe to message brokers they’re interested in to listen to these messages (similar to how a user subscribes to a channel on YouTube).

![https://miro.medium.com/v2/resize:fit:1155/1*3PuN-NsIk6b_HlEy8bAkkA.png](https://miro.medium.com/v2/resize:fit:1155/1*3PuN-NsIk6b_HlEy8bAkkA.png)

Pub Subsystems are widely used in use cases where we need loosely coupled/decoupled components that are needed to be scaled individually. eg: Event Notification, Distributed Caching, Distributed Logging etc.