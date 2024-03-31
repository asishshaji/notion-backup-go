# Multiplexing vs Demultiplexing

Multiplexing is combining multiple TCP connections into one. Demultiplexing is the opposite. The issue with this is that chrome can only open at most 6 connections at a time. So this becomes a bottleneck. 

Http/1 requires multiple TCP connections for multiple requests.

But in Http/2 multiple requests can be combined into a single TCP connection. These multiple connections are essentially multiplexed into single connection. 

![Untitled](Multiplexing%20vs%20Demultiplexing%203fee22e869ea45ecb8b3ab61ade5555d/Untitled.png)

![Untitled](Multiplexing%20vs%20Demultiplexing%203fee22e869ea45ecb8b3ab61ade5555d/Untitled%201.png)

In the above figure, multiple http/1 requests are combined into a single http/2 connection by the proxy. The downside is that it puts lot of load on the http/2 server. The end server need to parse the single request and separate it back to three request.

![Untitled](Multiplexing%20vs%20Demultiplexing%203fee22e869ea45ecb8b3ab61ade5555d/Untitled%202.png)

This figure is just the opposite.

## Connection Pooling

![Untitled](Multiplexing%20vs%20Demultiplexing%203fee22e869ea45ecb8b3ab61ade5555d/Untitled%203.png)

We will have a pool of connections open to database server from backend server. When a request comes from the client we will assign the connection which is free to that request instead of creating new one.

![Untitled](Multiplexing%20vs%20Demultiplexing%203fee22e869ea45ecb8b3ab61ade5555d/Untitled%204.png)

Lets say the max open connection is 4. Now all of them are busy. The new request is blocked. When one of the database connection is free new connection is accepted by the database.  

### **What is the problem with normal connections?**

Assume we have a Postgres database and you want to run this query on it.

`select chairs from dining_table;`

To get the result, along with this query we need to provide the database details and our authentication details, usually the *host, port, database, username and password*. These are needed because before actually running the query, the program has to:

→ find the database → connect to it → authenticate our connection → use the given database → and then run our query.

Once the result is returned the connection is closed. When we want to run a different query we need to start over again.

As you can see we this is time consuming. Authentication is a very expensive process and most of the times creating connections take relatively more time compared to the actual query they run. So it becomes very inefficient for us to create a connection for every query.

So instead of closing the connection, keeping it active and running all the queries using the same connection gets rid of the creation latency. But in a service, how do we share a connection between different clients? That is where connection pooling comes in.

Connection Pooling alleviates this problem by creating a pool of connections at the start and keeping them alive (ie not closing the connections) till the end.

![Untitled](Multiplexing%20vs%20Demultiplexing%203fee22e869ea45ecb8b3ab61ade5555d/Untitled%205.png)

Whenever some part of the application wants to query the database, they borrow the connection from the pool and once they are done with the query instead of closing the connection, it is returned back to the pool.

If all connections are used up, the pool will do a blocking wait on new requests, and once a connection is free let the newer request through. To our program, this mostly looks like a longer query than normal.

### **What does this look like in code?**

Often in Go you will see the pattern of a *DB being created in the main function, or at least a function is called that creates it. This instance is then passed to each data layer entity or service for use. Usually as part of a struct for a service. This means that it acts as a singleton, each request that comes in is using the same database pool. That is an important part here, **each incoming request is using the same DB instance.**

```go
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kohofinancial/go-and-databases/services"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type app struct {
	UserService services.UserService
}

func main() {
	var dsn string
	var port string
	var setLimits bool
	flag.StringVar(&dsn, "dsn", "", "PostgreSQL DSN")
	flag.StringVar(&port, "port", "8080", "Service Port")
	flag.BoolVar(&setLimits, "limits", false, "Sets DB limits")
	flag.Parse()

	db, err := openDB(dsn, setLimits)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)

	application := app{UserService: services.NewPostgresUserService(db)}

	r := mux.NewRouter()
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/{id}", application.GetUser).Methods(http.MethodGet)
	userRouter.HandleFunc("", application.AddUser).Methods(http.MethodPost)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func openDB(dsn string, setLimits bool) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if setLimits {
		fmt.Println("setting limits")
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
```

I think you can see why it is great that Go does pooling on its own. By creating a single instance of the DB pool Go can make as many connections it deems necessary based on load.

## Browser demultiplexing

![Untitled](Multiplexing%20vs%20Demultiplexing%203fee22e869ea45ecb8b3ab61ade5555d/Untitled%206.png)

6 connections here represents the max no of connections that chrome browser can make. Up to 6 connections in HTTP/1.1, after that requests are blocked.