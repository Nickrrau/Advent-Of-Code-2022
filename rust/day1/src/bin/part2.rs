use std::env;
use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

fn main() {
    println!("AoC Day 1 Part 2");
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        println!("Need to provide input");
        return;
    }

    let filename = args[1].clone();
    process_input(filename).expect("Failed to process input");
}

fn process_input(filename: String) -> std::io::Result<()> {
    let file = File::open(filename);
    let file = match file {
        Ok(file) => file,
        Err(error) => return Err(error),
    };

    let buf_reader = BufReader::new(file);

    let mut highest_cal: [u64; 3] = [0, 0, 0];
    let mut tmp_cal = 0;

    for line in buf_reader.lines() {
        let line = match line {
            Ok(line) => line,
            Err(_e) => break,
        };

        if line.len() == 0 {
            push_pop(&mut highest_cal, tmp_cal);
            tmp_cal = 0;
        } else {
            let cal: u64 = line.parse().unwrap();
            tmp_cal = tmp_cal + cal;
        }
    }
    push_pop(&mut highest_cal, tmp_cal);

    println!("{:?}", highest_cal[0] + highest_cal[1] + highest_cal[2]);
    Ok(())
}

fn push_pop(stack: &mut [u64; 3], value: u64) {
    if stack[0] < value {
        stack[2] = stack[1];
        stack[1] = stack[0];
        stack[0] = value;
    } else if stack[1] < value {
        stack[2] = stack[1];
        stack[1] = value;
    } else if stack[2] < value {
        stack[2] = value;
    }
}
