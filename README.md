## GO:

1.  Go is `faster` than the interpreted languages like Python, JavaScript, Php.
2.  Go has a `fast compilation` speed than Rust, C++, Java.
3.  It doesn't have the same `execution speed` as other languges.

### Commands:

```bash
 go mod init <module-name>
 go run main.go
```

### Compilation:

1. main.go will not be understood by the computer.
2. During the compilation we create an executable program.

- `Distributing` programs that are natively compiled is much easier.
- If we distribute python code, then the other person needs to have python `installed`.

### Statically & Strongly Typed:

- Go is both `statically` typed and `strongly` typed language.
- Go enforces static typing meaning variable types are `known` before the code runs.
- Strongly typed meaning the variable types are `fixed` and cannot be changed.

### Garbage Collection:

- Go has garbage collection, which allows to `clean up` resources.
- However, it does not have `JVM` like Java, allowing minimal memory usage.
- A small code is added known as `Go Runtime`, which handles memory management.

### Functions:

- `defer` keyword is used to close a resource.
- The parameters will be checked but the function will be called before the current function `returns`.

### Interfaces:

1. Keep interfaces `small`.
2. Interfaces should have `no knowledge` about satisfying types. (Ex: isFireTruck())
3. Interfaces are not `classess`.

### Errors:

1. Go has `errors` package, which allows to create custom errors.
2. We should avoid using `panic` and `recover` keywords.
3. Alternative to that is the `Log.Fatal()` method.