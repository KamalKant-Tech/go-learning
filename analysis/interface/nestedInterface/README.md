# Nested interface
Like a struct, an interface can also be nested in a struct. In Layman’s terms, it means that a field can have a data type of an interface.

Since we know that, an interface type is a declaration of method signatures. Any data type that implements an interface can also be represented as a type of that interface (polymorphism).

Using this polymorphism principle, we can have a struct field of an interface type and its value can be anything that implements that interface. Let’s see this in action using the earlier example.