# Monolithic Architecture

Developing complete application as a single unit. In some cases monolithic performs better than other architectures. 

![Untitled](Monolithic%20Architecture%200c56f00839e24e94a555e8a36957f3bb/Untitled.png)

## When to use it

- if you are building small application
- It very easy to build, test, deploy, troubleshoot and scale vertically
- Small team
- Simple application with predictable scale and complexity
- POC and quick launch (For easy validation of a new idea on the market.)

## Benifits

- Easy development, testing, deployment and troubleshoot
- There is no network latency

## Challenges

- Become complex over time - Hard to understand
- Hard to make new changes
- Barrier to new technology adoption
- Any changes to the framework can effect the entire application
- Difficult to scale.

## Design Principles of Monolithic Architecture

- DRY
- KISS
- YAGNI

### DRY

Don’t repeat yourself.

### KISS

Keep It Simple, Stupid

### YAGNI

You Ain’t Gonna Need It : don’t create features that’s not really necessary.