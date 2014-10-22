import random
from datetime import datetime

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

def mergeSort(input):
	if len(input) < 2:
		return input

	split = len(input) / 2

	return merge(mergeSort(input[:split]), mergeSort(input[split:]))

def merge(first, second):

	size = len(first) + len(second)
	result = []

	firstPointer = 0
	secondPointer = 0
	for x in range(0, size):
		if firstPointer == len(first):
			for el in second[secondPointer:]:
				result.append(el)
			break
		elif secondPointer == len(second):
			for el in first[firstPointer:]:
				result.append(el)
			break
		elif first[firstPointer] < second[secondPointer]:
			result.append(first[firstPointer])
			firstPointer += 1
		else:
			result.append(second[secondPointer])
			secondPointer += 1
	return result

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
	for i, val in enumerate(input):
		print(val),
		if i < len(input) -1:
			print(','),
	print("]")

def toReadable(timeDelta):
	timeString = ""
	hours, remainder = divmod(timeDelta.seconds, 3600)
	minutes, seconds = divmod(remainder, 60)
	milliseconds = timeDelta.microseconds // 1000
	if hours > 0:
		timeString += str(hour) + ":" + str(minutes) + ":"
	elif minutes > 0:
		timeString += str(minutes) + ":"
	
	milliString = str(milliseconds)
	if milliseconds < 10:
		milliString = "00" + str(milliseconds)
	elif milliseconds < 100:
		milliString = "0" + str(milliseconds)

	timeString += str(seconds) + "." + milliString + " seconds"
	return timeString

original = generateArray(10000)
li = original[:]
# printList(original)
start = datetime.now()
selectionSort(li)
runTime = datetime.now() - start
if isSorted(li) and listsEqual(li, original):
	print "SelectionSort Success. Ran in " + toReadable(runTime)
else:
	print "SelectionSort Failed"
	print printList(li)

li = original[:]
start = datetime.now()
insertionSort(li)
runTime = datetime.now() - start
if isSorted(li) and listsEqual(li, original):
	print "InsertionSort Success. Ran in " + toReadable(runTime)
else:
	print "insertionSort Failed"
	print printList(li)

li = original[:]
start = datetime.now()
bubbleSort(li)
runTime = datetime.now() - start
if isSorted(li) and listsEqual(li, original) and not isSorted(original):
	print "BubbleSort Success. Ran in " + toReadable(runTime)
else:
	print "bubbleSort Failed"
	print printList(li)

li = original[:]
start = datetime.now()
result = mergeSort(li)
runTime = datetime.now() - start
if isSorted(result) and listsEqual(result, original):
	print "MergeSort Success. Ran in " + toReadable(runTime)
else:
	print "MergeSort Failed"
	print printList(result)

li = original[:]
start = datetime.now()
li.sort()
runTime = datetime.now() - start
if isSorted(li) and listsEqual(li, original):
	print "Built-in Sort Success. Ran in " + toReadable(runTime)
else:
	print "MergeSort Failed"
	print printList(li)
# printList(result)