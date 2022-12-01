

def main():
	fname = "./day1_1.txt"
	elves = [0]
	i = 0
	with open(fname) as f:
		for line in f.readlines():
			if line == "\n":
				i = i + 1
				elves.append(0)
				continue
			elves[i] += int(line)
	print(max(elves))
	z = 0
	for i in sorted(elves, reverse=True)[:3]:
		z += i
	print(z)

if __name__ == "__main__":
    main()
