use crate::{puzzles::get_data, Process, SolvePuzzle};

#[derive(Default)]
pub struct Day1 {
    input: String,
}

impl SolvePuzzle for Day1 {
    fn solve_p1(&mut self) -> crate::Solved {
        crate::Solved::Usize(part_one(&self.input))
    }
    fn solve_p2(&mut self) -> crate::Solved {
        crate::Solved::Usize(part_two(&self.input))
    }
}
impl Process for Day1 {
    fn process(&mut self) {
        self.input = get_data(2024, 1);
    }
}

pub fn part_one(input: &str) -> usize {
    let (mut left, mut right) = parse_input(input);
    left.sort();
    right.sort();

    let sum: usize = left.iter().zip(right).map(|(a, b)| a.abs_diff(b)).sum();

    sum
}

pub fn part_two(input: &str) -> usize {
    let (left, right) = parse_input(input);

    let mut seen_ids = std::collections::HashMap::new();

    let sum = left
        .iter()
        .map(|id| {
            let count = seen_ids
                .entry(id)
                .or_insert_with(|| count_occurences(id, &right));
            *id * *count
        })
        .sum();
    sum
}

fn count_occurences(id: &usize, list: &Vec<usize>) -> usize {
    list.iter().filter(|value| **value == *id).count() as usize
}

fn parse_input(input: &str) -> (Vec<usize>, Vec<usize>) {
    let lines = input.lines();

    let mut left_list = vec![];
    let mut right_list = vec![];

    for line in lines {
        let mut split = line.split("   ");
        left_list.push(split.next().unwrap().parse::<usize>().unwrap());
        right_list.push(split.next().unwrap().parse::<usize>().unwrap());
    }

    (left_list, right_list)
}
