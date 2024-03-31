# SSL

SSL allows web clients and servers to 

- Verify each other’s identity
- Encrypt messages they send to each other
- Ensure the integrity of messages sent between them.

Uses a combination of symmetric and asymmetric encryption. 

Symmetric encryption uses the same key to encrypt and decrypt. eg: **DES, AES**

Asymmetric encryption (public key encryption) uses a public and private key. The public key is used for encryption and the private key is used for decryption. eg: RSA

Asymmetric encryption can also be used to generate digital signatures. Generating the signature is done with the private key, while verifying the signature's authenticity is done using the public key. SSL uses these digital signatures to authenticate the identity of web servers.

![Untitled](SSL%20fd5c50efb2eb4fbc9153117c930ef3ca/Untitled.png)

## SSL Certificate

An SSL certificate is a digital document tied to a specific subject, such as a server hosted on a web domain. It contains

- Who the SSL certificate is issued to (**subject**)
- expiration date
- SSL certificate’s public key, which will be used for encryption
- The digital signature of the **certificate authority** who issued the certificate

![Untitled](SSL%20fd5c50efb2eb4fbc9153117c930ef3ca/Untitled%201.png)

SSL protocol uses the SSL certificate to 

- Verify the authenticity of the server
- Encrypt messages sent between client and server
- Check the integrity of the messages sent between client and server.

SSL certificates are signed and issued by CA, third-party corporations who verify the identity of organizations in exchange for the certificate.

## How SSL protocol actually works

[https://namangupta01.medium.com/how-ssl-works-23d8e5ed0cfa#:~:text=SSL is known as Secure Socket Layer.,and other being Web Server](https://namangupta01.medium.com/how-ssl-works-23d8e5ed0cfa#:~:text=SSL%20is%20known%20as%20Secure%20Socket%20Layer.,and%20other%20being%20Web%20Server).

SSL is built on top of the TCP layer, so after a TCP connection is established, the client and server engage in what is called the SSL handshake. The client will send the server the following

- Which version of SSL it’s running
- Which cipher-suites it supports
- Which compression methods it supports

The server will select the highest version of SSL it supports, the cipher-suites, and compression methods.

The server then offers an SSL certificate to confirm its identity. The client attempts to verify the server’s certificate.

***A random secret key is generated at the client end. The client will then take the public key of the server from the certificate, encrypt the secret key and send it to the server. Because the server has the corresponding private key, it can decrypt the client’s random secret key.***

Both the client and server will now have this random secret key. The SSL handshake is over and the two applications can communicate securely by symmetrically encrypting messages with this secret random key.

![Untitled](SSL%20fd5c50efb2eb4fbc9153117c930ef3ca/Untitled%202.png)

[Certificates and Certificate Authority Explained](https://www.youtube.com/watch?v=x_I6Qc35PuQ)