class Point
    attr_accessor :x, :y
    def initialize(x, y)
        self.x = x
        self.y = y
    end

    def to_s
        "(#{self.x}, #{self.y})"
    end
end

class Vector
    attr_accessor :direction, :distance
    def initialize(direction, distance)
        self.direction = direction
        self.distance = distance
    end
end

def parse_direction(dir)
    Vector.new(dir[0], dir[1..-1].to_i)
end

def draw_coords(coords)
    cursor = Point.new(0, 0)
    grid = []
    coords.each do |coord|
        i = coord.distance
        puts "Direction: #{coord.direction}"
        if coord.direction == "R"
            puts "Moving right"
            n = 0
            i.times do
                grid << Point.new(x, cursor.y)
                cursor = Point.new(x, cursor.y)
            end
        elsif coord.direction == "U"
            puts "Moving up"
            ((cursor.y - i)...cursor.y).each do |y|
                grid << Point.new(cursor.x, y)
                cursor = Point.new(cursor.x, y)
            end
        elsif coord.direction == "L"
            puts "Moving left"
            (cursor.x...(cursor.x-i)).each do |x|
                grid << Point.new(x, cursor.y)
                cursor = Point.new(x, cursor.y)
            end
        elsif coord.direction == "D"
            puts "Moving down"
            (cursor.y...cursor.y + i).each do |y|
                grid << Point.new(cursor.x, y)
                cursor = Point.new(cursor.x, y)
            end
        end
    end
    puts grid
end

def main
    input1 = "R8,U5,L5,D3".split(",").freeze
    result = input1.map { |d| parse_direction(d) }
    draw_coords(result)
end

main