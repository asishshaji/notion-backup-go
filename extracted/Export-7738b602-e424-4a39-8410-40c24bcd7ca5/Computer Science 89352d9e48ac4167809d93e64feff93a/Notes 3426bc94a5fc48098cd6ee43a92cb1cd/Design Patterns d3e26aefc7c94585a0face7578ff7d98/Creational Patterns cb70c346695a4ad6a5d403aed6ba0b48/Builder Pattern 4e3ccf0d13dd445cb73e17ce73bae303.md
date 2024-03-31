# Builder Pattern

[Design Patterns and Refactoring](https://sourcemaking.com/design_patterns/builder)

- lets you construct complex objects step by step.
- allows you to produce different types and representations of an object using the same construction code.
- Let’s imagine you need to create an object which has a lot of parameters in its constructor, such initialization code is usually buried inside a monstrous constructor with lots of parameters

[](https://github.com/asishshaji/interview-prep/blob/main/design_patterns/creational_patterns/builder/builder.go)

- The director accepts the builder, and then the builder builds the object that you want.

[Builder in Java / Design Patterns](https://refactoring.guru/design-patterns/builder/java/example)

- Builder interface defines the methods that the concrete builders should follow.
- Concrete builder sets the values

![Untitled](Builder%20Pattern%204e3ccf0d13dd445cb73e17ce73bae303/Untitled.png)

**The key components of the Builder Pattern are:**

- **Product**: Represents the complex object being built.
- **Builder**: Provides an interface for constructing the parts of the product.
- **Concrete Builder**: Implements the builder interface and constructs the parts of the product.
- **Director**: Orchestrates the construction process by using the builder interface.

[Understanding the Builder Pattern in Go](https://dev.to/kittipat1413/understanding-the-builder-pattern-in-go-gp9)