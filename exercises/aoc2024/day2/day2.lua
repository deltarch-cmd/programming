local function shallow_copy(t)
    local copy = {}
    for k, v in pairs(t) do
        copy[k] = v
    end
    return copy
end

local function prepare_data(lines_table)
    local aux_table = {}
    for _, val in ipairs(lines_table) do
        local numbers = {}
        for w in val:gmatch("%w+") do
            table.insert(numbers, tonumber(w))
        end
        table.insert(aux_table, numbers)
    end
    return aux_table
end

local function is_safe(numbers)
    if #numbers < 2 then
        return true
    end

    local asc = false
    local desc = false
    local anterior = numbers[1]

    for i = 2, #numbers, 1 do
        local next = numbers[i]

        if anterior == next then
            return false
        end

        if asc == false and desc == false then
            asc = next > anterior
            desc = not asc
        end

        if (asc and anterior > next) or (desc and next > anterior) then
            return false
        end

        if math.abs(next - anterior) < 1 or math.abs(next - anterior) > 3 then
            return false
        end
        anterior = next
    end
    return true
end

local function part_one(m_table)
    local unsafe_count = 0
    for _, numbers in ipairs(m_table) do
        if not is_safe(numbers) then
            unsafe_count = unsafe_count + 1
        end
    end

    print("Unsafe reports: " .. unsafe_count)
    print("Safe reports: " .. #m_table - unsafe_count)

    return unsafe_count
end

local function part_two(m_table)
    local unsafe_count = 0
    for _, numbers in ipairs(m_table) do
        if is_safe(numbers) then
            goto continue
        end

        local safe_combo = false
        for i = 1, #numbers do
            local aux = shallow_copy(numbers)
            table.remove(aux, i)
            if is_safe(aux) then
                safe_combo = true
                break
            end
        end

        if not safe_combo then
            unsafe_count = unsafe_count + 1
        end
        ::continue::
    end

    print("Unsafe reports: " .. unsafe_count)
    print("Safe reports: " .. #m_table - unsafe_count)

    return unsafe_count
end

local function main()
    local file = io.open("input.txt", "r")
    local lines = {}
    if file ~= nil then
        for line in file:lines() do
            table.insert(lines, line)
        end
        file:close()
    end
    local m_table = prepare_data(lines)
    _ = part_one(m_table)
    _ = part_two(m_table)
end

main()
