import java.util.Random;
import java.util.Arrays;
import java.util.List;
import java.util.ArrayList;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.File;
import java.io.IOException;

public class Sorting {
	

	public static void main(String[] args) {

		boolean verbose = false;
		int size = 1000;
		String fileName = "";
		int[] input = null;

		for (int i = 0; i < args.length; i++) {
			if ("verbose".equalsIgnoreCase(args[i]))
				verbose = true;
			else {
				try {
					size = Integer.parseInt(args[i]);
				} catch (NumberFormatException e) {
					// must be file
					fileName = args[i];
				}
			}
		}

		if (fileName != "") {
			try {
				input = readArrayFromFile(fileName);
			} catch (IOException e) {
				// ignore
			}
		} else {
			input = generateRandomArray(size);
		}
		System.out.println("Sorting " + input.length + " items...");

		int[] newInput = new int[input.length];
		System.arraycopy(input, 0, newInput, 0, input.length);
		if (verbose)
			System.out.println("Input: " + arrayToString(input));
		
		long start = System.currentTimeMillis();
		int[] output = bubbleSort(newInput);
		long elapsed = System.currentTimeMillis() - start;
		// System.out.println("Output: " + arrayToString(output));
		System.out.print(isSorted(output) && compareArrays(output, input) ? "Success" : "Failed");
		System.out.println(" BubbleSort Took " + elapsed + "ms.");

		newInput = new int[input.length];
		System.arraycopy(input, 0, newInput, 0, input.length);
		start = System.currentTimeMillis();
		output = insertionSort(newInput);
		elapsed = System.currentTimeMillis() - start;
		// System.out.println("Output: " + arrayToString(output));
		System.out.print(isSorted(output) && compareArrays(output, input) ? "Success" : "Failed");
		System.out.println(" InsertionSort Took " + elapsed + "ms.");		

		newInput = new int[input.length];
		System.arraycopy(input, 0, newInput, 0, input.length);
		start = System.currentTimeMillis();
		output = selectionSort(newInput);
		elapsed = System.currentTimeMillis() - start;
		// System.out.println("Output: " + arrayToString(output));
		System.out.print(isSorted(output) && compareArrays(output, input) ? "Success" : "Failed");
		System.out.println(" SelectionSort Took " + elapsed + "ms.");		

		newInput = new int[input.length];
		System.arraycopy(input, 0, newInput, 0, input.length);
		start = System.currentTimeMillis();
		output = mergeSort(newInput);
		elapsed = System.currentTimeMillis() - start;
		// System.out.println("Output: " + arrayToString(output));
		System.out.print(isSorted(output) && compareArrays(output, input) ? "Success" : "Failed");
		System.out.println(" MergeSort Took " + elapsed + "ms.");

		newInput = new int[input.length];
		System.arraycopy(input, 0, newInput, 0, input.length);
		start = System.currentTimeMillis();
		Arrays.sort(newInput, 0, newInput.length);
		elapsed = System.currentTimeMillis() - start;
		// System.out.println("Output: " + arrayToString(output));
		System.out.print(isSorted(newInput) && compareArrays(newInput, input) ? "Success" : "Failed");
		System.out.println(" Built-in Took " + elapsed + "ms.");
		
		if (verbose)
			System.out.println("Output: " + arrayToString(newInput));
	}

	public static int[] insertionSort(int[] input) {

		int val = 0;

		for (int pointer = 1; pointer < input.length; pointer++) {
			val = input[pointer];
			for (int i = 0; i < pointer; i++) {
				if (val < input[i]) {
					// shift
					for (int j = pointer; j > i; j--) {
						input[j] = input[j-1];

					}
					input[i] = val;
					break;
				}
			}
		}
		return input;
	}

	public static int[] bubbleSort(int[] input) {

		boolean done = false;
		int holder = 0;
		while (!done) {
			done = true;
			for (int i = 1; i < input.length; i++) {
				if (input[i-1] > input[i]) {
					// swap
					holder = input[i];
					input[i] = input[i-1];
					input[i-1] = holder;
					done = false;
				}
			}
		}
		return input;
	}

	public static int[] selectionSort(int[] input) {

		// int pointer = 0;
		int currentMin;
		int pointerVal;
		for (int i = 0; i < input.length; i++) {
			currentMin = i;
			for (int j = i; j < input.length; j++) {
				if (input[j] < input[currentMin])
					currentMin = j;
			}
			// swap current min with pointer
			pointerVal = input[i];
			input[i] = input[currentMin];
			input[currentMin] = pointerVal;
		} 
		return input;
	}

	public static int[] mergeSort(int[] input) {

		if (input.length < 2)
			return input;

		int split = input.length / 2;
		int[] firstHalf = mergeSort(Arrays.copyOfRange(input, 0, split));
		int[] secondHalf = mergeSort(Arrays.copyOfRange(input, split, input.length));

		return merge(firstHalf, secondHalf);
	}

	private static int[] merge(int[] one, int[] two) {
		int[] result = new int[one.length + two.length];
		int pointerOne = 0;
		int pointerTwo = 0;
		for (int i = 0; i < result.length; i++) {

			if (pointerOne >= one.length) {
				System.arraycopy(two, pointerTwo, result, i, two.length - pointerTwo);
				return result;
			} else if (pointerTwo >= two.length) {
				System.arraycopy(one, pointerOne, result, i, one.length - pointerOne);
				return result;
			} else if (one[pointerOne] < two[pointerTwo]) {
				result[i] = one[pointerOne++];
			} else {
				result[i] = two[pointerTwo++];
			}
		}
		return result;
	}

	// Util Methods

	private static boolean isSorted(int[] input) {
		for (int i=1; i<input.length; i++) {
			if (input[i-1] > input[i])
				return false;
		}
		return true;
	}

	private static boolean compareArrays(int[] inputOne, int[] inputTwo) {
		if (inputOne.length != inputTwo.length)
			return false;

		int[] copyTwo = new int[inputTwo.length];
		int[] used = new int[inputTwo.length];
		System.arraycopy(inputTwo, 0, copyTwo, 0, inputTwo.length);
		for (int i = 0; i < inputOne.length; i++) {
			for (int j = 0; j < copyTwo.length; j++) {
				if (used[j] != 1 && inputOne[i] == copyTwo[j]) {
					used[j] = 1;
					break;
				}
			}
		}

		for (int i = 0; i < used.length; i++) {
			if (used[i] != 1)
				return false;
		}

		return true;
	}

	private static int[] generateRandomArray(int size) {
		Random rand = new Random();
		int[] random = new int[size];
		for (int i = 0; i < size; i++) {
			random[i] = rand.nextInt();
		}
		return random;
	}

	private static String arrayToString(int[] input) {
		StringBuilder buf = new StringBuilder();
		buf.append("[");
		for (int i = 0; i < input.length; i++) {
			buf.append(input[i]);
			if (i != (input.length - 1)) {
				buf.append(", ");
			}
		}
		buf.append("]");
		return buf.toString();
	}

	private static int[] readArrayFromFile(String fileName) throws IOException {
		List<Integer> result = new ArrayList<Integer>();
		File file = new File(fileName);
		BufferedReader reader = new BufferedReader(new FileReader(file));
		String line;
		while ((line = reader.readLine()) != null) {
			result.add(Integer.parseInt(line));
		}
		reader.close();

		int[] arrayResult = new int[result.size()];
		for(int i = 0; i < arrayResult.length; i++)
    		arrayResult[i] = result.get(i);
		
		return arrayResult;
	}
}