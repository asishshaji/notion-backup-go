# Finite State Machine (FSM)

There are a finite number of states in which a program can be in. Within any unique state, the program behaves differently. The program can be switched from one state to another.

Depending on the current state the program may or may not switch to certain other state. These switching rules are called *transitions*. 

State machines are usually implemented with lots of conditional statements that select the appropriate behaviour depending on the current state of the object. “State” is just a set of values of the object’s fields.

![Untitled](Finite%20State%20Machine%20(FSM)%20efb8375d76014c9b80c574f7656d8416/Untitled.png)

Most methods will contain monstrous conditional that pick the proper behaviour of a method according to the current state. Code like this is complicated to maintain because any change to the transition logic may require changing the state conditionals in every method.

The solution to this is a state design pattern.

## State Design Pattern

The s**tate** is a behavioural design pattern that lets an 
object alter its behaviour when its internal state changes. It appears as
if the object changed its class.

[State](https://refactoring.guru/design-patterns/state)

```python
class LightBulb:

    def __init__(self):
        self._state = 'OFF'    # initial state of bulb

    def switch(self):
        if self._state == 'ON':
            self._state = 'OFF'
            print("The light bulb is turned off")
        elif self._state == 'OFF':
            self._state = 'ON'
            print("Let there be light, and there was light")

    def get_bulb_status(self):
        print(f'{self.__class__.__name__} is {self._state}')

bulb = LightBulb()
bulb.get_bulb_status()
bulb.switch()
bulb.get_bulb_status()
bulb.switch()
bulb.get_bulb_status()
```

![Untitled](Finite%20State%20Machine%20(FSM)%20efb8375d76014c9b80c574f7656d8416/Untitled%201.png)

The state pattern suggests that you create new classes for all possible states of an object and extract all state-specific behaviours into these classes.

Instead of implementing all behaviours on its own, the original object, called **context**, stores a reference to one of the state objects that represent its current state and delegates all the state-related work to that object.

Context stores the current state, states are represented in their own classes, and these concrete states must follow a `state` interface.

![Untitled](Finite%20State%20Machine%20(FSM)%20efb8375d76014c9b80c574f7656d8416/Untitled%202.png)

[State in Python / Design Patterns](https://refactoring.guru/design-patterns/state/python/example)