

def main():
	fname = "./day1_1.txt"
	elves = []
	i = 0
	with open(fname) as f:
		elves.append(0)
		for line in f.readlines():
			if line == "\n":
				i = i + 1
				elves.append(0)
				continue
			elves[i] += int(line)
	print(max(elves))
	top3 = sorted(elves, reverse=True)[:3]
	z = 0
	for i in top3:
		z += i
	print(z)

if __name__ == "__main__":
    main()