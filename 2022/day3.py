def atm(a):
    val = ord(a)
    if val >= 65 and val <= 90:
        return (val - 64) + 26
    if val >= 97 and val <= 122:
        return val - 96
    return None

def splitCompart(a):
    return a[:len(a)//2], a[len(a)//2:]

def getDups(a,b):
    x = [l for l in a if l in b]
    return set(x)

def getGroups(f):
    group = []
    tgroup = []
    for count, line in enumerate(f):
        tgroup.append(line.strip())
        if (count+1) %3 == 0 and count != 0:
            group.append(tgroup)
            tgroup = []
    return group

def puzzle1(f):
    acc = 0
    for line in f.readlines():
        for i in getDups(*splitCompart(line)):
            acc += atm(i)
    return acc

def puzzle2(f):
    acc = 0
    grps = getGroups(f)
    for grp in grps: 
        for i in getDups(getDups(grp[0], grp[1]), grp[2]):
            acc += atm(i)
    return acc

def main():
	fname = "./day3.txt"
	with open(fname) as f:
		print(puzzle1(f))
		f.seek(0)
		print(puzzle2(f))


if __name__ == "__main__":
	main()