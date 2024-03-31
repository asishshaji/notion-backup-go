# OSI Model (Open Systems Interconnection)

[OSI Model – Practical Networking .net](https://www.practicalnetworking.net/series/packet-traveling/osi-model/)

Why do we need a communication model?

Without a standard model, your application must have knowledge of the underlying network medium. In short both your client and server have to essentially speak the same language. 

Imagine building different versions of your application so that it works on wifi vs ethernet vs LTE vs fibre. That will be a disaster. 

So we needed a standard that works globally, so thus OSI model was created. OSI is everywhere. 

Without a standard model, different devices may not be able to talk to each other.

Innovations can be done in each layer separately without affecting the rest of the models. 

7 layers each describe a specific networking component.

![Untitled](OSI%20Model%20(Open%20Systems%20Interconnection)%20508ed37a5806493d81e4114eb52fc76d/Untitled.png)

Backend developers live in layers 7 and 4. 

## OSI Layers.

### POV Sender

- L7 - Application
    - POST request with JSON data to HTTPS server
- L6 - Presentation
    - Serialize JSON to flat byte strings (for HTTP server, converted to plain string)
- L5 - Session
    - Request to establish TCP connection
- L4 - Transport
    - Sends SYN request to target port 443
- L3 - Network
    - SYN is placed on an IP packet(s) and adds the source and dest IPs
- L2 - Datalink
    - Each packet goes into a single frame and adds the source and dest MAC
- L1 - Physical
    - Each frame becomes strings of bits which is converted into either a radio signal (wifi), electric signal (ethernet) or light (fibre)

 

### POV Receiver

- L1 - Physical
    - Radio, electric or light is received and converted into digital bits
- L2 - Datalink
    - The bits from L1 are assembled into frames
- L3 - Network
    - The frames are assembled into IP packets
- L4 - Transport
    - IP packets are assembled into TCP segments
- L5 - Session
    - A connection session is established or identified
    - We only arrive at this layer when necessary (when three-way handshake is done)
- L6 - Presentation
    - Deserialise flat byte strings back to JSON for the app to consume
- L7 - Application
    - The application understands the JSON request and your server request receive event is triggered.
    
    The below pictures represent the flow of the request from client to server. 
    

![Untitled](OSI%20Model%20(Open%20Systems%20Interconnection)%20508ed37a5806493d81e4114eb52fc76d/Untitled%201.png)

![Untitled](OSI%20Model%20(Open%20Systems%20Interconnection)%20508ed37a5806493d81e4114eb52fc76d/Untitled%202.png)

## What are the shortcomings

- OSI model has too many layers which can be hard to comprehend.
- Hard to argue about which layer does what
- Simpler to deal with Layers 5-6-7 as just one layer
- TCP/IP model does just that.

![Untitled](OSI%20Model%20(Open%20Systems%20Interconnection)%20508ed37a5806493d81e4114eb52fc76d/Untitled%203.png)

## How the OSI model works on Youtube?

[how the OSI model works on YouTube (Application and Transport Layers) // FREE CCNA // EP 5](https://www.youtube.com/watch?v=oIRkXulqJA4)

[REAL LIFE example!! (TCP/IP and OSI layers) // FREE CCNA // EP 4](https://www.youtube.com/watch?v=3kfO61Mensg)