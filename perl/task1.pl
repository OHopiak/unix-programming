#!/usr/bin/env perl
use warnings;
use strict;
use File::Spec::Functions 'catfile';

sub usage {
    print <<EOF
Usage: $0 DIRECTORY
       $0 [OPTION]
EOF
}

sub help {
    usage;
    print <<EOF
  -h, --help                 give this help list

Moves all files in the DIRECTORY to a folder inside DIRECTORY
named as a part ot the file before the first period
EOF
}

sub convert {
    my $dirname = $_[0];
    unless (-d $dirname) {
        help;
        exit 2;
    }
    opendir my $dir, $dirname or die "Cannot open directory: $!";
    my @files = readdir $dir;
    closedir $dir;

    for (@files) {
        my $file = $_;

        if (!-f catfile($dirname, $file) || index($file, '.') == -1) {
            next;
        }
        my @split = split('\.', $file);
        my $sub_dir = catfile($dirname, $split[0]);
        unless(-e $sub_dir or mkdir($sub_dir, 0755)) {
            die "Unable to create $sub_dir\n";
        }
        rename(catfile($dirname, $file), catfile($sub_dir, $file))
    }
}

my $argc = @ARGV;
if ($argc != 1) {
    help;
    exit 1
}
my $first_arg = $ARGV[0];
if ($first_arg eq "--help" || $first_arg eq "-h") {
    help;
    exit 0;
}

convert $first_arg;