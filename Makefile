go1:
	@go run 01/day01.go

fleng1:
	@fleng 01/day01.pcn -o out
	@./out
	@rm out

fleng2:
	@fleng 02/day02.pcn -o out
	@./out
	@rm out
