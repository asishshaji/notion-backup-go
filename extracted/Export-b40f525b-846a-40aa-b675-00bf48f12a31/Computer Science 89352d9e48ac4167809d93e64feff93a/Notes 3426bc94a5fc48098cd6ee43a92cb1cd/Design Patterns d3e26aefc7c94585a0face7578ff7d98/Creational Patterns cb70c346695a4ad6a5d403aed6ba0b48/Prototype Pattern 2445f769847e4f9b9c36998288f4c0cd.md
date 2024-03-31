# Prototype Pattern

[Prototype](https://refactoring.guru/design-patterns/prototype)

- used for creating a new object out of the previous one
- delegates the actual cloning process to the objects that are being cloned.
    - Why?? Because some attributes of the class might be private, so it makes sense to have the clone method inside the class itself.
- Defines a common interface for objects that want to support cloning
    - Usually, the interface contains only the clone() method.

- An object that supports cloning is called a `prototype`
- Various actors are
    - Prototype Interface, which contains clone() method
    - Concrete Prototype which implements Prototype Interface
    
    ![Real-world example](Prototype%20Pattern%202445f769847e4f9b9c36998288f4c0cd/Untitled.png)
    
    Real-world example
    
    ![Untitled](Prototype%20Pattern%202445f769847e4f9b9c36998288f4c0cd/Untitled%201.png)
    

![Untitled](Prototype%20Pattern%202445f769847e4f9b9c36998288f4c0cd/Untitled%202.png)

- You can also maintain a prototype registry to keep access to frequently accessed prototypes
- That’s optional