package tasks;

public class Common {
	public static void checkArgs(String[] args, String help){
		if (args.length != 1) {
			System.err.println(help);
			System.exit(1);
		}
		String firstArg = args[0];
		if (firstArg.equals("--help") || firstArg.equals("-h")) {
			System.err.println(help);
			System.exit(0);
		}
	}
}
