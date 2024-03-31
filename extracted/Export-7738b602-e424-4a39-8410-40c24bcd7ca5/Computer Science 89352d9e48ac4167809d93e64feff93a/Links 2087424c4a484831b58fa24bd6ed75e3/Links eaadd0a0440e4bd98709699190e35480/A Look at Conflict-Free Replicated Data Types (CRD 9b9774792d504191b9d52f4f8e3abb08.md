# A Look at Conflict-Free Replicated Data Types (CRDT)

URL: https://medium.com/@istanbul_techie/a-look-at-conflict-free-replicated-data-types-crdt-221a5f629e7e
Category: Design, System Design

![https://miro.medium.com/v2/resize:fit:625/1*Vn9ZGiUAcxQHw27RH3bz8g.png](https://miro.medium.com/v2/resize:fit:625/1*Vn9ZGiUAcxQHw27RH3bz8g.png)

---

Top highlight

You may have heard about CRDTs in the past few years if you are into distributed systems. In this post I will give a brief summary of what they are and what kind of guarantees they provide. In short, CRDTs are objects that can be updated without expensive synchronization/consensus and they are guaranteed to converge eventually if all concurrent updates are commutative (see below) and if all updates are executed by each replica eventually. For giving these guarantees these objects have to satisfy certain conditions, which I will briefly describe below. For more details and proofs please take a look at Marc Shapiro’s papers given in the references section below.

In their papers Shapiro et al. consider two models of replication in an eventually consistent distributed system: state-based and operation-based approach, and based on the model of replication they define two types of CRDTs, CvRDT (convergent replicated data type) and CmRDT ( commutative replicated data type). Interestingly, they show that these two replication models and these two types of CRDTs are equivalent. First let’s take a look at these two replication approaches, and then we can look at a simple CRDT example to materialize the concept.

**State-based replication:** When a replica receives an update from a client it first updates its local state, and then some time later it sends its full state to another replica. So occasionally every replica is sending its **full state** to some other replica in the system. And a replica that receives the state of another replica applies a **merge** function to merge its local state with the state it just received. Similarly this replica also occasionally sends its state to another replica, so every update eventually reaches all replicas in the system. In their paper Shapiro et al. show that, if the set of values that the state can take forms a [semi-lattice](https://en.wikipedia.org/wiki/Semilattice) (a partially ordered set with a join/least upper bound operation) and updates are increasing (e.g., say, state is an integer and update is an increment operation), and if merge function computes the least upper bound, then replicas are guaranteed to converge to the same value (which is the least upper bound of the most recent updates). And for the set of all possible system states to be a semi-lattice, this merge operation has to be idempotent, associative, and commutative. A replicated object satisfying this property (called monotonic semi-lattice property in the paper) is one type of CRDT, namely CvRDT — convergent replicated data type.

State-based approach. “s” denotes the source replica where the initial update is applied. From [2].

![A%20Look%20at%20Conflict-Free%20Replicated%20Data%20Types%20(CRD%209b9774792d504191b9d52f4f8e3abb08/1Vn9ZGiUAcxQHw27RH3bz8g.png](A%20Look%20at%20Conflict-Free%20Replicated%20Data%20Types%20(CRD%209b9774792d504191b9d52f4f8e3abb08/1Vn9ZGiUAcxQHw27RH3bz8g.png)

**Operation-based replication**: In this approach a replica doesn’t send its full state to another replica, which can be huge. Instead it just sends/broadcasts the **update** operation to **all** the other replicas in the system and expects them to replay that update (similar to [state machine replication](https://en.wikipedia.org/wiki/State_machine_replication)). Since this is a broadcast operation, if there are two updates, *u_1* and *u_2*, applied at some replica *i* and if *i* sends these updates to two other replicas *r_1* and *r_2*, these updates may arrive to these replicas in different orders, that is, *r_1* can receive them in the order *u_1* followed by *u_2*, while *r_2* can receive the updates in the order *u_2* followed by *u_1*. How do these replicas converge then? Well, they can converge if these updates are **commutative —** no matter which order these updates are applied at a replica the resulting state will be the same. In this model where the updates are broadcast to all replicas, an object for which all concurrent updates are commutative is called a CmRDT (commutative replicated data type).

Operation-based approach. “s” denotes source replicas and “d” denotes the downstream replicas. From [2].

![A%20Look%20at%20Conflict-Free%20Replicated%20Data%20Types%20(CRD%209b9774792d504191b9d52f4f8e3abb08/1E-_IN_tTSiirbJ1-7XTh5w.png](A%20Look%20at%20Conflict-Free%20Replicated%20Data%20Types%20(CRD%209b9774792d504191b9d52f4f8e3abb08/1E-_IN_tTSiirbJ1-7XTh5w.png)

The simplest CRDT example is the following integer vector. Assume that we are using the state-based replication model. To have an integer vector CRDT, we need to show that the set of integer vectors is a semi-lattice (has a partial order among its elements and a join/least upper bound operation). In fact it is, because we can (partially) order two vectors v and v’ by defining a binary relation v **<=** v’ as ∀*i* v[*i*] <= v’[*i*], that is, a vector v is less than a vector v’ if each integer in v is less than or equal to the integer in v’ at the same index (e.g., [3,6] <= [4,7]). And we also need to define a join/least upper bound operation for the merge operation, which we define as the per-index maximum operation. For example, assume that a replica has state [3,5] and it sends its state to another replica that has state [4,2] then the result of the merge operation at this replica will be [4,5]. The final condition is that the state should be monotonically increasing as a result of updates, and this holds if we define the update operation to be the increment operation for index *i.* If you think about it, since the state (in this case the integers in a vector) is just monotonically increasing, and since each replica does state merges by taking the per-index-maximum, then eventually the final value at every index will be the maximum value it has ever been updated to, so the state will eventually converge on all the replicas. This is just a simple example, it’s surprising to see that with the principles I mentioned here it’s possible to define complex CRDTs such as sets, maps, and graphs — please see the papers below for more complex examples.

CRDTs are addressing an interesting and a fundamental problem in distributed systems, but they have some important limitations which Shapiro et al. acknowledge in [2]: “Since, by design, a CRDT does not use consensus, the approach has strong limitations; nonetheless, some interesting and non-trivial CRDTs are known to exist.”. The limitation is CRDTs address only a part of the problem space as not all of the possible update operations are commutative, and so not all problems can be cast to CRDTs. On the other hand, for some types of applications CRDTs can definitely be useful as they provide a nice abstraction to implement replicated distributed systems while at the same time giving theoretical consistency guarantees.

**References:**

[1] [http://arxiv.org/abs/0907.0929](http://arxiv.org/abs/0907.0929)

[2] [http://hal.upmc.fr/inria-00555588/document](http://hal.upmc.fr/inria-00555588/document)

[3] [https://hal.inria.fr/inria-00609399v1/document](https://hal.inria.fr/inria-00609399v1/document)

[4] [http://christophermeiklejohn.com/crdt/2014/07/22/readings-in-crdts.html](http://christophermeiklejohn.com/crdt/2014/07/22/readings-in-crdts.html) (Has a good list of references, talks, etc.)

[5] [http://basho.com/posts/technical/distributed-data-types-riak-2-0/](http://basho.com/posts/technical/distributed-data-types-riak-2-0/)

**Talks/Slides:**

[1] [http://research.microsoft.com/apps/video/default.aspx?id=153540&r=1](http://research.microsoft.com/apps/video/default.aspx?id=153540&r=1)

[2] [http://research.microsoft.com/en-us/um/redmond/events/mysorepark/talkslides/marc%20shapiro%20eventual-consistency-mysore-2011-02.pdf](http://research.microsoft.com/en-us/um/redmond/events/mysorepark/talkslides/marc%20shapiro%20eventual-consistency-mysore-2011-02.pdf)

[3][http://boemund.dagstuhl.de/mat/Files/13/13081/13081.ShapiroMarc.Slides.pdf](http://boemund.dagstuhl.de/mat/Files/13/13081/13081.ShapiroMarc.Slides.pdf)

[4] [http://events.telecom-sudparis.eu/wetice/invited_speakers/ShapiroKeyNote.pdf](http://events.telecom-sudparis.eu/wetice/invited_speakers/ShapiroKeyNote.pdf)

[5] [https://2014.nosql-matters.org/cgn/wp-content/uploads/2013/02/russell-crdt-nosql-cgn-deck.pdf](https://2014.nosql-matters.org/cgn/wp-content/uploads/2013/02/russell-crdt-nosql-cgn-deck.pdf)

[6] [http://zvonimir.info/events/ec2-2013/marc-shapiro-reorder-ec2_2013.pdf](http://zvonimir.info/events/ec2-2013/marc-shapiro-reorder-ec2_2013.pdf)