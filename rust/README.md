# Learn programming with Rust

## Data Types

### Number Types

**Integer**

* Integers are whole numbers
* They are either signed or unsigned
* Signed variants
    * Can be positive or negative numbers
    * Store -(2^(n-1)) to 2^(n-1)-1
* Unsigned variants:
    * Are always positive
    * Store: 0 to 2^(n-1)

| Length  | Signed | Unsigned |
| ------  | ------ | -------- |
| 8-bit   | i8     | u8       |
| 16-bit  | i16    | u16      |
| 32-bit  | i32    | u32      |
| 64-bit  | i64    | u64      |
| 128-bit | i128   | u128     |
| arch    | isize  | usize    |

* Rust compiler must know the data types for each variable
* The data type can be inferred
* Rust defaults to i32

**Floating-point**

* Numbers with decimal points
* Types: **f32** and **f64**
* Default type is f64 for modern CPUs

**Numerical operations**

* Number types support mathematical operations
    * **( + )** = addition
    * **( - )** = subtraction
    * **( &ast; )** = multiplication
    * **( / )** = division
    * **( % )** = remainder

### Booleans

* Have value of either **true** or **false**
* Specified using the keyword **bool**
* One byte size
* Used most in conditional and control flow statements

### Characters

* Represent letters
* Specified using the **char** keyword
* Use single quotes
* Four bytes in size

## Compound Data Types

### Arrays

* Continuous group of items
* Fixed length
* Length is known at compile time
* Homogeneous (only contains same data type)
* i.e., `let array: [u32, 4] = [1, 2, 3, 4];`
* Access item using index `array[0]`

### Tuples

* Continuous group of items
* Fixed length
* Length is known at compile time
* Heterogeneous (can contains multiple data types)
* i.e., `let tuple: (bool, u32, u16) = (true, 8, 9);`
* Access item using index `tuple.0`
* Empty tuple is called **unit**

## Functions

* Function definition
    * Start with keyword **fn**
    * The function name
    * A list of arguments
    * An optional return type
    * The function body
* Arguments' type always required
* Return type always required if value is returned
* If not, return type is unit or empty tuple

## Structs

* A type that is composed of other types
* Can contain different types
* Struct is defined using keyword **struct** followed by the name
* The name should describe the object appropriately
* Instantiate struct by supplying **key: value** pairs
* Three flavors of structs
    * Classic
    * Tuple
    * Unit

**Classic**

* Most commonly used
* Each field has a name and a type

**Tuple**

* Similar to classic structs
* Their fields have no name

**Unit**

* Have no fields
* Similar to `()` unit type

## Enums

* List all variations of some data
* Common feature across programming languages (i.e., F#; Haskell)
* Referred to as algebraic data types
* Enum is defined using keyword **enum** followed by the name
* List all variations inside the body
* The enum is now a custom data type that can be used in the code
* The definition didn't include a type, need to associate a type another way
* The enum's variants can also specify their type
* An enum variant can include any kind of data
* An enum can have a variety of types
* Similar to structs but with more flexibility and advantages
    * Describe what kind of data will be stored
    * Each variant can have a different type
    * All variants stored under the custom enum type

## Control Flow: If/Else and Match

### If expressions

* The most common ways to introduce control flow and branch code
* Provide a condition, then execute a block of code if the condition is met

### If/else expressions

* An else expression can be added optionally
* If no else is provided the block will be skipped if the condition is false

### Else if expressions

* If there are more than two conditions, if/else can be combined with else if
* In the case all **if** and **else if** is false, the else block is executed

### Match

* Rust provides pattern matching with the **match** keyword
* Similar to **switch** in other languages
* A **scrutinee** expression is provided to compare to the pattern
* Arms are evaluated and compared with the scrutinee expression
* Each arm has a pattern and some code separated by "=>" operator
* The first arm with a matching pattern is chosen as the branch target
* If there is no matching pattern, arm with pattern "_" will be chosen

## Loops

### Loop

* Execute over a block of code forever
* Use keyword **loop**
* Use **break** keyword to stop the loop

### While

* Conditional loops
* Run until condition is met or become false

### For

* Iterate over elements in a collection
* Each pass of the loop extracts values

## Error Handling

### Panic Macro

* Simplest way to handler errors
    * Failure message is printed
    * Program unwindes and cleans up the stack
    * Program quits
* Should only be used when a program comes to an unrecoverable state
* Rust emits a panic during code execution

### The Option Enum

* Manages the possibility of nonexistent values
* Type **T** is generic and associated with the **Some** variant
* **None** indicates that no element was found
* **Some** means that an element of type **T** was found

### The Result Enum

* Used for recoverable errors that are more common
* The Ok(T) variant represents a success and contains a value
* The Err(E) variant represents an error
* Used for input/output operations (I/O)
    * Parsing strings
    * File access
    * Data validation
* Best for expected failures

### Unwrap and Expect

* Helper methods of the Result type
* **Unwrap** returns the value inside Ok variant
* **Unwrap** returns a panic! macro for the Err variant
* **Expect** returns a value or called the panic! macro with a detailed error

### The ? Operator

* Similar to a match statement
* For Result type:
    * Unwraps the value if Ok variant
    * Returns an error if Err variant
* For Option type:
    * Returns a value is with the Some variant
    * Returns nothing for the None variant

## Ownership

* Related to how your program manages memory
* Rust stores data in two different structure parts of memory:
    * Stack
    * Heap
* Stack:
    * Stack data structure
    * LIFO (last in, first out)
    * Known, fixed size
    * i.e., i32:
        * An i32 has a known, fixed size at compile time
        * Will take up 32 bits of memory
        * Variable and value stored on the **stack**
* Vector:
    * Vector is **mutable**
    * Vector size can **change when the program is running**
    * Vector object store on **stack** with **pointer to heap**
    * **Value** of vector is stored on **heap**
* String:
    * The **String** is **mutable**
    * A **String** stores data on both the **stack** and the **heap**
    * **String** object stored on **stack** with **pointer to heap**
* Dangling pointer:
    * Pointer on the stack pointing to data on the heap that no longer valid
    * C and C++ depends on the programmer to explicitly manages the memory
    * Python, Ruby, JS (and more) use garbage collector
    * Rust use concept of ownership
* Rules of ownership in Rust:
    * Each value in Rust has a variable that is called its **owner**
    * There can only be **one owner** of a value at a time
    * When the owner goes **out of scope**, the **value** will be **dropped**:
        * Variable's scope:
            * **Range** within a program for which that **variable** and the **value** are valid
* Ownership prevents memory safety issues:
    * Dangling pointer
    * Double free: trying to free memory that has already been freed
    * Memory leaks: not freeing memory that should have been freed

## Borrowing

### Clone

* Copies data on both stack and heap

### Borrowing

* Uses ampersand (&) operator
* Creates a reference to a value
* Reference is immutable by default
* Rules of borrowing:
    * At any given time, you can have either:
        * One mutable reference, or
        * Any number of immutable references
    * References must always be **valid**

## Strings

* UTF-8 encoded
* Non-null-byte terminated
* Not collections of chars
* String types:
    * *String* and *&str*
    * *CString* and *&CStr*
    * *OsString* and *&OsStr*
* The *String* type:
    * An owned string
    * Owns string data
    * Data freed when dropped
    * *String* memory:
        * Three parts: Length, Capacity, Data pointer
* The *&str* type:
    * A *borrowed* string slice
    * Does *not* own string data
    * Data *not* freed when dropped
    * View/window into string data
    * *&str* memory:
        * Two parts: Length, Data pointer
* String literals:
    * Embedded into the binary
    * Have type *&str*
* Characters:
    * Strings are *not* arrays of characters
    * Not every byte is a full characters
    * Access data using explicit methods

## Collections

* Vec<T> :
    * Access by index
    * *Vec* memory:
        * Three parts: Length, Capacity, Data pointer
* Slices:
    * Continuous chunks of memory
    * Refer to data owned by another value
    * Written as &[T]
    * *&[T]* memory:
        * Two parts: Length, Data pointer
* HashMap<K, V> :
    * Associate keys and values
* HashSet<T> :
    * Unique items & no access by index
* VecDeque<T> :
    * A queue
* LinkedList<T>

## Traits

* Object-oriented programming:
    * Object: Data and Behavior
* Rust:
    * Enum/Structs: Data
    * Traits: Behavior
* Struct:
    * Stores data
    * i.e., Person; Cat
* Trait:
    * Stores behavior
    * i.e., Eat
    * Can define default behavior for a method
* Implementing a Trait:
    * Use the **impl** and **for** keywords
* Using a Trait:
    * Create an instance of a struct
    * Call a trait method on the instance
