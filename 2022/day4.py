def compareRanges(r1, r2):
    if set(r1).issubset(set(r2)) or set(r2).issubset(set(r1)):
        return True
    return False


def splitToInt(a):
    x = list(map(int,a.split("-")))
    return list(range(x[0], x[1]+ 1))

def CheckForOverlap(r1, r2):
    for r in r1:
        if r in r2:
            return True
    for r in r2:
        if r in r1:
            return True
    return False

def puzzle1(f):
    fullyContained = 0
    for line in f.readlines():
        a,b = line.strip().split(',')
        if compareRanges(splitToInt(a),splitToInt(b)):
            fullyContained = fullyContained + 1
    return fullyContained

def puzzle2(f):
    overlapping = 0
    for line in f.readlines():
        a,b = line.strip().split(',')
        if CheckForOverlap(splitToInt(a),splitToInt(b)):
            overlapping = overlapping + 1
    return overlapping



def main():
	fname = "./day4.txt"
	with open(fname) as f:
		print(puzzle1(f))
		f.seek(0)
		print(puzzle2(f))


if __name__ == "__main__":
	main()