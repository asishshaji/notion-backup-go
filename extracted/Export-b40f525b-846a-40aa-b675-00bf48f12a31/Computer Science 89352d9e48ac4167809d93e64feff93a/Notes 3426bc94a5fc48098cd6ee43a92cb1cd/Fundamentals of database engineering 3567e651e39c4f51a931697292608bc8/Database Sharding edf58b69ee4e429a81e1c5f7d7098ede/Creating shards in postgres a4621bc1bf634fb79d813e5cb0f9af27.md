# Creating shards in postgres

![Untitled](Creating%20shards%20in%20postgres%20a4621bc1bf634fb79d813e5cb0f9af27/Untitled.png)

Create an init.sql file, which is supposed to run everytime a new shard is created.

![Untitled](Creating%20shards%20in%20postgres%20a4621bc1bf634fb79d813e5cb0f9af27/Untitled%201.png)

![Untitled](Creating%20shards%20in%20postgres%20a4621bc1bf634fb79d813e5cb0f9af27/Untitled%202.png)

Build the docker image.

![Untitled](Creating%20shards%20in%20postgres%20a4621bc1bf634fb79d813e5cb0f9af27/Untitled%203.png)

![Untitled](Creating%20shards%20in%20postgres%20a4621bc1bf634fb79d813e5cb0f9af27/Untitled%204.png)

Run the shards. Now we three database instances running. Start the pgadmin instance.

```bash
docker run --name pgadmin -p 5555:80 -e PGADMIN_DEFAULT_EMAIL=asish@gmail.com -e PGADMIN_DEFAULT_PASSWORD=pass -d dpage/pgadmin4:latest
```

Add the servers to the pgadmin panel.

![Untitled](Creating%20shards%20in%20postgres%20a4621bc1bf634fb79d813e5cb0f9af27/Untitled%205.png)

```jsx
const app = require("express")();
const { Client } = require("pg");
const ConsistentHash = require("consistent-hash")
const crypto = require("crypto")

const hashRing = new ConsistentHash();
hashRing.add("5432")
hashRing.add("5433")
hashRing.add("5434")

const clients = {
    "5432": new Client({
        "host": "192.168.29.111",
        "port": "5432",
        "user": "postgres",
        "password": "postgres",
        "database": "postgres",
    }),
    "5433": new Client({
        "host": "192.168.29.111",
        "port": "5433",
        "user": "postgres",
        "password": "postgres",
        "database": "postgres",
    }),
    "5434": new Client({
        "host": "192.168.29.111",
        "port": "5434",
        "user": "postgres",
        "password": "postgres",
        "database": "postgres",
    }),
}

connect();
async function connect() {
    await clients["5432"].connect();
    await clients["5433"].connect();
    await clients["5434"].connect();
}

app.get("/", (req, res) => {

})

app.post("/", async (req, res) => {
    const url = req.query.url;

    const hash = crypto.createHash("sha256").update(url).digest("base64")
    const urlId = hash.substring(0, 5)

    const server = hashRing.get(urlId)
    await clients[server].query("INSERT INTO URL_TABLE (URL,URL_ID) VALUES($1,$2)",
        [url,
            urlId])

    res.send({
        "urlId": urlId,
        "url": url,
        "server": server,
    })

})

app.listen("8080", () => console.log("listening"))
```