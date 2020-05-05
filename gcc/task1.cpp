#include <filesystem>
#include <iostream>

namespace fs = std::filesystem;

void usage(const std::string& appName)
{
	std::cout << "Usage: " << appName << " DIRECTORY" << std::endl;
	std::cout << "       " << appName << " [OPTION]" << std::endl;
}

void help(const std::string& appName)
{
	usage(appName);
	std::cout << "  -h, --help                 give this help list" << std::endl;
	std::cout << "Moves all files in the DIRECTORY to a folder inside DIRECTORY" << std::endl;
	std::cout << "named as a part ot the file before the first period." << std::endl;
}

int restructure(const std::string& appName, const fs::path& dirPath)
{
	if (!fs::is_directory(dirPath)) {
		help(appName);
		return 2;
	}
	for (auto&& file : fs::directory_iterator(dirPath))
	{
		auto filename = file.path().filename().string();
		int periodPos = filename.find('.');
		// check if the file is regular and it contains a period
		if(!fs::is_regular_file(file.path()) || periodPos == -1)
			continue;
		auto subFolder = filename.substr(0, periodPos);
		auto newDir = dirPath / subFolder;
		// create a subdirectory with an appropriate format
		fs::create_directory(newDir);
		// move the file to the subdirectory
		fs::rename(file.path(), dirPath / subFolder / filename);
		std::cout << dirPath / subFolder / filename << std::endl;
	}

	return 0;
}

int main(int argc, char** argv)
{
	std::string appName = argv[0];

	if (argc != 2) {
		help(appName);
		return 1;
	}
	std::string firstParam = argv[1];
	if (firstParam == "--help" || firstParam == "-h") {
		help(appName);
		return 0;
	}

	return restructure(appName, firstParam);
}
