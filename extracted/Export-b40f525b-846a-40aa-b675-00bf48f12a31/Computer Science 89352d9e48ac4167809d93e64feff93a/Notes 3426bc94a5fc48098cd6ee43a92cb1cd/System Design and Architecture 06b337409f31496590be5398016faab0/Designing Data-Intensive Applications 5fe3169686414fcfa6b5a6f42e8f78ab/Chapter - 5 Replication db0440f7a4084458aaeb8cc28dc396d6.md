# Chapter - 5 : Replication

Replication means keeping a copy of the same data on multiple machines connected via a network. Why on multiple machines?

- To keep data geographically close to your users (reduces latency)
- To keep the system continue working, even if some parts fail
- To enable multiple read replicas for handling high read requirements(increase read throughput)

If the data to be replicated is very small you only need to copy the data once across all the nodes.

![Untitled](Chapter%20-%205%20Replication%20db0440f7a4084458aaeb8cc28dc396d6/Untitled.png)

All database writes must be processed by each replica to maintain uniform data. The leader writes data locally and sends it to its followers via a replication log. Each follower updates its database in the same order as the leader. Reads can be handled by any node, but writes are only accepted by the leader in this `master-slave` replication. In **Synchronous Replication**, the leader waits until the writes are committed to the followers before responding to the user.

![Leader-based replication with one synchronous and one asynchronous follower.](Chapter%20-%205%20Replication%20db0440f7a4084458aaeb8cc28dc396d6/Untitled%201.png)

Leader-based replication with one synchronous and one asynchronous follower.

**Asynchronous Replication**: The leader sends the message without waiting for the follower's response.

Synchronous replication ensures the follower has the latest data, allowing its promotion to leader if needed. However, if the synchronous follower doesn't respond, all writes must pause until it's available.

Typically, only one of the followers is synchronous, while others are asynchronous. If the synchronous follower is unavailable, an asynchronous one replaces it. This is termed semi-synchronous.

[](https://levelup.gitconnected.com/mastering-database-replication-an-essential-guide-for-2023-9fa6deb3efe4)

## Setting up new followers

- You will need to increase the number of nodes at some point in time (for increasing number of replicas or replacing failed nodes)
- Copying the entire data from a node to the new node is not always sufficient
- Clients are writing to the databases continuously, so a standard file copy would see different parts of database at different points in time.

### Process Flow

- Take a snapshot of the leader’s database at some point in time
- Copy the snapshot to the new follower node
- The follower connects to the leader node and requests for all the data changes that happened after the snapshot was created
- When the follower has processed the backlog of the data changes since the snapshot, the node is said to be ready.

# Handling Node outages

Nodes can go down due to failures, s/w updates etc. The goal is to keep the system as a whole running even during node outages.

### Follower Failure: Catch-up recovery

- Followers keep a log of data changes received from the master
- If a follower crashes, the follower can recover from its log, it knowns the last log it processed.
- And the follower can then request the master for all the data changes after the last processed one.
- When those changes are applied by the follower, the follower is caught up with the master and can receive stream of data changes as before.

### Leader failure: failover