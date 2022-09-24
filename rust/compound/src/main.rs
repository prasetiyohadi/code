fn main() {
    // Array using inferred type
    let array = [1u32, 2, 3];

    // Array with type declaration
    // let array: [u32; 3] = [1u32, 2, 3];

    // ERROR: mismatched types
    // Error because number actual of elements is larger than expected number
    // of elements in type declaration
    // let array: [u32; 3] = [1u32, 2, 3, 4];

    // ERROR: mismatched types
    // Error because the type of the elements is not homogenous
    // let array = [1, 2, true];

    // Access first element of the array
    let first_element = array[0];

    // ERROR: will panic at runtime: index out of bounds
    // let warning = array[100];

    // Get length of string
    let length = "Some text".len();

    // PANIC at runtime: index out of bounds
    // [1][length];

    // Tuple using inferred types
    let tuple = (1u32, 2, true);

    // Tuple with type declaration
    // let tuple: (u32, i8, bool) = (1u32, 2, true);

    // Access first element of tuple
    let first_element = (1, 2, true).0;

    // ERROR: no field `100`
    // let error = (1, 2, true).100;
}
