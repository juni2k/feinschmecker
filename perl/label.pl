#!/usr/bin/env perl

# [[ label.pl ]]
# Reads a dish label from STDIN.
# Then, it cleans up some pieces of
# text and prints the result to STDOUT.

use strict;
use warnings;
use v5.18.0;

binmode(STDIN,  ':utf8');
binmode(STDOUT, ':utf8');
binmode(STDERR, ':utf8');

$/ = undef;
my $label = <STDIN>;

# Remove '#allinONEPOT.....' bullshit
$label =~ s/#allinONEPOT\.*,?//g;

# Remove leading / trailing whitespace
$label =~ s/(^\s*|\s*$)//g;

# Remove annoying allergens
$label =~ s/\s*\(.*?\)\s*//g;

# Fix wrong spacing around commas
$label =~ s/\s*?,\s*/, /g;

# Replace n-whitespace with a single space
$label =~ s/\s+/ /g;

# Always start with an uppercase letter
$label = ucfirst $label;

print $label;
