// This implementation is a lot slower
//
// fn is_prime(n: &u32) -> bool {
//     if n < &2 {
//         return false;
//     }
//     !(2..n-1).any(|i| n % i == 0)
// }

fn is_prime(n: &u32) -> bool {
    if n <= &1 {
        return false;
    }

    let mut i = 2;
    while &i < n {
        if n % &i == 0 {
            return false;
        }
        i += 1;
    }
    true
}

fn next_prime(n: &u32) -> u32 {
    let mut curr = n + 1;
    while !is_prime(&curr) {
        curr += 1;
    }
    curr
}

pub fn nth(n: u32) -> u32 {
    let mut i = 0;
    let mut curr = 2;

    while i < n {
        curr = next_prime(&curr);
        i += 1;
    }
    curr
}

pub fn main() {
    let mut i = 2;
    while i < 200000 {
        if is_prime(&i) {
            println!("{}", i);
        }
        i += 1;
    }
}