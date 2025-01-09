let rec fibonacci n =
    if n = 0 then 1
    else if n = 1 then 1
    else fibonacci (n - 1) + fibonacci (n - 2)

let () = print_endline(string_of_int(fibonacci 5))
