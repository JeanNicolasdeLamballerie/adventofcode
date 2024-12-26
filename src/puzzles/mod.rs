use lazy_static::lazy_static;
use std::collections::HashMap;
use std::fs::read_to_string;

use aoc2024::day1;

use crate::Solvable;

pub mod aoc2024;

fn get_data(year: usize, day: usize) -> String {
    return read_to_string(format!("./input/{}-{}.txt", year, day)).unwrap();
}

const AVAILABLE_2024: [(bool, bool); 25] = [
    (true, true),   // Day 1
    (true, true),   // Day 2
    (true, true),   // Day 3
    (true, true),   // Day 4
    (true, true),   // Day 5
    (true, true),   // Day 6
    (true, true),   // Day 7
    (false, false), // Day 8
    (false, false), // Day 9
    (false, false), // Day 10
    (false, false), // Day 11
    (false, false), // Day 12
    (false, false), // Day 13
    (false, false), // Day 14
    (false, false), // Day 15
    (false, false), // Day 16
    (false, false), // Day 17
    (false, false), // Day 18
    (false, false), // Day 19
    (false, false), // Day 20
    (false, false), // Day 21
    (false, false), // Day 22
    (false, false), // Day 23
    (false, false), // Day 24
    (false, false), // Day 25
];

const AVAILABLE_YEARS: [usize; 1] = [2024];

pub fn get_years<'available_years>() -> &'static [usize] {
    return &AVAILABLE_YEARS;
}

pub fn get_availability(year: usize) -> [(bool, bool); 25] {
    match year {
        2024 => AVAILABLE_2024,

        _ => unreachable!(),
    }
}

pub fn get_puzzle(year: usize, day: usize) -> Box<dyn Solvable> {
    match year {
        2024 => match day {
            0 => Box::new(day1::Day1::default()),

            _ => unreachable!(),
        },
        _ => unreachable!(),
    }
}

pub struct Info<'a> {
    pub message: &'a str,
}

const INF2024_0: &[Info] = &[Info {
    message: "Hello, first solution !",
}];
const INF2024: &[Info] = &[Info {
    message: "Hello, deleted project !",
}];

lazy_static! {
   pub static ref INFORMATIONS: HashMap<&'static str, &'static [Info<'static>]> =
        [("2024", INF2024), ("2024-0", INF2024_0)]
            // .into_iter()
        .iter()
            .copied()
            .collect();
}
