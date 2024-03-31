# MapReduce: Simplified Data Processing on Large Clusters

URL: http://nil.lcs.mit.edu/6.824/2020/papers/mapreduce.pdf

- Programming model for processing and generating large datasets.

![Untitled](MapReduce%20Simplified%20Data%20Processing%20on%20Large%20Clus%20c65edf4dd22c4d0d8fb5ccf599a0ee45/Untitled.png)

- Takes key value as input, the map creates an intermediate output, the reducer merges all intermediate values associated with the same intermediate key.
- A typical map-reduce can handle terabytes of data

## why map-reduce

- Google engineers needed to process large amount of data, on various machines (distributed)
- They wanted to process the data to compute various kinds of derived data.
- eg: For indexing the whole internet
- the input data are usually very large, and computations have to be distributed across hundreds or thousands of machines.
- They thus created a map-reduce framework to address this problem, while hiding the messy details of parallelization, fault tolerance, and load balancing for the programmer.

```jsx
map(String key, String value):
// key: document name
// value: document contents
for each word w in value:
  EmitIntermediate(w, "1");
  
reduce(String key, Iterator values):
// key: a word
// values: a list of counts
	int result = 0;
	for each v in values:
		result += ParseInt(v);
	Emit(AsString(result));
```

## Examples

- Distributed Grep
    - The map function emits the line that matches the pattern from the input document
    - The reduce function groups together all the intermediary data to the output
- Count of URL Access Frequency
    - map function takes the log file as input, looks up the <url> in the file, and emits <url, 1> as an event
    - The reduce function adds together all the values for the same URL and emits a <url, total_count> event

# Execution Overview

- Reverse Web-Link Graph
    - map function outputs <target, source> pairs for each link to a target URL found in the page named source.
    - the reduce function concatenates the list of source URLS associated with a given target URL and emits <target, list(sources)>
    
    ![Untitled](MapReduce%20Simplified%20Data%20Processing%20on%20Large%20Clus%20c65edf4dd22c4d0d8fb5ccf599a0ee45/Untitled%201.png)
    
    Map invocations are distributed across multiple machines by automatically partitioning input data into a set of M splits. These input splits can be processed by multiple machines in parallel. 
    
    Reduce invocations are distributed into workers by partitioning the intermediate keys into R pieces using a partitioning function ( eg hash(key) mod R) where R is the number of partitions.
    
       
    
    # FLOW
    
    ![Untitled](MapReduce%20Simplified%20Data%20Processing%20on%20Large%20Clus%20c65edf4dd22c4d0d8fb5ccf599a0ee45/Untitled%202.png)
    
    - The mapreduce splits the input data into M segments varying from 16 Mb to 64 Mb(configurable by the user). It then starts up many copies of the program on a cluster of machines.
    - One copy is the master, the rest are workers that are assigned work by the master. There are M map tasks and R reducer tasks to assign. The master picks idle workers and assigns each one a map task or a reduce task.
    - A worker who is assigned a map task reads the contents of the input split data. It parses the k/v of input data and passes it to the map function, which then generates the intermediate output, which is then buffered in the memory.
    - Periodically, the buffered pairs are written into the local disk, partitioned into R partition by a hashing function, and the locations are sent to the master, who is responsible for forwarding these locations to the reduce workers.
    - When the reduce workers are notified about the locations of the partition by the master, the reducer then uses RPC calls to access the data from the local disk of map workers. When a reducer worker has read in all the intermediate data, it then sorts the data so that all occurrences of a key are grouped. If the intermediate data is too large to fit in the memory, external sort is used.
    - The reducer worker then iterates over the sorted intermediate data and for each unique intermediate data encountered, it passes the key and the corresponding set of intermediate values to the user’s reduce function. The output of the reduce function is then appended to the final output of the reducer’s partition.
    
    ---
    
    - Master tracks of various data structures.
    - It tracks the status of each map and reduce work, and the identity of idle worker machines
    - The master is the conduit through which the intermediate output from the map reaches the reducer.
    - The master stores the locations and sizes of the R intermediate regions produced by a map task.
    - Updates to these locations are received once the map task is complete.
    
    ### Worker Failure
    
    - The master pings every worker periodically, if no response is received from the worker in a certain amount of time, the worker is marked as failed.
    - Any map tasks completed by the worker are reset back to their initial state(idle), which makes it possible for the master to reschedule again on other workers.
        - Why mark completed tasks to the initial state?
            - The completed tasks are in the failed machine memory and therefore are inaccessible.
            - Completed reduces tasks need not be recalculated again, because they are stored in a global file system.
    - Any map tasks or reduce tasks in progress on a failed worker are marked as idle and become eligible for rescheduling.
    - Rescheduling and re-execution avoid most of the bottleneck in case of a worker failure.
    
    ### Master Failure
    
    - Better make the master periodically commit its local data structures.
    - If the master dies, a new copy can be started from the last check-pointed state.