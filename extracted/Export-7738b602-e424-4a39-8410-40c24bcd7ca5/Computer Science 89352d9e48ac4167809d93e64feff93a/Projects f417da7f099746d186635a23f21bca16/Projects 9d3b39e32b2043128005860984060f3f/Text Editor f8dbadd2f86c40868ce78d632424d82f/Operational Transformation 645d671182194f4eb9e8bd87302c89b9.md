# Operational Transformation

[Operational Transformations as an algorithm for automatic conflict resolution](https://medium.com/coinmonks/operational-transformations-as-an-algorithm-for-automatic-conflict-resolution-3bf8920ea447)

Earlier Google Docs versions maintained synchronized data between clients and servers. The server received changes, merged the data, and sent it to other clients. If other clients had unsent changes, they merged it with the server's version, repeating the process.

![An earlier version, which can lead to inconsistency in data](Operational%20Transformation%20645d671182194f4eb9e8bd87302c89b9/Untitled.png)

An earlier version, which can lead to inconsistency in data

This approach may not always work, and data consistency isn't always maintained. Additionally, the merging operation can be challenging.

Consider an alternative approach: store the document as a sequence of changes and apply them in order. View the document as the final state, with the changes acting as events that alter the state and lead to this final state (the document itself).

So, instead of comparing versions of documents, send events. These are applied in sequence to generate the actual document.

Basic operations that you need for a real-time text editor

- Insert text
- Delete text
- Apply style - not a necessity

When you edit the file, the changes are added to a changelog. You can then replay this changelog on other nodes to recreate the exact document.

## Example

Alice and Bob start from a synchronized document “LIFE 2017”

Alice changed 2017 to 2018 via two operations.

![Untitled](Operational%20Transformation%20645d671182194f4eb9e8bd87302c89b9/Untitled%201.png)

At the same time, Bob changes the text to “CRAZY LIFE 2017”

![Untitled](Operational%20Transformation%20645d671182194f4eb9e8bd87302c89b9/Untitled%202.png)

If Bob applies Alice’s delete operation when he receives it then he gets an incorrect document.

![Untitled](Operational%20Transformation%20645d671182194f4eb9e8bd87302c89b9/Untitled%203.png)

The digit 7 was intended to be deleted, but instead, F was removed. Despite the theory suggesting that the element at index 8 should be deleted, this is not semantically correct. The element at index 14 should have been removed.

Bob needs the transform the delete operation according to his local changes to get the correct state of the document.

![Untitled](Operational%20Transformation%20645d671182194f4eb9e8bd87302c89b9/Untitled%204.png)

![Untitled](Operational%20Transformation%20645d671182194f4eb9e8bd87302c89b9/Untitled%205.png)

![Operational Transformation](Operational%20Transformation%20645d671182194f4eb9e8bd87302c89b9/Untitled%206.png)

Operational Transformation