fn main() {
    let lc = last_char(String::from("Hello"));
    println!("The last character is {}", lc);
    let wn = last_char(String::from(""));
    println!("The last character is {}", wn);
}

fn last_char(string: String) -> char {
   if string.is_empty() {
       return '🚨';
   }
   string.chars().next_back().unwrap()
}
