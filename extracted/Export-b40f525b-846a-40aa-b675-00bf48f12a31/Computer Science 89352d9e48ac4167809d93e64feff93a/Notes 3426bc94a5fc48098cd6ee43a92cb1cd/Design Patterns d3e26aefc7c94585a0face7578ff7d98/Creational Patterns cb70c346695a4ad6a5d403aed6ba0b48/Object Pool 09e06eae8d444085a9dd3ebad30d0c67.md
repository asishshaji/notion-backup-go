# Object Pool

- Provides a mechanism for reusing objects that can be very hard to create.
- The keyword is `reusability`
- Useful in situations where the cost of initializing a class is high
- Used to cache the created objects
- Check if the pool has an object available, if so take that object.

```python
class ReusablePool:
    """
    Manage Reusable objects for use by Client objects.
    """

    def __init__(self, size):
        self._reusables = [Reusable() for _ in range(size)]

    def acquire(self):
        return self._reusables.pop()

    def release(self, reusable):
        self._reusables.append(reusable)

class Reusable:
    """
    Collaborate with other objects for a limited amount of time, then
    they are no longer needed for that collaboration.
    """

    pass

def main():
    reusable_pool = ReusablePool(10)
    reusable = reusable_pool.acquire()
    reusable_pool.release(reusable)

if __name__ == "__main__":
    main()
```