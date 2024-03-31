# How We Scaled to 100 Million Active Users Using Kafka and Golang — Eventual Consistency

URL: https://itnext.io/how-we-scaled-to-100-million-active-users-using-kafka-and-golang-eventual-consistency-6241cfeba7e8
Category: Design, Golang, System Design

![https://miro.medium.com/v2/resize:fit:858/1*EKg0Oo4woZeIdrX2vexPeA.png](https://miro.medium.com/v2/resize:fit:858/1*EKg0Oo4woZeIdrX2vexPeA.png)

---

Nowadays, we have reached an era where the most popular startups reach millions of users within less than a year. During my experience as a software developer, where I had the privilege to work on a couple of them, I’ve seen that the most common bottleneck within a backend service is caused by I/O overhead. In this article, we will discuss the **Eventual Consistency** technique and how we can overcome I/O bottlenecks in scale by utilizing Kafka.

# How simplicity allows Kafka to scale up to millions of messages per second

Basically, a Kafka cluster consists of multiple servers, called brokers in the Kafka vocabulary, and one or more topics.

![How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1gQYItzs6sjU95oCHhnPwxg.png](How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1gQYItzs6sjU95oCHhnPwxg.png)

**What is a topic in Kafka?** A topic is a stream of messages with similar behaviors. They can be similar events, JSON, or anything as bytes. For instance, we had topics for published posts, comments, and likes, as well as action logs to study user experiences.

**How is the structure of a topic in a cluster?** Topics have two main properties: the number of partitions and replications. The replication determines how many copies we need from a topic to increase resiliency in case of broker failures. Before describing partitions, let’s discuss a message's journey, from its publication to a topic until received and committed by a consumer.

![How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1cV_2kAAcjjPjcQzJi_cnsw.png](How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1cV_2kAAcjjPjcQzJi_cnsw.png)

First, we have a happy message published by a producer going toward a topic. What will happen when the topic receives the message? A topic is not just one queue. It consists of one or more queues that we call **partitions**. These partitions can be located in one or more different brokers.

![How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1AC-CNlH_aZiRm8tN21W1Qw.png](How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1AC-CNlH_aZiRm8tN21W1Qw.png)

When the topic receives the message, the topic leader decides where to put the message. We can configure this behavior by using keys and distribution strategies. However, you can allow Kafka to distribute your messages fairly and randomly among partitions.

![How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1EKg0Oo4woZeIdrX2vexPeA.png](How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1EKg0Oo4woZeIdrX2vexPeA.png)

Each consumer can read from one or more partitions. The key point is that no more than one consumer can read from a partition. **What does it mean?**

![How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/18huhjb-Wmq2sO2lzGxqiwA.png](How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/18huhjb-Wmq2sO2lzGxqiwA.png)

If we have more consumers than partitions, excess consumers will starve! At first, I thought it was going to limit us a lot. But this behavior allows Kafka to scale out partitions dependently and achieve millions of messages per second.

# Eventual Consistency — In eventuality, we trust!

Let’s explore different scenarios. Posting a new Instagram photo, a new tweet, a new comment on Amazon, liking a photo, uploading a new avatar, or looking at a photo! In the traditional way, when a client sends an HTTP request to upload a new photo, the backend holds that request until the request is either successful or failed. But there is another approach to it that suits well with asynchronous programming.

**How much should It take until our followers receive a new tweet?** 100 milliseconds, one second, one minute, or five minutes? As Twitter has mentioned in their articles, they are happy with less than five seconds. In our case, in request peaks with way fewer resources than Twitter, we were happy with less than a minute. The key point is that in many scenarios, we are okay with seconds of delays, but we do not need to hold the user until the request is finished. Imagine waiting 5 seconds to tweet each time.

In other scenarios, like gathering action logs to compute user behaviors later, we do not need messages instantly. Hence, making users wait until the application saves their log is not logical at all. Besides that, Kafka has powerful integrations with other tools, such as ELK and Flink, to work on streams of messages.

The second reason I should mention that makes an application slow is I/O bottlenecks. Assume we need to build a blogging service. By separating reads from writes, we can scale databases independently. For instance, we can have a relational database that supports transactions for insert queries, and after the data have been saved, it publishes a new event that synchronizes read databases. Read databases, therefore, don’t need to be as complicated as write databases and can be optimized just for read queries.

One of the most useful techniques that we used to handle around 10 million requests per hour was utilizing Redis as the read database. I was working on services that directly respond to GET requests for editorial pages. We processed every page and then saved them in the Redis. And whenever there was a change within a page, a Kafka event was triggered to update the cache. Therefore, the Redis was always updated, and we were able to answer requests as fast as possible.

![How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1UC_ZwL4T24kMLwxDs75K0w.png](How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1UC_ZwL4T24kMLwxDs75K0w.png)

Why do we have two different topics? Because in many cases, your team is only in charge of one part of the page, and other teams are updating other entities, and your team doesn’t need to be worried about it. Using a message broker provides decoupling, less complexity in services, and resilience and scalability by eventual consistency.

# Let’s Build a Blogging Service

Let’s build a sample service that utilizes Kafka and Redis to implement eventual consistency. While setting up Kafka and Redis in the development environment is not that challenging, having a resilient and scalable Kafka and Redis cluster on the cloud requires years of experience.

> We will be using Upstash’s free Kafka and Redis clusters in this article. Their free-tier Kafka and Redis clusters support up to 10k messages per day which is more than enough for development environment.
> 

To create your cluster, **[click here](https://upstash.com/?utm_source=Mohammad1)** and create a new Upstash account. Go to the Kafka tab and click on the create cluster button.

![How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1jLCZ7LYyHvObsCaaBJGwZg.png](How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1jLCZ7LYyHvObsCaaBJGwZg.png)

As you can see, you can select between multiple regions and choose the replication factor. We are going to use this topic for the testing environment, so I’ll leave it on Single Replica.

![How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1cwSzx0e1VyLCI4eqE-Hv-A.png](How%20We%20Scaled%20to%20100%20Million%20Active%20Users%20Using%20Ka%207e4f71f3951941cc98f7d02d079b2dc2/1cwSzx0e1VyLCI4eqE-Hv-A.png)