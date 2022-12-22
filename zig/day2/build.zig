const std = @import("std");

pub fn build(b: *std.build.Builder) void {
    // Standard target options allows the person running `zig build` to choose
    // what target to build for. Here we do not override the defaults, which
    // means any target is allowed, and the default is native. Other options
    // for restricting supported target set are available.
    const target = b.standardTargetOptions(.{});

    // Standard release options allow the person running `zig build` to select
    // between Debug, ReleaseSafe, ReleaseFast, and ReleaseSmall.
    const mode = b.standardReleaseOptions();

    part1Setup(b, target, mode);
    part2Setup(b, target, mode);
}

fn part1Setup(b: *std.build.Builder, target: std.zig.CrossTarget, mode: std.builtin.Mode) void {
    const exe = b.addExecutable("part1", "src/part1.zig");
    exe.setTarget(target);
    exe.setBuildMode(mode);
    exe.install();

    const run_cmd = exe.run();
    run_cmd.step.dependOn(b.getInstallStep());
    if (b.args) |args| {
        run_cmd.addArgs(args);
    }

    const run_step = b.step("run-part1", "Run the part 1 solution");
    run_step.dependOn(&run_cmd.step);
}

fn part2Setup(b: *std.build.Builder, target: std.zig.CrossTarget, mode: std.builtin.Mode) void {
    const exe = b.addExecutable("part2", "src/part2.zig");
    exe.setTarget(target);
    exe.setBuildMode(mode);
    exe.install();

    const run_cmd = exe.run();
    run_cmd.step.dependOn(b.getInstallStep());
    if (b.args) |args| {
        run_cmd.addArgs(args);
    }

    const run_step = b.step("run-part2", "Run the part 2 solution");
    run_step.dependOn(&run_cmd.step);
}
