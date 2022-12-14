open Core
let content = In_channel.read_all "input-2"

let split_groups = Str.split (Str.regexp "\n\n")
let split_lines = Str.split (Str.regexp "\n")
let list_of_ints = List.map ~f:int_of_string
let sum = List.fold_left ~f:(+) ~init:0
let invert f a b = f b a
let take_inverted = invert List.take

let () =
  let result = content
    |> split_groups
    |> List.map ~f:split_lines
    |> List.map ~f:list_of_ints
    |> List.map ~f:sum
    |> List.sort ~compare:Int.compare
    |> List.rev
    |> take_inverted 3
    |> sum
in
  print_endline (string_of_int result)





