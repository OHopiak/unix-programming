#!/usr/bin/env perl
use warnings;
use strict;
use Data::Dumper qw(Dumper);

sub usage {
    print <<EOF
Usage: $0 SRC
       $0 [OPTION]
EOF
}

sub help {
    usage;
    print <<EOF
  -h, --help                 give this help list

Parses file SRC, changes all letters to lowercase, changes every blank space to the newline,
sorts the result, counts repeated words, sorts the count results descending by count,
show top 10 results.
EOF
}

sub convert {
    my $filename = $_[0];
    unless (-r $filename) {
        help;
        exit 2;
    }

    open my $file_handle, '<', $filename or die "Can't open file $!";
    read $file_handle, my $text, -s $file_handle;
    $text =~ s/[^a-zA-Z \n]//g;
    $text =~ s/\s+/\n/g;
    $text = lc($text);
    my %frequency;
    for (split('\n', $text)) {
        if (!exists $frequency{$_}) {
            $frequency{$_} = 1;
        }
        else {
            $frequency{$_} += 1;
        }
    }
    my @keys = sort {
        -$frequency{$a} <=> -$frequency{$b}
            or $a cmp $b
    } keys(%frequency);
    my @values = @frequency{@keys};

    for ((0 .. 9)) {
        print("$values[$_] $keys[$_]", "\n");
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