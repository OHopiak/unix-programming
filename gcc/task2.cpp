#include <filesystem>
#include <iostream>
#include <fstream>
#include <regex>
#include <map>
#include <vector>

namespace fs = std::filesystem;

void usage(const std::string& appName)
{
	std::cout << "Usage: " << appName << " SRC" << std::endl;
	std::cout << "       " << appName << " [OPTION]" << std::endl;
}

void help(const std::string& appName)
{
	usage(appName);
	std::cout << "  -h, --help                 give this help list" << std::endl;
	std::cout << "Parses file SRC, changes all letters to lowercase, changes every blank space to the newline,"
			  << std::endl;
	std::cout << "sorts the result, counts repeated words, sorts the count results descending by count," << std::endl;
	std::cout << "show top 10 results." << std::endl;
}

int processFile(const std::string& appName, const std::string& dir)
{
	fs::path filePath = dir;
	// check if the file exists
	if (!fs::is_regular_file(filePath)) {
		help(appName);
		return 2;
	}

	// read file to the string
	std::ifstream f{filePath};
	const auto fileSize = fs::file_size(filePath);
	std::string text(fileSize, ' ');
	f.read(text.data(), fileSize);

	// remove irrelevant characters
	text = std::regex_replace(text, std::regex("[^a-zA-Z \\n]"), "");
	// change any number of blank spaces to newlines
	text = std::regex_replace(text, std::regex("\\s+"), "\n");
	// convert the text to lowercase
	std::transform(text.begin(), text.end(), text.begin(), [](unsigned char c) { return std::tolower(c); });

	// count the word frequencies
	size_t pos = 0;
	std::map<std::string, int> frequencyMap;
	while ((pos = text.find('\n')) != std::string::npos) {
		std::string token = text.substr(0, pos);
		frequencyMap[token] += 1;
		text.erase(0, pos + 1);
	}

	// define the types for sorting, just to keep the code clean
	typedef std::pair<std::string, int> StrInt;
	typedef std::function<bool(StrInt, StrInt)> Comparator;

	// move the values to a sortable type
	std::vector<StrInt> frequencyVector;
	for (auto&& pair : frequencyMap) {
		frequencyVector.emplace_back(pair.first, pair.second);
	}

	// first sort by value, if it is the same, sort by the key
	Comparator descendingOrder = [](const StrInt& elem1, const StrInt& elem2) {
		if (elem1.second > elem2.second)
			return true;
		return elem1.second == elem2.second && elem1.first < elem2.first;
	};

	// sort the vector
	std::sort(frequencyVector.begin(), frequencyVector.end(), descendingOrder);

	// print the top 10 items by count
	int iter = 0;
	for (auto&& pair : frequencyVector) {
		std::cout << pair.second << " " << pair.first << std::endl;
		iter++;
		if (iter == 10) break;
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
	return processFile(appName, firstParam);
}