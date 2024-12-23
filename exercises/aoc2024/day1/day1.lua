local function create_tables(input)
    local izq = true
    local left_table = {}
    local right_table = {}
    for word in input:gmatch("%w+") do
        if izq then
            table.insert(left_table, word)
        else
            table.insert(right_table, word)
        end
        izq = not izq
    end

    table.sort(left_table, function(a, b) return a < b end)
    table.sort(right_table, function(a, b) return a < b end)

    return left_table, right_table
end

local function part_one(left, right)
    local acum = 0
    for i = 1, #left, 1 do
        acum = acum + math.abs(left[i] - right[i])
    end
    print(acum)
end

local function part_two(left, right)
    local acum = 0
    for i = 1, #left, 1 do
        local rep = 0
        local x = left[i]
        for _, val in ipairs(right) do
            if x == val then
                rep= rep + 1
            end
        end
        acum = acum + (x * rep)
    end
    print(acum)
end

local function main()
    local file = io.open("input.txt", "r")
    local content = nil
    if file ~= nil then
        content = file:read("*all")
        file:close()
    end
    local left, right = create_tables(content)

    part_one(left, right)
    part_two(left, right)
end
main()
