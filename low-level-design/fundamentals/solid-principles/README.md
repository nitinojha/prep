# Solid Principles

S - Single Responsibility Principle
O - Open / Closed Principle
L - Liskov Substitution Principle
I - Interface Segmented Principle
D - Dependency Inversion Principle

Advantages:
1. Avoid duplicate code
2. Easy to maintain
3. Easy to understand
4. Flexible software
5. Reduce complexity

## [S — Single Responsibility](./single-responsibilty-principle.md)
- A class should have single responsibility
- A class should have only 1 reason to change

**Goal** - This principle aims to separate behaviours so that if bugs arise as a result of your change, it won’t affect other unrelated behaviours.

## [O - Open/Closed Principle](./open-closed-principle.md)
- Open for extension but closed for modification
  
**Goal** - This principle aims to extend a Class’s behaviour without changing the existing behaviour of that Class. This is to avoid causing bugs wherever the Class is being used.

## [L- Liskov Substitution Principle](./liskov-substitution-principle.md)
- If Class B us subtype of Class A, then we should be able to replace object of A with B without breakign the behaviour of the program.
- Subclass should extend the capabilty of parent class not narrow it down.

**Goal:** This principle aims to enforce consistency so that the parent Class or its child Class can be used in the same way without any errors.


## [I — Interface Segregation Principle](interface-segregation-principle.md)
- Interface should be such that client should implement unnecessary functions they do not need.
- Clients should not be forced to depend on methods that they do not use.

**Goal:** This principle aims at splitting a set of actions into smaller sets so that a Class executes ONLY the set of actions it requires.

## [D — Dependency Inversion Principle](./dependency-invserion-principle.md)
- Class should depend on interfaces rather than concrete classes
- High-level modules should not depend on low-level modules. Both should depend on the abstraction.
- Abstractions should not depend on details. Details should depend on abstractions.

**Goal:** This principle aims at reducing the dependency of a high-level Class on the low-level Class by introducing an interface.

## **References:**
- https://medium.com/backticks-tildes/the-s-o-l-i-d-principles-in-pictures-b34ce2f1e898
- https://www.youtube.com/watch?v=XI7zep97c-Y