# Why do we need certificates

What are the available solutions that we have for encrypting the communication between client and server??

- Symmetric
- Asymmetric (Public-key encryption)

[What is asymmetric encryption? | Asymmetric vs. symmetric encryption | Cloudflare](https://www.cloudflare.com/learning/ssl/what-is-asymmetric-encryption/)

[Certificates and Certificate Authority Explained](https://www.youtube.com/watch?v=x_I6Qc35PuQ)

## Solution 1: Use symmetric encryption

![Untitled](Why%20do%20we%20need%20certificates%20aed17af78ef047d481f3584ac85f4352/Untitled.png)

This is very simple, the client and server essentially shares a key. Symmetric cryptography is much faster than asymmetric. If someone manages to get a hold of the shared key, it’s game over.

## Solution 2: Using asymmetric encryption

![Untitled](Why%20do%20we%20need%20certificates%20aed17af78ef047d481f3584ac85f4352/Untitled%201.png)

The server will have a public and private key. When the client asks for a key, the server sends its public key. The client can now use the public key to encrypt the messages, and only the server will be able to decrypt the message since the private key is available only with the server. 

This solution is good, but there are some downsides to it.

- During every request/response, the server/client [will need to do an asymmetric encryption/decryption.](https://crypto.stackexchange.com/questions/30777/exactly-how-slow-is-asymmetric-encryption)
    - We could combine symmetric key property with this.
    - Once you have the public key, the client can generate a random secret and sign it to the server
    - Voila, you have a shared secret key with the client and server.
- Using an MITM attack, a user in the middle can intercept the request and forward its public key to the client, the client has no way of making sure that the public key belongs to the server.

![Untitled](Why%20do%20we%20need%20certificates%20aed17af78ef047d481f3584ac85f4352/Untitled%202.png)

The above solution is a good one, but we need some kind of mechanism to verify the authenticity of the public key. That’s where certificates come into the picture.

![Untitled](Why%20do%20we%20need%20certificates%20aed17af78ef047d481f3584ac85f4352/Untitled%203.png)

So to solve the problem of asymmetric encryption, we introduce a certificate that verifies the public key belongs to the server itself.

So before all these, when the server is published. The server asks some certificate authority to give a certificate for it. The server sends its public key to the CA. The CA will generate a certificate for the server. This is the certificate which is sent to all the users during the connection.

The certificate contains information about the server, the validity of the certificate and details of the authority. It has the public key of the server and a signature. The CA signs the public key of the server with its private key to get the signature. Now the clients who get the certificate can use the public key of the CA to decrypt the signature and get a value, if the certificate wasn’t tampered then the public key in the certificate and the value will match. This can verify the authenticity of the public key.

This is ***TLS1.2***

[Kazakhstan Begins Intercepting HTTPS Internet Traffic Of All Citizens Forcefully](https://thehackernews.com/2019/07/kazakhstan-https-security-certificate.html)