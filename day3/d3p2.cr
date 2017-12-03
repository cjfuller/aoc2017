# let's just brute force it...

# MAX_SIZE needs to be large enough to get to the desired number, otherwise,
# it's any arbitrary odd number
MAX_SIZE = 11
OFFSET = (MAX_SIZE - 1) / 2

class Array
  def i2(tup)
    self[tup[0] + OFFSET][tup[1] + OFFSET]
  end

  def set_i2(tup, val)
    self[tup[0] + OFFSET][tup[1] + OFFSET] = val
  end

  def sum_neighbors(tup)
    (-1..1).reduce(0) do |acci, i|
      (-1..1).reduce(acci) do |accj, j|
        coord = {tup[0] + i, tup[1] + j}
        accj + self.i2(coord)
      end
    end
  end
end

def next_inc(pos, grid_size, increment)
    if increment == {0, 1} && pos[1] == (grid_size - 1) / 2
      # upper right corner, change increment to go left
      {-1, 0}
    elsif increment == {-1, 0} && pos[0] == -(grid_size - 1) / 2
      # upper left corner, change increment to go down
      {0, -1}
    elsif increment == {0, -1} && pos[1] == -(grid_size - 1) / 2
      # lower left corner, change increment to go right
      {1, 0}
    else
      increment
    end
end

def solve(num)
  grid = Array.new(MAX_SIZE) { Array.new(MAX_SIZE) { 0 } }
  pos = {0, 0}
  last_pos = {0, 0}
  incr = {1, 0}
  i = 1
  ring = 1

  grid.set_i2(pos, 1)

  while grid.i2(last_pos) <= num
    grid.set_i2(pos, grid.sum_neighbors(pos))
    last_pos = pos
    pos = {pos[0] + incr[0], pos[1] + incr[1]}
    incr = next_inc(pos, ring, incr)
    i += 1
    if i == ring**2 + 1
      incr = {0, 1}
      ring += 2
    end
  end

  grid.i2(last_pos)
end

puts solve(1) # expect: 2
puts solve(6) # expect: 10
puts solve(20) # expect: 23
puts solve(60) # expect: 122
puts solve(800) # expect: 806

puts solve(368078)
