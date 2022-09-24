fn main() {
    let say = String::from("Cat");
    // Error because value moved when used to call print_out function
    // print_out(say);
    // Use clone() to duplicate the stack and the heap
    // print_out(say.clone());
    // Use borrowing to get the reference to the value
    print_out(&say);
    println!("Again: {}", say);

    let mut my_vec = vec![1, 2, 3];
    println!("{:?}", my_vec);

    add_to_vec(&mut my_vec);
    println!("{:?}", my_vec);
}

// original function
// fn print_out(to_print: String) {
//     println!("{}", to_print);
// }

// function using reference to get the argument value
fn print_out(to_print: &String) {
    println!("{}", to_print);
}

fn add_to_vec(a_vec: &mut Vec<i32>) {
    a_vec.push(4);
}
