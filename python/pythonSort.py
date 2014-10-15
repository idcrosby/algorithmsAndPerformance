import random

def selectionSort(input):
	for i, v in enumerate(input):
		min = i
		for idx, val in enumerate(input[i:]):
			if val < input[min]:
				min = (idx + i)
		# swap
		hold = input[min]
		input[min] = input[i]
		input[i] = hold

def insertionSort(input):

	for i, v in enumerate(input):
		for j, v2 in enumerate(input[:i]):
			if v < v2:
				# shift and insert
				input[j+1:i+1] = input[j:i]
				input[j] = v2
				break

def bubbleSort(input):

	done = False
	while not done:
		done = True
		for i, v, in enumerate(input[:-1]):
			if v > input[i+1]:
				# swap
				input[i] = input[i+1]
				input[i+1] = v
				done = False


# Util Methods

def isSorted(input):
	first = input[0]
	for i in input:
		if first > i:
			return False
		first = i
	return True

def listsEqual(list1, list2):
	list2Copy = list2[:]
	for i in list1:
		try:
			list2Copy.remove(i)
		except ValueError:
			return False

	return len(list2Copy) == 0

def generateArray(size):
	output = []
	for x in range(0, size):
		output.append(random.randint(-10,10))
	return output

def printList(input):
	print("["),
	for val in input:
		print(val),
		print(','),
	print("]")


original = generateArray(20)
li = original[:]
printList(original)
print "SelectionSort. Before: Sorted? " + `isSorted(li)`
selectionSort(li)
print "SelectionSort. After: Sorted? " + `isSorted(li)`
print "Same list? " + `listsEqual(original, li)`
# printList(sorted)
printList(li)

li = original[:]
print "InsertionSort. Before: Sorted? " + `isSorted(li)`
insertionSort(li)
print "InsertionSort. After: Sorted? " + `isSorted(li)`
print "Same list? " + `listsEqual(li, original)`
printList(li)

li = original[:]
print "BubbleSort. Before: Sorted? " + `isSorted(li)`
bubbleSort(li)
print "BubbleSort. After: Sorted? " + `isSorted(li)`
print "Same list? " + `listsEqual(li, original)`
printList(li)