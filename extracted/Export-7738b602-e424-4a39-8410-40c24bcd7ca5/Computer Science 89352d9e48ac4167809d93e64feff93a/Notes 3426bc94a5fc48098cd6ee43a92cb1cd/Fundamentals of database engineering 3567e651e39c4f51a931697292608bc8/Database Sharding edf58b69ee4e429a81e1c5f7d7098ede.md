# Database Sharding

Let’s say we have a table will 1 million pages, then your operations will slow down.

![Untitled](Database%20Sharding%20edf58b69ee4e429a81e1c5f7d7098ede/Untitled.png)

![Untitled](Database%20Sharding%20edf58b69ee4e429a81e1c5f7d7098ede/Untitled%201.png)

We can partition the data using **database partitioning.**

## Consistent Hashing

[Consistent hashing algorithm - High Scalability -](http://highscalability.com/blog/2023/2/22/consistent-hashing-algorithm.html)

![Untitled](Database%20Sharding%20edf58b69ee4e429a81e1c5f7d7098ede/Untitled%202.png)

[Consistent Hashing | The Backend Engineering Show](https://www.youtube.com/watch?v=p6wwj0ozifw)

Let’s assume that I’ve three shards. Hashfunction gives you information on which db instance to connect to.

![Untitled](Database%20Sharding%20edf58b69ee4e429a81e1c5f7d7098ede/Untitled%203.png)

![Untitled](Database%20Sharding%20edf58b69ee4e429a81e1c5f7d7098ede/Untitled%204.png)

4 is your server pool size.

The cost to figure out which instance to connect is zero.

Now if we add a new shard or remove a shard. The hashing will produce an inconsistent index for the server to access. 

![Untitled](Database%20Sharding%20edf58b69ee4e429a81e1c5f7d7098ede/Untitled%205.png)

Keys donot match nomore. That’s a huge problem. 

[Consistent Hashing and why it might not be the correct answer to your system design interview | Shing's Blog](https://shinglyu.com/web/2022/02/11/consistent-hashing-and-why-it-might-not-be-the-correct-answer-to-your-system-design-interview.html)

Consistent hashing is really a way to distribute traffic among chaning web-server population. 

![Untitled](Database%20Sharding%20edf58b69ee4e429a81e1c5f7d7098ede/Untitled%206.png)

![Untitled](Database%20Sharding%20edf58b69ee4e429a81e1c5f7d7098ede/Untitled%207.png)

[Consistent Hashing](https://medium.com/@animeshblog/consistent-hashing-d23379273ade)

[Consistent Hashing: Algorithmic Tradeoffs](https://dgryski.medium.com/consistent-hashing-algorithmic-tradeoffs-ef6b8e2fcae8)

![https://miro.medium.com/v2/resize:fit:640/format:webp/1*_n6tY98Rtd--6RGWaxrZBg.jpeg](https://miro.medium.com/v2/resize:fit:640/format:webp/1*_n6tY98Rtd--6RGWaxrZBg.jpeg)

[Creating shards in postgres](Database%20Sharding%20edf58b69ee4e429a81e1c5f7d7098ede/Creating%20shards%20in%20postgres%20a4621bc1bf634fb79d813e5cb0f9af27.md)

## When to use sharding

It’s the last thing that you want to use. Try to stick with partitioning and replication.