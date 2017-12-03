def all_odd_nums()
  i = 1
  while true
    yield i
    i += 2
  end
end

def solve(num)
  # num will be in the ring of size n, such that (n-1)**2 < num < n**2, and n
  # is odd
  # first let's find n:
  grid_size = 1
  all_odd_nums do |s|
    if num <= s**2
      grid_size = s
      break
    end
  end

  # which (0-indexed) position is it in in the current ring?
  pos_in_ring = num - (grid_size - 2) ** 2 - 1

  # now just brute force the coords in the ring
  pos = {(grid_size - 1)/2, -Math.max((grid_size - 1)/2 - 1, 0)}
  increment = {0, 1}
  counter = 0
  while counter < pos_in_ring
    pos = {pos[0] + increment[0], pos[1] + increment[1]}
    if increment == {0, 1} && pos[1] == (grid_size - 1) / 2
      # upper right corner, change increment to go left
      increment = {-1, 0}
    elsif increment == {-1, 0} && pos[0] == -(grid_size - 1) / 2
      # upper left corner, change increment to go down
      increment = {0, -1}
    elsif increment == {0, -1} && pos[1] == -(grid_size - 1) / 2
      # lower left corner, change increment to go right
      increment = {1, 0}
    end
    counter += 1
  end
  pos[0].abs + pos[1].abs
end

puts solve(1)
puts solve(12)
puts solve(23)
puts solve(1024)

puts solve(368078)
