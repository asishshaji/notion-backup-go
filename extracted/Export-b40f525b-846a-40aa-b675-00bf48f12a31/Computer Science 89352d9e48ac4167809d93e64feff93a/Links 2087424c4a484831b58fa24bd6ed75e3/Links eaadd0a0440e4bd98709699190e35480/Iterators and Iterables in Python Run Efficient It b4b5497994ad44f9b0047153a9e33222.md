# Iterators and Iterables in Python: Run Efficient Iterations – Real Python

URL: https://realpython.com/python-iterators-iterables/
Category: Python

![https://files.realpython.com/media/Iterators-and-Iterables-in-Python-What-Are-They-and-How-to-Use-Them_Watermarked.5df74332ac58.jpg](https://files.realpython.com/media/Iterators-and-Iterables-in-Python-What-Are-They-and-How-to-Use-Them_Watermarked.5df74332ac58.jpg)

Iterators power and control the iteration process, while iterables typically hold data you want to iterate over one value at a time.

An iterator is an object that allows you to iterate over data collections, such as lists, tuples, dictionaries, and sets.

Python iterators implement the [iterator design pattern](https://refactoring.guru/design-patterns/iterator)

Iterators take responsibility for two main actions

- Returning the data from a stream or container one item at a time
- Keeping track of the current and visited items.

In summary, an iterator will yield each item or value from a stream or a container while doing all the internal bookkeeping required to maintain the state of the iteration process.

## Python Iterator Protocol

A Python object is considered an iterator when it implements two methods collectively known as the iterator protocol.

.__iter__() and .__next__()

The .__iter__() method returns self object.

The .__next__() method must return the next element in the stream and should return a StopIteration expectation when no more items are available in the data stream.

Python uses iterators under the hood to support every operation requiring iteration, including loops, comprehensions, iterable unpacking, and more.

The most attractive feature of iterators is that it doesn't load the whole stream thing into the memory, it processes one item at a time without exhausting the memory resources.

## Creating different types of iterators

- take the stream of data and yield the data as they appear - classic iterator
- take the stream of data and apply a transformation to the stream
- take no input data, generate new data as a result of some computation to finally yield the generated items.

![classic iterator](Iterators%20and%20Iterables%20in%20Python%20Run%20Efficient%20It%20b4b5497994ad44f9b0047153a9e33222/Untitled.png)

classic iterator

![Untitled](Iterators%20and%20Iterables%20in%20Python%20Run%20Efficient%20It%20b4b5497994ad44f9b0047153a9e33222/Untitled%201.png)

![for loop working](Iterators%20and%20Iterables%20in%20Python%20Run%20Efficient%20It%20b4b5497994ad44f9b0047153a9e33222/Untitled%202.png)

for loop working

![Generator functions](Iterators%20and%20Iterables%20in%20Python%20Run%20Efficient%20It%20b4b5497994ad44f9b0047153a9e33222/Untitled%203.png)

Generator functions

![Generator comprehensions](Iterators%20and%20Iterables%20in%20Python%20Run%20Efficient%20It%20b4b5497994ad44f9b0047153a9e33222/Untitled%204.png)

Generator comprehensions

Generator functions are special types of functions that allow you to create iterators using a functional style. Unlike regular functions, which typically compute the value and then return it to the caller, generator functions return a generator iterator that yields a stream of data one value at a time.

A generator function returns an iterator that supports the ***iterator protocol*** out of the box. So, generators are also iterators.

## Iterables

An Iterable is an object that you can iterate over. Pure iterable objects typically hold the data themselves. 

Iterators are also iterables.

![Untitled](Iterators%20and%20Iterables%20in%20Python%20Run%20Efficient%20It%20b4b5497994ad44f9b0047153a9e33222/Untitled%205.png)

The `.__iter__()` method fulfils the ***iterable protocol***. This method must return an iterator object, which usually doesn't coincide with self unless your iterable is also an iterator.

![Untitled](Iterators%20and%20Iterables%20in%20Python%20Run%20Efficient%20It%20b4b5497994ad44f9b0047153a9e33222/Untitled%206.png)

All iterators are iterables because they implement __iter__(), ie iterable protocol.

Not all iterables are iterators, because iterables don't implement .__next__() method, only .__iter__() method.

![Untitled](Iterators%20and%20Iterables%20in%20Python%20Run%20Efficient%20It%20b4b5497994ad44f9b0047153a9e33222/Untitled%207.png)