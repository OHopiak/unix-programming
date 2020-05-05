package tasks.task2;

import tasks.Common;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.*;

import static java.util.Map.Entry;

public class Main {
	static String USAGE = "" +
			"Usage: <COMMAND> SRC\n" +
			"       <COMMAND> [OPTION]\n";

	static String HELP = USAGE +
			"  -h, --help\t\t\t\t give this help list\n" +
			"\n" +
			"Parses file SRC, changes all letters to lowercase, changes every blank space to the newline,\n" +
			"sorts the result, counts repeated words, sorts the count results descending by count,\n" +
			"show top 10 results.\n";

	static void convert(String filename) throws IOException {
		File file = new File(filename);
		if (!file.isFile()) {
			System.err.println(HELP);
			System.exit(2);
		}

		String[] dataSplit = new String(Files.readAllBytes(Paths.get(filename)))
				.replaceAll("[^a-zA-Z \\n]", "")
				.replaceAll("\\s+", "\n")
				.toLowerCase()
				.split("\\n");
		HashMap<String, Integer> frequencies = new HashMap<>();
		for (String item : dataSplit) {
			frequencies.put(item, frequencies.getOrDefault(item, 0) + 1);
		}

		Comparator<Entry<String, Integer>> valueComparator = (e1, e2) -> {
			Integer v1 = e1.getValue();
			Integer v2 = e2.getValue();
			int intCmp = -v1.compareTo(v2);
			if (intCmp == 0)
				return e1.getKey().compareTo(e2.getKey());
			return intCmp;
		};
		List<Entry<String, Integer>> listOfEntries = new ArrayList<>(frequencies.entrySet());
		listOfEntries.sort(valueComparator);

		for (int i = 0; i < 10; i++) {
			Entry<String, Integer> entry = listOfEntries.get(i);
			System.out.printf("%d %s\n", entry.getValue(), entry.getKey());
		}
	}

	public static void main(String[] args) throws IOException {
		Common.checkArgs(args, HELP);
		convert(args[0]);
	}
}
