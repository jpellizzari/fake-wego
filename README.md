## Overview

### Models

- The `models` directory contains data structures for the different WeGO domain entities. These are generally `structs` that do not contain logic.
- Models objects may have helper methods for common situations, such as converting between types, accessing a computed property, etc.
- Models are typically thought of as "immutable", meaning that methods that mutates a model should return a new copy of that model (pass by value).

### Services

- Services represent the business logic of our program
- Services typically consume models and interact with adapters to execute operations
- Services will often consume other services
- Services are generally named based on what they do: `Applier`, `Creator`, `Fetcher`

### Adapters

- Adapters are the "machinery" that is used to interact with the outside world
- Adapters can be incoming or outgoing
- Models do not use Adapters directly
- An example incoming adapter would be the HTTP API server
- An example outgoing adapter would be a Kubernetes client or `go-git-providers` client.

## FAQ

Q: **When does a function belong on the model and when does it belong in a services?**
A: Models reciever methods should be focused on performing common data manipulation operations on their parent object. An example of a good use of model methods would be a method that converts the model object into another type. An example of a bad use of model methods would be a method that retrieves something from a cluster or external HTTP API server.

Q: **Can two models satisfy the same interface**?
A: Yup

Q: **Should model functions be on pointers (func (f \*foo) vs (f foo)) so that they can manipulate the model?**
A: Typically, we would want the reciever to be passed by value. If we are manipulating data of a Model instance, that may be an indication that we should use a service or return another type instead.

Q: **Can service methods accept other services as arguments?**
A: You betcha

Q: **How do I instantiate a Model? Should I use a constructor or fill out a `struct`**?
A: Constructors are preferable, but if a constructor requires too many arguments (more than 4), or a "params" `struct`, consider filling out a `struct` and using a `Validate` method instead.

Q: **Do Models always have a corresponding Service (and vice versa)?**:
A: No, not necessarily
