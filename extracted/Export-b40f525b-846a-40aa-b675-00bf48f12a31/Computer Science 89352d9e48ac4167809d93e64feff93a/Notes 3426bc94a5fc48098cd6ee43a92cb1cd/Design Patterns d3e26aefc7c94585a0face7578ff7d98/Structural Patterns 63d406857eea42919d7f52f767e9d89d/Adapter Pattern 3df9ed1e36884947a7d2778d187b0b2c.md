# Adapter Pattern

- also known as the wrapper
- structural pattern that allows objects with incompatible interfaces to collaborate

## How it works

- create an adapter interface, which is compatible with one of the existing objects
- Using this interface, the existing object can safely call the adapter’s methods

We are trying to make the target compatible with our object, the adapter should have the same API as the target, so that it’s compatible.

Once the adapter has the function that the target has, when a call with the target’s signature is made to the adapter, we can trigger the processing. DAFAQ!!

![Untitled](Adapter%20Pattern%203df9ed1e36884947a7d2778d187b0b2c/Untitled.png)