def find_divisors_and_divide(row)
  row.each.with_index do |m, i|
    row.each_with_index do |n, j|
      next if i == j
      if m % n == 0
        return m / n
      end
    end
  end
  0
end

def solve(spreadsheet)
  spreadsheet
    .map { |x| find_divisors_and_divide(x) }
    .sum
end

def parse_input(filename)
  File.read_lines(filename)
    .map { |line| line.split(/\s+/).map(&.to_i) }
end

test_input = parse_input("./d2p2_test_input.txt")
puts solve(test_input)

input = parse_input("./d2p1_input.txt")
puts solve(input)
