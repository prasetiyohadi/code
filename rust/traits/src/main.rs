pub struct Person {
    name: String
}

pub struct Cat {
    name: String
}

pub struct Rabbit {
    name: String
}

pub trait Eat {
    fn eat_dinner(&self) {
        println!("I eat from a dish")
    }
}

impl Eat for Person {
    fn eat_dinner(&self) {
        println!("I eat from a plate")
    }
}

impl Eat for Cat {
    fn eat_dinner(&self) {
        println!("I eat from a cat bowl")
    }
}

impl Eat for Rabbit {}

struct Film {
    title: String,
    director: String,
    studio: String
}

struct Book {
    title: String,
    author: String,
    publisher: String
}

trait Catalog {
    fn describe(&self) {
        println!("We need more information about this type of media")
    }
}

impl Catalog for Film {
    fn describe(&self) {
        println!(
            "{} was directed by {} through {} studios",
            self.title,
            self.director,
            self.studio
        )
    }
}

impl Catalog for Book {
    fn describe(&self) {
        println!(
            "{} was written by {} and published by {}",
            self.title,
            self.author,
            self.publisher
        )
    }
}

struct Albuum {
    title: String,
    artist: String,
    label: String
}

impl Catalog for Albuum {}

fn main() {
    let person = Person {
        name: String::from("Nell")
    };

    person.eat_dinner();

    let cat = Cat {
        name: String::from("Marvin")
    };

    cat.eat_dinner();

    let rabbit = Rabbit {
        name: String::from("Leia")
    };

    rabbit.eat_dinner();

    let capt_marvel = Film {
        title: String::from("Captain Marvel"),
        director: String::from("Anna Boden and Ryan Fleck"),
        studio: String::from("Marvel")
    };

    capt_marvel.describe();

    let elantris = Book {
        title: String::from("Elantris"),
        author: String::from("Brandon Sanderson"),
        publisher: String::from("Tor Books")
    };

    elantris.describe();

    let let_it_be = Albuum {
        title: String::from("Let it Be"),
        artist: String::from("Beatles"),
        label: String::from("Apple")
    };

    let_it_be.describe();
}
