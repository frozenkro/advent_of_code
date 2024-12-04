const std = @import("std");
const print = std.debug.print;
const ArrayList = std.ArrayList;
const Allocator = std.mem.Allocator;
const MyError = error{
    BadDataError,
};

pub fn main() !void {
    std.debug.print("starting..\n", .{});

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const gpaa = gpa.allocator();
    var arena = std.heap.ArenaAllocator.init(gpaa);
    defer arena.deinit();
    const allocator = arena.allocator();

    var data = ArrayList(*ArrayList(i32)).init(allocator);
    defer data.deinit();

    try getData(&data, allocator);

    const safeRows = getSafeRows(data);
    std.debug.print("{}", .{safeRows});
}

fn getSafeRows(data: ArrayList(*ArrayList(i32))) i32 {
    var res: i32 = 0;
    for (data.items) |row_ptr| {
        const row = row_ptr.*;

        const increasing = isIncreasing(row);
        var unsafe: bool = false;
        for (row.items, 0..) |item, iter| {
            if (iter == row.items.len - 1) break;

            const d1 = if (increasing) item + 1 else item - 1;
            const d3 = if (increasing) item + 3 else item - 3;

            const min = @min(d1, d3);
            const max = @max(d1, d3);
            const next = row.items[iter + 1];
            if (next < min or next > max) {
                unsafe = true;
            }
        }
        if (unsafe == false) {
            res += 1;
        }
    }
    return res;
}

fn isIncreasing(data: ArrayList(i32)) bool {
    // This would blow up if any of the test records were less than 4 items :)
    const d1 = data.items[0] < data.items[1];
    const d2 = data.items[1] < data.items[2];
    const d3 = data.items[2] < data.items[3];
    return ((d1 and d2) or (d2 and d3) or (d1 and d3));
}

fn getData(data: *ArrayList(*ArrayList(i32)), allocator: Allocator) !void {
    const file = try std.fs.cwd().openFile("./input", .{});
    defer file.close();
    var br = std.io.bufferedReader(file.reader());
    var in_stream = br.reader();

    var buf: [1024]u8 = undefined;
    var lineNum: usize = 0;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var newSet = try allocator.create(ArrayList(i32));
        newSet.* = ArrayList(i32).init(allocator);
        try data.append(newSet);

        var iterator = std.mem.splitSequence(u8, line, " ");
        var intVal = try std.fmt.parseInt(i32, iterator.first(), 10);
        try newSet.append(intVal);

        while (iterator.next()) |value| {
            intVal = try std.fmt.parseInt(i32, value, 10);
            try newSet.append(intVal);
        }

        lineNum += 1;
    }
}
