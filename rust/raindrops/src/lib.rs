use std::string::String;

pub fn raindrops(n: u32) -> String {
    let is_factor = |factor: u32| n % factor == 0;
    let mut res = String::new();
    if is_factor(3) {
        res.push_str("Pling")
    }
    if is_factor(5) {
        res.push_str("Plang")
    }
    if is_factor(7) {
        res.push_str("Plong")
    }
    if res.is_empty() {
        res = n.to_string();
    }
    res
}
