enum WebEvent {
    PageLoad,
    PageUnload,
    KeyPress(char),
    Paste(String),
    Click {x:i64, y:i64},
}

// Option enum
// Defined by standard library
enum Option<T> {
    Some(T),
    None,
}

fn main() {
    let quit = WebEvent::KeyPress('q');

    let something = Some(1);
}
