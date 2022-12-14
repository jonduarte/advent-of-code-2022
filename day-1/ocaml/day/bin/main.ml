open Core
let content = In_channel.read_all "input-2"

let split_groups = Str.split (Str.regexp "\n\n")
let split_lines = Str.split (Str.regexp "\n")
let list_of_ints = List.map ~f:int_of_string
let sum = List.fold_left ~f:(+) ~init:0
let (>>) f g x = g(f(x))
let total_calories = split_lines >> list_of_ints >> sum

let () =
  let result = content
    |> split_groups
    |> List.map ~f:total_calories
    |> List.sort ~compare:Int.compare
    |> List.rev
    |> Fn.flip (List.take) 3
    |> sum
in
  print_endline (string_of_int result)





