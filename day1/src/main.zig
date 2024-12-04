const std = @import("std");
const print = std.debug.print;
const ArrayList = std.ArrayList;

pub fn main() !void {
    // Prints to stderr (it's a shortcut based on `std.io.getStdErr()`)
    std.debug.print("All your {s} are belong to us.\n", .{"codebase"});
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var list1 = std.ArrayList(usize).init(allocator);
    defer list1.deinit();
    var list2 = std.ArrayList(usize).init(allocator);
    defer list2.deinit();

    try getItems(&list1, &list2);
    print("got items\n", .{});

    quickSort(list1.items);
    print("first list sorted\n", .{});
    quickSort(list2.items);
    print("second list sorted\n", .{});
    if (list1.items[0] == 61967) {
        print("error", .{});
    }

    var total: usize = 0;
    for (list1.items, 0..) |item, iter| {
        total += getAbsDiff(item, list2.items[iter]);
    }
    print("{}\n", .{total});
}

fn getAbsDiff(item1: usize, item2: usize) usize {
    return if (item1 < item2) item2 - item1 else item1 - item2;
}

fn getItems(list1: *ArrayList(usize), list2: *ArrayList(usize)) !void {
    const file = try std.fs.cwd().openFile("./input", .{});
    defer file.close();
    var br = std.io.bufferedReader(file.reader());
    var in_stream = br.reader();

    var buf: [1024]u8 = undefined;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var iterator = std.mem.splitSequence(u8, line, "   ");
        const item1 = try std.fmt.parseInt(usize, iterator.first(), 10);
        const item2 = try std.fmt.parseInt(usize, iterator.next().?, 10);

        try list1.append(item1);
        try list2.append(item2);
    }
}

fn quickSort(arr: []usize) void {
    sortSub(arr, 0, arr.len - 1);
}
fn sortSub(arr: []usize, low: usize, high: usize) void {
    if (low < high) {
        const pivot = partition(arr, low, high);
        sortSub(arr, low, @min(pivot, pivot -% 1));
        sortSub(arr, pivot + 1, high);
    }
}
fn partition(arr: []usize, low: usize, high: usize) usize {
    const pivot: usize = high;
    var lowP = low;
    var hiP = high;

    while (lowP < hiP) {
        while (arr[lowP] < arr[pivot]) {
            lowP += 1;
        }
        while (arr[hiP] > arr[pivot]) {
            hiP -= 1;
        }
        if (lowP > hiP) {
            break;
        }

        const temp: usize = arr[lowP];
        arr[lowP] = arr[hiP];
        arr[hiP] = temp;
        lowP += 1;
    }
    const temp: usize = arr[hiP];
    arr[hiP] = arr[pivot];
    arr[pivot] = temp;
    return hiP;
}

test "simple test" {
    var list = std.ArrayList(usize).init(std.testing.allocator);
    defer list.deinit(); // Try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(usize, 42), list.pop());
}

test "fuzz example" {
    const global = struct {
        fn testOne(input: []const u8) anyerror!void {
            // Try passing `--fuzz` to `zig build test` and see if it manages to fail this test case!
            try std.testing.expect(!std.mem.eql(u8, "canyoufindme", input));
        }
    };
    try std.testing.fuzz(global.testOne, .{});
}

test "quick sort desc" {
    var arr = [_]usize{ 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1 };
    quickSort(&arr);
    for (arr, 0..) |value, idx| {
        try std.testing.expectEqual(@as(usize, idx + 1), value);
    }
}
test "quick sort dupe" {
    // should add timeout here actually, bug is infinite loop
    var arr = [_]usize{ 15, 14, 14, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1 };
    const exp = [_]usize{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 14, 15 };
    quickSort(&arr);
    for (arr, 0..) |value, idx| {
        try std.testing.expectEqual(exp[idx], value);
    }
}
