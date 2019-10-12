#!/usr/bin/env perl

# [[ icons.pl ]]
# Reads "\n"-separated icon descriptions
# from STDIN. Prints one line where each
# icon description has been mapped to
# an emoji.

use strict;
use warnings;
use v5.18.0;

my $icons = {
  en => [
    ["\x{1f41f}", 'Contains Fish', ['Contains fish']],
    ["\x{24cb}", 'Vegan', ['Vegan']],
    ["\x{1f3cb}\x{fe0f}", 'Mensa vital', ['Mensa vital']],
    ["\x{1f6ab}\x{1f95b}", 'Lactose-free', ['Lactose-free']],
    ["\x{1f437}", 'Contains pork', ['Contains pork']],
    ["\x{1f37a}", 'Contains alcohol', ['Contains alcohol']],
    ["\x{1f955}", 'Vegetarian', ['Vegetarian']],
    ["\x{1f42e}", 'Contains beef', ['Contains beef']],
    ["\x{1f98c}", 'Contains wild', ['Contains wild', 'Contains venison']],
    ["\x{2764}\x{fe0f}", 'Favorite Food', ['Favorite Food']],
    ["\x{1f414}", 'Contains poultry', ['Contains poultry']],
    ["\x{1f332}", 'Climate plate', ['Climate plate']],
    ["\x{1f195}", 'New meal', ['New meal']]
  ],
  de => [
    ["\x{1f332}", 'Klima Teller', ['Klima Teller']],
    ["\x{1f98c}", 'Mit Wild', ['Mit Wild']],
    ["\x{1f414}", "Mit Gefl\x{fc}gel", ["Mit Gefl\x{fc}gel"]],
    ["\x{1f437}", 'Mit Schwein', ['Mit Schwein']],
    ["\x{1f6ab}\x{1f95b}", 'Laktosefrei', ['Laktosefrei']],
    ["\x{2764}\x{fe0f}", 'Lieblingsessen', ['Lieblingsessen']],
    ["\x{1f955}", 'Vegetarisch', ['Vegetarisch']],
    ["\x{1f42e}", 'Mit Rind', ['Mit Rind']],
    ["\x{1f41f}", 'Mit Fisch', ['Mit Fisch']],
    ["\x{24cb}", 'Vegan', ['Vegan']],
    ["\x{1f3cb}\x{fe0f}", 'Mensa Vital', ['Mensa Vital']],
    ["\x{1f37a}", 'Mit Alkohol', ['Mit Alkohol']],
    ["\x{1f195}", 'Neues Gericht', ['Neues Gericht']]
  ]
};

my @combined_icons = (@{ $icons->{en} }, @{ $icons->{de} });

sub find_icon {
  my ($combined, $desc) = @_;

  for my $mapping (@{ $combined }) {
    for my $match (@{ $mapping->[2] }) {
      if ($desc =~ /$match/i) {
        return $mapping->[0];
      }
    }
  }

  return '(no icon)';
}

binmode(STDIN,  ':utf8');
binmode(STDOUT, ':utf8');
binmode(STDERR, ':utf8');

$/ = "\n";
my @alts = <STDIN>;

chomp for @alts;

my @mapped_icons =
    map { find_icon(\@combined_icons, $_) }
    @alts;

print join ', ', @mapped_icons;
