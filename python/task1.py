#!/usr/bin/env python
import sys
from os import listdir, mkdir, path, rename
from os.path import isdir, isfile, basename


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

Moves all files in the DIRECTORY to a folder inside DIRECTORY
named as a part ot the file before the first period
""")


def restructure(app_name, dir_name):
	"""
	Moves all files in the dir_name to a folder inside dir_name
	named as a part ot the file before the first period


	:param app_name: the name of the command being executed
	:type app_name: str
	:param dir_name: the directory to be restructured
	:type dir_name: str
	"""

	if not isdir(dir_name):
		print_help(app_name)
		sys.exit(2)

	for filename in listdir(dir_name):
		file = path.join(dir_name, filename)
		if not isfile(file) or '.' not in filename:
			continue
		sub_folder = path.join(dir_name, filename.split('.')[0])
		if not isdir(sub_folder):
			mkdir(sub_folder, 0o755)
		rename(file, path.join(sub_folder, filename))


def main():
	app_name = sys.argv[0]

	if len(sys.argv) != 2:
		print_help(app_name)
		sys.exit(1)

	first_arg = sys.argv[1]
	if first_arg == '--help' or first_arg == '-h':
		print_help(app_name)
		sys.exit(0)

	restructure(app_name, first_arg)


if __name__ == '__main__':
	main()
