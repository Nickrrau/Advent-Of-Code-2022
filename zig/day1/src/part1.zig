const std = @import("std");
const proc = std.process;
const io = std.io;
const fs = std.fs;
const os = std.os;

pub fn main() !void {
    const stdout_file = std.io.getStdOut().writer();
    var bw = io.bufferedWriter(stdout_file);
    const stdout = bw.writer();

    try stdout.print("Aoc Day 1 Part 1.\n", .{});
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

fn processInput(filename: []const u8) !void {
    const file = try fs.cwd().openFile(filename, .{});
    defer file.close();

    var highest_cal: u64 = 0;
    var tmp_cal: u64 = 0;

    var buf: [1024]u8 = undefined;
    while (file.reader().readUntilDelimiter(&buf, '\n')) |line| {
        if (line.len == 0) {
            tmp_cal = 0;
            continue;
        }

        var cal = try std.fmt.parseInt(u64, line, 10);
        tmp_cal = tmp_cal + cal;

        if (highest_cal < tmp_cal) highest_cal = tmp_cal;
    } else |err| {
        if (err != error.EndOfStream) return err;
    }

    try io.getStdOut().writer().print("{}", .{highest_cal});
}
