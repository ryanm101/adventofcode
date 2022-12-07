import re

def BuildCrates(line):
    crate = []
    for i in range(int(line.strip()[len(line.strip())-1:])):
        crate.append([])
    return crate

def getIndexes(line):
    pos = 0
    offset = 0
    idx = []
    while pos != -1:
        pos = line.find('[', offset)
        offset = pos +1
        if pos != -1:
            idx.append(pos)
    return idx


def GetStartingCrates(crates):
    cr = BuildCrates(crates[len(crates)-1:][0])
    for c in crates:
        idx = getIndexes(c)
        for i in idx:
            cratenum = int(i/4)
            cr[cratenum].append(c[i+1])
    i = 0
    while i < len(cr):
        cr[i].reverse()
        i = i + 1
    return cr

def ParseActions(lines):
    found = False
    actions = []
    for line in lines:
        if line.strip() == "":
            found = True
            continue
        if not found:
            continue
        pat = re.compile('move (\d+) from (\d+) to (\d+)')
        m = re.match(pat, line)
        actions.append(m.groups())
    return actions

def MoveCrates(action, stack):
    for i in range(int(action[0])):
        stack[int(action[2])-1].append(stack[int(action[1])-1].pop())
    return stack

def MoveCrates2(action, stack):
    tarr = stack[int(action[1])-1][len(stack[int(action[1])-1])-int(action[0]):]
    del stack[int(action[1])-1][len(stack[int(action[1])-1])-int(action[0]):]
    stack[int(action[2])-1] = stack[int(action[2])-1] + tarr
    return stack

def decode(stack):
    x = ""
    for i in stack:
        x = x + i.pop()
    return x

def puzzle1(f):
    crates = []
    for line in f.readlines():
        if line == "\n":
            break
        crates.append(line)
    stacks = GetStartingCrates(crates)
    actionlines = []
    f.seek(0)
    for line in f.readlines():
        actionlines.append(line)
    actions = ParseActions(actionlines)
    for action in actions:
        stacks = MoveCrates(action, stacks)
    return decode(stacks)
    

def puzzle2(f):
    crates = []
    for line in f.readlines():
        if line == "\n":
            break
        crates.append(line)
    stacks = GetStartingCrates(crates)
    actionlines = []
    f.seek(0)
    for line in f.readlines():
        actionlines.append(line)
    actions = ParseActions(actionlines)
    for action in actions:
        stacks = MoveCrates2(action, stacks)
    return decode(stacks)

def main():
	fname = "./day5.txt"
	with open(fname) as f:
		print(puzzle1(f))
		f.seek(0)
		print(puzzle2(f))


if __name__ == "__main__":
	main()