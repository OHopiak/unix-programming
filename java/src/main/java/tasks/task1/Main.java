package tasks.task1;

import tasks.Common;

import java.io.File;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Objects;

@SuppressWarnings("ResultOfMethodCallIgnored")
public class Main {
	static String USAGE = "" +
			"Usage: <COMMAND> DIRECTORY\n" +
			"       <COMMAND> [OPTION]\n";

	static String HELP = USAGE +
			"  -h, --help\t\t\t\t give this help list\n" +
			"\n" +
			"Moves all files in the DIRECTORY to a folder inside DIRECTORY\n" +
			"named as a part ot the file before the first period\n";

	static void restructure(String dirName) {
		File dir = new File(dirName);
		if (!dir.isDirectory()) {
			System.err.println(HELP);
			System.exit(2);
		}
		for (String pathname : Objects.requireNonNull(dir.list())) {
			Path filename = Paths.get(dirName, pathname);
			File file = new File(filename.toString());
			if (!file.isFile() || !pathname.contains("."))
				continue;
			String[] split = pathname.split("\\.");
			Path subDirName = Paths.get(dirName, split[0]);
			File subDir = new File(subDirName.toString());
			if(!subDir.isDirectory())
				subDir.mkdir();
			file.renameTo(new File(Paths.get(dirName, split[0], pathname).toString()));
		}

	}

	public static void main(String[] args) {
		Common.checkArgs(args, HELP);
		restructure(args[0]);
	}
}
