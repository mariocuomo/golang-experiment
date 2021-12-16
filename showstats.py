# Importing packages
import matplotlib.pyplot as plt


fileObject = open("stats.txt", "r")
data = fileObject.readlines()
inputSize = []
bucketSort = []
insertionSort = []
countingSort = []
mergeSort = []
bubbleSort = []
treeSort = []
selectionSort = []

for riga in data:
	pos=riga.find(':')
	if (pos != -1):
		inputSize.append(int(riga[pos+2 : len(riga)]))

	pos=riga.find('bucketSort')
	if (pos != -1):
		bucketSort.append(int(riga[pos+11 : len(riga)]))

	pos=riga.find('insertionSort')
	if (pos != -1):
		insertionSort.append(int(riga[pos+14 : len(riga)]))

	pos=riga.find('countingSort')
	if (pos != -1):
		countingSort.append(int(riga[pos+13 : len(riga)]))

	pos=riga.find('mergeSort')
	if (pos != -1):
		mergeSort.append(int(riga[pos+10 : len(riga)]))

	pos=riga.find('bubbleSort')
	if (pos != -1):
		bubbleSort.append(int(riga[pos+11 : len(riga)]))

	pos=riga.find('treeSort')
	if (pos != -1):
		treeSort.append(int(riga[pos+9 : len(riga)]))

	pos=riga.find('selectionSort')
	if (pos != -1):
		selectionSort.append(int(riga[pos+14 : len(riga)]))


# Define data values
inputSize_sorted = inputSize.copy()
inputSize_sorted.sort()


bucketSort_copy = bucketSort.copy()
insertionSort_copy = insertionSort.copy()
countingSort_copy = countingSort.copy()
mergeSort_copy = mergeSort.copy()
bubbleSort_copy = bubbleSort.copy()
treeSort_copy = treeSort.copy()
selectionSort_copy = selectionSort.copy()



for i in range(len(inputSize_sorted)):
	pos = inputSize.index(inputSize_sorted[i])

	bucketSort[i]=bucketSort_copy[pos]
	insertionSort[i]=insertionSort_copy[pos]
	countingSort[i]=countingSort_copy[pos]
	mergeSort[i]=mergeSort_copy[pos]
	bubbleSort[i]=bubbleSort_copy[pos]
	treeSort[i]=treeSort_copy[pos]
	selectionSort[i]=selectionSort_copy[pos]



# Plot a simple line chart
plt.plot(inputSize_sorted, bucketSort, 'b', label='bucketSort')
plt.plot(inputSize_sorted, insertionSort, 'g', label='insertionSort')
plt.plot(inputSize_sorted, countingSort, 'r', label='countingSort')
plt.plot(inputSize_sorted, mergeSort, 'c', label='mergeSort')
plt.plot(inputSize_sorted, bubbleSort, 'm', label='bubbleSort')
plt.plot(inputSize_sorted, treeSort, 'y', label='treeSort')
plt.plot(inputSize_sorted, selectionSort, 'k', label='selectionSort')
plt.xlabel('# of element in input')
plt.ylabel('time in ms')
plt.legend()
plt.show()