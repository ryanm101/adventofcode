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
	print(sum(int(v) for v in sorted(elves, reverse=True)[:3]))

if __name__ == "__main__":
    main()
