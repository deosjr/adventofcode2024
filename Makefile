go1:
	@go run 01/day01.go

go4:
	@go run 04/day04.go

go6:
	@go run 06/day06.go

go7:
	@go run 07/day07.go

fleng1:
	@fleng 01/day01.pcn -o out
	@./out
	@rm out

fleng2:
	@fleng 02/day02.pcn -o out
	@./out
	@rm out

fleng3:
	@fleng 03/day03.pcn -o out
	@./out
	@rm out

fleng4:
	@fleng 04/day04.pcn -o out
	@./out
	@rm out

fleng5:
	@fleng 05/day05.pcn -o out
	@./out
	@rm out
	
fleng7:
	@fleng 07/day07.pcn -o out
	@./out
	@rm out
