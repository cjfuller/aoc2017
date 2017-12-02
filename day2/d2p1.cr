def checksum(spreadsheet)
  spreadsheet.map { |row| row.max - row.min }
    .sum
end


def parse_input(filename)
  File.read_lines(filename)
    .map { |line| line.split(/\s+/).map(&.to_i) }
end

test_input = parse_input("./d2p1_test_input.txt")
puts checksum(test_input)

input = parse_input("./d2p1_input.txt")
puts checksum(input)
