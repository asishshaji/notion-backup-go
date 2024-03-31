# Tier 1 Analysis

![Untitled](Tier%201%20Analysis%20635c48fbc3ff4412a4647e0b6c62933d/Untitled.png)

## DevTools

![Untitled](Tier%201%20Analysis%20635c48fbc3ff4412a4647e0b6c62933d/Untitled%201.png)

**Preserve Log :** If checked clears the log on page reload/ navigation

**Disable cache :** disables browser cache, eg : disables etag. 

### Navigating to [https://example.com](https://example.com)

![Untitled](Tier%201%20Analysis%20635c48fbc3ff4412a4647e0b6c62933d/Untitled%202.png)

![Untitled](Tier%201%20Analysis%20635c48fbc3ff4412a4647e0b6c62933d/Untitled%203.png)

First size represents the compressed size and second represents the uncompressed size. 

**Content-Encoding** can be enabled to find if the data is compressed or not.

## Waterfall

![Untitled](Tier%201%20Analysis%20635c48fbc3ff4412a4647e0b6c62933d/Untitled%204.png)

[Network features reference - Chrome Developers](https://developer.chrome.com/docs/devtools/network/reference/?utm_source=devtools#timing-explanation)

**Queueing**: if a request is in front of your request, the new request is queued. 

The browser queues requests when:

- There are higher priority requests.
- There are already six TCP connections open for this origin, which is the limit. This applies to HTTP/1.0 and HTTP/1.1 only.
- The browser is briefly allocating space in the disk cache

**Stalled:** The request could be delayed for any of the reasons described in **Queueing**. If the resource is not queued and if it waits then it’s called stalled.

Time the request spent waiting before it could be sent. This time is inclusive of any time spent in proxy negotiation. Additionally, this time will include when the browser waits for an already established connection to become available for re-use, obeying Chrome's [maximum six](https://code.google.com/p/chromium/issues/detail?id=12066) TCP connections per origin rule.

The first thing that happened when the request is made is a DNS lookup. Browser did a DNS query, based on the DNS server configured in the browser. 

**Initial Connection** is the tcp connection + SSL/TLS.  The browser is establishing a connection, including TCP handshakes/retries and negotiating an SSL.

**Waiting for server response: or TIME TO FIRST BYTE (TTFB)** the time till the server responds with the first byte. TTFB is the time it takes from when a client makes an HTTP request to it receiving its first byte of data from the web server.