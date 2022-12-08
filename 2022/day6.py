def getMakerPos(stream):
	for character in stream:
		if stream.count(character) > 1:
			return False
	return True


def searchPacket(stream, seqLen):
	i = 0
	found = False
	while i < len(stream) and not found:
		found = getMakerPos(stream[i:seqLen+i])
		if found:
			return seqLen+i
		i = i + 1


def puzzle1(f):
	stream = f.readline()
	return searchPacket(stream, 4)


def puzzle2(f):
	stream = f.readline()
	return searchPacket(stream, 14)


def main():
	fname = "./day6.txt"
	with open(fname) as f:
		print(puzzle1(f))
		f.seek(0)
		print(puzzle2(f))


if __name__ == "__main__":
	main()