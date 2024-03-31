# Single Listener, Single Worker thread

eg: Node

One process does all the jobs for us.A simple, elegant and easy to understand architecture. Where the backend application consists of a single thread that acts as a listener, so it binds on an IP:Port, it accepts connections on the socket, and it also reads the TCP stream from all connections it accepts (It uses epolling to do this asynchronously).  [NodeJS](https://www.youtube.com/watch?v=gMtchRodC2I) for instance uses this model and as a result any application you build using NodeJS.

![Untitled](Single%20Listener,%20Single%20Worker%20thread%208f282194774d4027a418d7c83a2890a9/Untitled.png)