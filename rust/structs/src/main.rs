// Classic struct
struct Person {
    name: String,
    age: u8,
    likes_orange: bool
}

// Tuple struct
struct Point2D(u32, u32);

fn main() {
    // Instantiate a classic struct
    let person = Person {
        name:String::from("Adam"),
        likes_orange: true,
        age: 25,
    };

    println!("Person name is: {} and age is: {} and likes orange: {}",
        person.name,
        person.age,
        person.likes_orange
    );

    // Instantiate a tuple struct
    let origin = Point2D(100, 200);

    println!("Point contains {:?} and {:?}", origin.0, origin.1);

    // Destruct a tuple struct instance
    let Point2D(x, y) = origin;

    println!("Point contains {:?} and {:?}", x, y);
}
