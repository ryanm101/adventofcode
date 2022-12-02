ROCK = 1
PAPER = 2 
SCISSORS = 3
WIN = 6
LOSE = 0
DRAW = 3

shape = {
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
	"A": "X",
	"B": "Y",
	"C": "Z"
}

outcome = {
	"A": {
		"X": DRAW,
		"Y": WIN,
		"Z": LOSE,
		"lose": "Z",
		"draw": "X",
		"win": "Y"
	},
	"B":
	{
		"X": LOSE,
		"Y": DRAW,
		"Z": WIN,
		"lose": "X",
		"draw": "Y",
		"win": "Z"
	},
	"C":
	{
		"X": WIN,
		"Y": LOSE,
		"Z": DRAW,
		"lose": "Y",
		"draw": "Z",
		"win": "X"
	}
}

def result(a, b):
	outcome[a][b] + shape[b]
	return int(outcome[a][b] + shape[b])

def strategy(b):
	if b == "X":
		return "lose"
	elif b == "Y":
		return "draw"
	elif b == "Z":
		return "win"

def puzzle1(f):
	acc = 0
	for line in f.readlines():
		acc += result(line[0], line[2])
	return acc

def	puzzle2(f):
	acc = 0
	for line in f.readlines():
		strat = outcome[line[0]][strategy(line[2])]
		acc += result(line[0], strat)
	return acc

def main():
	fname = "./day2.txt"
	with open(fname) as f:
		print(puzzle1(f))
		f.seek(0)
		print(puzzle2(f))


if __name__ == "__main__":
	main()