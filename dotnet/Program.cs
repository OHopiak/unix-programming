using System;
using System.IO;

namespace task1
{
    class Program
    {
    	static string USAGE = @"\
Usage: <COMMAND> DIRECTORY
       <COMMAND> [OPTION]
";

    	static string HELP = USAGE + @"\
  -h, --help				 give this help list

Moves all files in the DIRECTORY to a folder inside DIRECTORY
named as a part ot the file before the first period
";

		static void restructure(string dirName){
			if (!Directory.Exists(dirName)) {
				Console.Error.WriteLine(HELP);
				Environment.Exit(2);
			}

			foreach (string pathname in Directory.GetFiles(dirName)) {
				string filename = Path.GetFileName(pathname);
				if(!File.Exists(pathname) || !filename.Contains('.'))
					continue;

				string subDirName = Path.Combine(dirName, filename.Split(new Char[]{'.'})[0]);
				if(!Directory.Exists(subDirName))
					Directory.CreateDirectory(subDirName);
				File.Move(pathname, Path.Combine(subDirName, filename));
			}
		}

        static void Main(string[] args)
        {
			if (args.Length != 1) {
				Console.Error.WriteLine(HELP);
				Environment.Exit(1);
			}
			String firstArg = args[0];
			if (firstArg == "--help" || firstArg == "-h") {
				Console.Error.WriteLine(HELP);
				Environment.Exit(0);
			}

			restructure(firstArg);
        }
    }
}
