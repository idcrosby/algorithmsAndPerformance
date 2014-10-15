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
				# print("before: "),
				# printList(input)
				input[j+1:i+1] = input[j:i]
				input[j] = v
				# print("after: "),
				# printList(input)
				break

def bubbleSort(input):

	done = False
	while not done:
		done = True
		for i, v, in enumerate(input):
			if i+1 < len(input) and v > input[i+1]:
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
		output.append(random.randint(-100,100))
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
if isSorted(li) and listsEqual(li, original):
	print "SelectionSort Success"
else:
	print "SelectionSort Failed"
	print printList(li)

li = original[:]
print "InsertionSort. Before: Sorted? " + `isSorted(li)`
insertionSort(li)
if isSorted(li) and listsEqual(li, original):
	print "insertionSort Success"
else:
	print "insertionSort Failed"
	print printList(li)

li = original[:]
print "BubbleSort. Before: Sorted? " + `isSorted(li)`
bubbleSort(li)
if isSorted(li) and listsEqual(li, original):
	print "BubbleSort Success"
else:
	print "bubbleSort Failed"
	print printList(li)

printList(li)