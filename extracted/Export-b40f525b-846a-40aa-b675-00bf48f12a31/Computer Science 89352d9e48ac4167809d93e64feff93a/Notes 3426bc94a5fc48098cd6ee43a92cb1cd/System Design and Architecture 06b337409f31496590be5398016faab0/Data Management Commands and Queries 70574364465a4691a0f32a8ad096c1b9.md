# Data Management: Commands and Queries

## Cross-Service Queries

In monolithic architectures, it’s very easy and common to query different entities. 

![Untitled](Data%20Management%20Commands%20and%20Queries%2070574364465a4691a0f32a8ad096c1b9/Untitled.png)

Multiple data query is very common. 

Microservice architecture uses polyglot persistence, therefore we need strategies to manage queries between services.

### How can we manage cross-service queries

- Direct HTTP Communication
    
    Not a good solution, there is more coupling between services
    
- Asynchronous communication
- Materialized View Pattern
    
    Reduces inter-service communication