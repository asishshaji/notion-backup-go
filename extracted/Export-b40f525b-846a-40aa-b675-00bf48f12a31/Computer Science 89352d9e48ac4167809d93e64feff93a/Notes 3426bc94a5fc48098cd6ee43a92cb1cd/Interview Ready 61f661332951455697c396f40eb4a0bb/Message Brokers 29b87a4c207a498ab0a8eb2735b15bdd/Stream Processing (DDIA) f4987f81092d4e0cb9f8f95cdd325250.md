# Stream Processing (DDIA)

> notes from DDIA
> 

![Untitled](Stream%20Processing%20(DDIA)%20f4987f81092d4e0cb9f8f95cdd325250/Untitled.png)

- Related events in streaming systems are usually grouped into a *topic* or *stream*.
- You can use a file or a database as a broker.
- The producer's purpose is to write to some datastore, and the consumers periodically poll the datastore to check for events.
- If we use that for a high-traffic system, the throughput won’t be that good. Because a file or a database is not designed for that purpose(eg if the datastore is not designed for large polling requests).
- 

- The input for stream processing is an event.
- Specifies the details of something that happened at some point in time.
- It contains a timestamp indicating when it happened.
- event can be text, JSON, or binary.