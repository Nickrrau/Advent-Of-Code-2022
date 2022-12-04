use std::env;
use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

fn main() {
    println!("AoC Day 1 Part 1");
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

    let mut highest_cal = 0;
    let mut tmp_cal = 0;

    for line in buf_reader.lines() {
        let line = match line {
            Ok(line) => line,
            Err(_e) => break,
        };

        if line.len() == 0 {
            tmp_cal = 0;
            continue;
        }

        let cal: u64 = line.parse().unwrap();
        tmp_cal = tmp_cal + cal;

        if highest_cal < tmp_cal {
            highest_cal = tmp_cal;
        }
    }

    println!("{}", highest_cal);
    Ok(())
}
