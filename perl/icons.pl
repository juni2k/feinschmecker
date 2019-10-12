#!/usr/bin/env perl

# [[ icons.pl ]]
# Reads "\n"-separated icon descriptions
# from STDIN. Prints one line where each
# icon description has been mapped to
# an emoji.
# When called with an argument, tries to
# find the argument as a key in $icons
# and prints its contents.

use strict;
use warnings;
use v5.18.0;

my $icons = {
  en => [
    ["\x{1f41f}", 'w/ ghoti', qr/contains fish/i],
    ["\x{24cb}", 'vegan', qr/vegan/i],
    ["\x{1f3cb}\x{fe0f}", 'healthy', qr/mensa vital/i],
    ["\x{1f6ab}\x{1f95b}", 'w/o lactose', qr/lactose/i],
    ["\x{1f437}", 'w/ pork', qr/contains pork/i],
    ["\x{1f37a}", 'w/ alcohol', qr/contains alcohol/i],
    ["\x{1f955}", 'vegetarian', qr/vegetarian/i],
    ["\x{1f42e}", 'w/ beef', qr/contains beef/i],
    ["\x{1f98c}", 'w/ wild', qr/contains (wild|venison)/i],
    ["\x{2764}\x{fe0f}", 'favourite food', qr/favorite food/i],
    ["\x{1f414}", 'w/ poultry', qr/contains poultry/i],
    ["\x{1f332}", 'climate plate', qr/climate plate/i],
    ["\x{1f195}", 'new meal', qr/new meal/i]
  ],
  de => [
    ["\x{1f332}", 'Klimateller', qr/klima teller/i],
    ["\x{1f98c}", 'Mit Wild', qr/mit wild/i],
    ["\x{1f414}", "Mit Gefl\x{fc}gel", qr/mit gefl\x{fc}gel/i],
    ["\x{1f437}", 'Mit Schwein', qr/mit schwein/i],
    ["\x{1f6ab}\x{1f95b}", 'Laktosefrei', qr/laktose/i],
    ["\x{2764}\x{fe0f}", 'Lieblingsessen', qr/lieblingsessen/i],
    ["\x{1f955}", 'Vegetarisch', qr/vegetarisch/i],
    ["\x{1f42e}", 'Mit Rind', qr/mit rind/i],
    ["\x{1f41f}", 'Mit Fisch', qr/mit fisch/i],
    ["\x{24cb}", 'Vegan', qr/vegan/i],
    ["\x{1f3cb}\x{fe0f}", 'Mensa Vital', qr/mensa vital/i],
    ["\x{1f37a}", 'Mit Alkohol', qr/mit alkohol/i],
    ["\x{1f195}", 'Neues Gericht', qr/neues gericht/i]
  ]
};

my @combined_icons = (@{ $icons->{en} }, @{ $icons->{de} });

sub find_icon {
  my ($combined, $desc) = @_;

  for my $mapping (@{ $combined }) {
    my $match = $mapping->[2];
    if ($desc =~ $match) {
      return $mapping->[0];
    }
  }

  return '(no icon)';
}

binmode(STDIN,  ':utf8');
binmode(STDOUT, ':utf8');
binmode(STDERR, ':utf8');

if (@ARGV) {
  my $lang = $ARGV[0];

  unless (exists $icons->{$lang}) {
    die "Language $lang not defined, please choose: "
        . join(', ', sort keys %{ $icons }) . "\n";
  }

  my @mappings = @{ $icons->{$lang} };
  for my $mapping (sort { $a->[1] cmp $b->[1] } @mappings) {
    printf "%s: %s\n", $mapping->[1], $mapping->[0];
  }

  exit;
} else {
  local $/ = "\n";
  my @alts = <STDIN>;

  chomp for @alts;

  my @mapped_icons =
      map { find_icon(\@combined_icons, $_) }
  @alts;

  print join ', ', @mapped_icons;
}
