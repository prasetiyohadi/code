fn main() {
    println!("Variables");

    // Bound variable to value using keyword `let`
    let x: i32 = 5;
    println!("The value of x is {}", x);

    let y: bool = true;
    println!("The value of y is {}", y);

    let y: bool = false;
    println!("The value of y is {}", y);

    // Variable is immutable by default
    let a = 5;
    println!("The value of a is {}", a);

    // Error because a is immutable
    //a = 6;
    //println!("The value of a is {}", a);

    // Use keyword `mut` to set variable mutable
    let mut b = 5;
    println!("The original value of b is {}", b);

    b = 6;
    println!("The new value of b is {}", b);

    // Immutability creates predictability in our code
    // Can be more convenient to make variables mutable
    // There are trade-offs either way
    // A similar concept as **constants**

    println!("Constants");

    // Declare a constant using keyword `const`
    // Constant is always immutable
    // Name of the constant should be capitalized with underscores
    // between words
    // Can only be set to an expression
    const SCORE_LIMIT: u32 = 100;
    println!("The value of constant SCORE_LIMIT is {}", SCORE_LIMIT);

    const STRING: &str = "Hullo";
    println!("The value of constant STRING is {}", STRING);

    println!("Variable Shadowing");

    // Declare a new variable with the same name as the previous, creating a new binding
    // The new variable *shadows* the previous variable
    let c = 5;
    println!("The original value of c is {}", c);

    let c = c + 1; // c is now 6
    println!("The new value of c is {}", c);
}
