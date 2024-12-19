:- use_module(library(clpfd)).
:- use_module(library(dcg/basics)).

:- dynamic mem/2.
mem([], 1).

run :-
    phrase_from_file(parse(Patterns, Designs), '19/day19.input'),
    foldl(p1(Patterns), Designs, 0, Ans1),
    format("Part 1: ~w~n", [Ans1]),
    foldl(p2(Patterns), Designs, 0, Ans2),
    format("Part 2: ~w~n", [Ans2]).

p1(Patterns, Cur, Acc, Next) :-
    ( phrase(possible(Patterns), Cur) -> Next #= Acc+1 ; Next #= Acc ).

p2(Patterns, Cur, Acc, Next) :-
    design_options(Patterns, Cur, Options),
    Next #= Acc + Options.

design_options(Patterns, Design, N) :-
    ( mem(Design, X) -> N #= X ;
    findall(S, (member(P,Patterns), append(P, S, Design)), Suffixes),
    ( Suffixes == [] -> N #= 0 ;
    maplist(design_options(Patterns), Suffixes, Nums),
    sum(Nums, #=, N),
    assertz(mem(Design, N))
    )).

possible(_) --> [].
possible(Patterns) --> {member(X, Patterns)}, X, possible(Patterns).

parse(Patterns, Designs) --> parse_patterns(Patterns), "\n", parse_designs(Designs).

parse_patterns([H|T]) --> string_without(",", H), ", ", parse_patterns(T).
parse_patterns([H]) --> string_without("\n", H), "\n".

parse_designs([H|T]) --> string_without("\n", H), "\n", parse_designs(T).
parse_designs([]) --> eol.
