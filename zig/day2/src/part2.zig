const std = @import("std");
const proc = std.process;
const io = std.io;
const fs = std.fs;
const os = std.os;

pub fn main() !void {
    const stdout_file = std.io.getStdOut().writer();
    var bw = io.bufferedWriter(stdout_file);
    const stdout = bw.writer();

    try stdout.print("Aoc Day 2 Part 2.\n", .{});
    try bw.flush();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const args = try proc.argsAlloc(allocator);
    defer proc.argsFree(allocator, args);

    if (args.len < 2) {
        try stdout.print("Need to provide input\n", .{});
        try bw.flush();
        return;
    }

    try processInput(args[1]);
    try bw.flush();
}

const Move = u32;
const Rock: Move = 1;
const Paper: Move = 2;
const Scissors: Move = 3;

fn byteToMove(b: u8) Move {
    switch (b) {
        'X', 'A' => return Rock,
        'Y', 'B' => return Paper,
        'Z', 'C' => return Scissors,
        else => return 0,
    }
}

fn processMove(outcome: u8, oppMove: Move) u32 {
    if (outcome == 'Y') return oppMove + 3;

    if (outcome == 'X') {
        switch (oppMove) {
            Rock => return Scissors,
            Paper => return Rock,
            Scissors => return Paper,
            else => {},
        }
    }

    switch (oppMove) {
        Rock => return Paper + 6,
        Paper => return Scissors + 6,
        Scissors => return Rock + 6,
        else => return 0,
    }
}

fn processInput(filename: []const u8) !void {
    const file = try fs.cwd().openFile(filename, .{});
    defer file.close();

    var bufr = io.bufferedReader(file.reader());
    var reader = bufr.reader();

    var score: u32 = 0;

    var buf: [1024]u8 = undefined;
    while (try reader.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        score += processMove(line[2], byteToMove(line[0]));
    }
    try io.getStdOut().writer().print("{}", .{score});
}
