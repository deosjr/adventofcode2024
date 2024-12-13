:- use_module(library(clpfd)).
:- use_module(library(dcg/basics)).

run :-
    phrase_from_file(parse(List), '13/day13.input'),
    convlist(p1, List, Out1),
    sum(Out1, #=, Ans1),
    format("Part 1: ~w~n", [Ans1]),
    convlist(p2, List, Out2),
    sum(Out2, #=, Ans2),
    format("Part 2: ~w~n", [Ans2]).

p1(machine(XA/YA, XB/YB, XP/YP), Cost) :-
    XP #= A*XA + B*XB,
    YP #= A*YA + B*YB,
    [A,B] ins 0..100,
    label([A,B]),
    Cost #= 3*A + B.

p2(machine(XA/YA, XB/YB, XP/YP), Cost) :-
    XP + 10000000000000 #= A*XA + B*XB,
    YP + 10000000000000 #= A*YA + B*YB,
    [A,B] ins 0..10000000000000,
    label([A,B]),
    Cost #= 3*A + B.

parse([machine(A,B,P)|T]) --> parse_buttonA(A), parse_buttonB(B), parse_prize(P), "\n", parse(T).
parse([machine(A,B,P)]) --> parse_buttonA(A), parse_buttonB(B), parse_prize(P).

parse_buttonA(X/Y) --> "Button A: X+", integer(X), ", Y+", integer(Y), "\n".
parse_buttonB(X/Y) --> "Button B: X+", integer(X), ", Y+", integer(Y), "\n".
parse_prize(X/Y) --> "Prize: X=", integer(X), ", Y=", integer(Y), "\n".
