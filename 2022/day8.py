def checkvisible(data, x, y, rows, columns):
	val = data[x][y]

	u = y - 1
	d = y + 1
	l = x - 1
	r = x + 1

	Ublocked = False
	Dblocked = False
	Lblocked = False
	Rblocked = False
	while u >= 0:
		if val <= data[x][u]:
			Ublocked = True
			break
		u = u - 1

	while d < rows:
		if val <= data[x][d]:
			Dblocked = True
			break
		d = d + 1

	while l >= 0:
		if val <= data[l][y]:
			Lblocked = True
			break
		l = l - 1
	
	while r < columns:
		if val <= data[r][y]:
			Rblocked = True
			break
		r = r + 1

	if Ublocked and Dblocked and Lblocked and Rblocked:
		return False
	return True


def iterateTrees(data, rows, columns):
	visible = 0
	c = 1
	while c < (columns - 1):
		r = 1
		while r < (rows - 1):
			if checkvisible(data, c, r, rows, columns):
				visible = visible + 1
			r = r + 1
		c = c + 1
	return visible


def getScenicScore(data, x, y, rows, columns):
	#print("x{}y{}".format(x,y))
	val = data[x][y]
	
	u = y - 1
	d = y + 1
	l = x - 1
	r = x + 1

	us = 0
	ds = 0
	ls = 0
	rs = 0

	# x is c, y is r

	if x != 3 or y != 2:
		return 0
	print(data[x][y])
	print("x=3,y=2")
	print("X={} Y={}".format(x,y))
    # 1 1 1 2
	# 1 1 2 2

	# 2 1 1 1
	# 2 2 1 2
	while u >= 0:
		t = data[y][u]
		print("X={} Y={}".format(x,u))
		print(t)
		if val >= t:
			us = us + 1
			if val == t:
				break
			u = u - 1
			continue
		break

	while l >= 0:
		t = data[y][l]
		if val >= t:
			ls = ls + 1
			if val == t:
				break
			l = l - 1
			continue
		break
		
	while r < columns:
		t = data[y][r]
		if val >= t :
			rs = rs + 1
			if val == t:
				break
			r = r + 1
			continue
		break
	
	while d < rows:
		t = data[x][d]
		if val >= t:
			ds = ds + 1
			if val == t:
				break
			d = d + 1
			continue
		break

	if us == 0:
		us = 1
	if ds == 0:
		ds = 1
	if ls == 0:
		ls = 1
	if rs == 0:
		rs = 1

	print("  U L R D")
	print("5 2 2 1 2")
	print("{} {} {} {} {}".format(val, us, ls, rs, ds)) 
	
	return (us * ds * ls * rs)


def iterateTreesVisDistance(data, rows, columns):
	score = 0
	c = 0
	while c < columns:
		r = 0
		while r < rows:
			ss = getScenicScore(data, c, r, rows, columns)
			if ss > score:
				score = ss
			r = r + 1
		c = c + 1
	return score

def GetGrid(f):
	grid = []
	for line in f:
		grid.append(line.strip())
	return grid, len(grid), len(grid[0])

def puzzle1(f):
	grid, rows, columns = GetGrid(f)
	InsideVisible = iterateTrees(grid, rows, columns)
	OutsideVisible = (2 * columns) + (2 * rows) - 4
	return OutsideVisible + InsideVisible


def puzzle2(f):
	grid, rows, columns = GetGrid(f)
	return iterateTreesVisDistance(grid, rows, columns)


def main():
	fname = "./day8s.txt"
	with open(fname) as f:
		print(puzzle1(f))
		f.seek(0)
		print(puzzle2(f))


if __name__ == "__main__":
	main()