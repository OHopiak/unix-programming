#!/usr/bin/env python
import sys
import re
from os.path import isfile


def print_usage(app_name):
	"""
	Prints the usage info

	:param app_name: the name of the command being executed
	:type app_name: str
	"""
	print(f"""\
Usage: {app_name} DIRECTORY
       {app_name} [OPTION]
""")


def print_help(app_name):
	"""
	Prints the detailed help info

	:param app_name: the name of the command being executed
	:type app_name: str
	"""
	print_usage(app_name)
	print(f"""\
  -h, --help				 give this help list

Parses file SRC, changes all letters to lowercase, changes every blank space to the newline,
sorts the result, counts repeated words, sorts the count results descending by count,
show top 10 results.
""")


def convert(app_name, filename):
	"""
	Parses file filename, changes all letters to lowercase, changes every blank space to the newline,
	sorts the result, counts repeated words, sorts the count results descending by count,
	show top 10 results.

	:param app_name: the name of the command being executed
	:type app_name: str
	:param filename: the directory to be restructured
	:type filename: str
	"""

	if not isfile(filename):
		print_help(app_name)
		sys.exit(2)
	with open(filename) as f:
		text = '\n'.join(f.readlines())
	if not text:
		sys.exit(0)
	pattern = re.compile(r'[^a-zA-Z \n]')
	text = pattern.sub('', text)
	pattern = re.compile(r'\s+')
	text = pattern.sub('\n', text)
	text = text.lower()
	frequency = {}
	for word in text.split('\n'):
		frequency[word] = frequency.get(word, 0) + 1

	top_list = [(key, value) for key, value in sorted(frequency.items(), key=lambda pair: (-pair[1], pair[0]))]
	for i in range(min(10, len(top_list))):
		word, count = top_list[i]
		print(f'{count} {word}')


def main():
	app_name = sys.argv[0]

	if len(sys.argv) != 2:
		print_help(app_name)
		sys.exit(1)

	first_arg = sys.argv[1]
	if first_arg == '--help' or first_arg == '-h':
		print_help(app_name)
		sys.exit(0)

	convert(app_name, first_arg)


if __name__ == '__main__':
	main()
