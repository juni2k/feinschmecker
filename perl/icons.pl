#!/usr/bin/env perl

# [[ icons.pl ]]
# Reads "\n"-separated icon descriptions
# from STDIN. Prints one line where each
# icon description has been mapped to
# an emoji.
# When called with an argument, tries to
# find the argument as a key in $ICONS
# and print its contents.

use strict;
use warnings;
use v5.18.0;

my $ICONS = {
  en => [
    ["\x{1f41f}", 'w/ ghoti', qr/(contains )?fish/i],
    ["\x{24cb}", 'vegan', qr/vegan/i],
    ["\x{1f3cb}\x{fe0f}", 'healthy', qr/mensa vital/i],
    ["\x{1f6ab}\x{1f95b}", 'w/o lactose', qr/lactose/i],
    ["\x{1f437}", 'w/ pork', qr/(contains )?pork/i],
    ["\x{1f37a}", 'w/ alcohol', qr/(contains )?alcohol/i],
    ["\x{1f955}", 'vegetarian', qr/vegetarian/i],
    ["\x{1f404}", 'w/ beef', qr/(contains )?beef/i],
    ["\x{1f98c}", 'w/ wild', qr/(contains )?(wild|venison)/i],
    ["\x{2764}\x{fe0f}", 'favourite food', qr/favorite food/i],
    ["\x{1f414}", 'w/ poultry', qr/(contains )?poultry/i],
    ["\x{1f332}", 'climate dish', qr/climate dish/i],
    ["\x{1f195}", 'new dish', qr/new dish/i]
  ],
  de => [
    ["\x{1f332}", 'Klimateller', qr/klima ?teller/i],
    ["\x{1f98c}", 'Mit Wild', qr/(mit )?wild/i],
    ["\x{1f414}", "Mit Gefl\x{fc}gel", qr/(mit )?gefl\x{fc}gel/i],
    ["\x{1f437}", 'Mit Schwein', qr/(mit )?schwein/i],
    ["\x{1f6ab}\x{1f95b}", 'Laktosefrei', qr/laktose/i],
    ["\x{2764}\x{fe0f}", 'Lieblingsessen', qr/lieblingsessen/i],
    ["\x{1f955}", 'Vegetarisch', qr/vegetarisch/i],
    ["\x{1f404}", 'Mit Rind', qr/(mit )?rind/i],
    ["\x{1f41f}", 'Mit Fisch', qr/(mit )?fisch/i],
    ["\x{24cb}", 'Vegan', qr/vegan/i],
    ["\x{1f3cb}\x{fe0f}", 'Mensa Vital', qr/mensa vital/i],
    ["\x{1f37a}", 'Mit Alkohol', qr/(mit )?alkohol/i],
    ["\x{1f195}", 'Neues Gericht', qr/neu(es gericht)?/i]
  ]
};

sub find_icon {

  my $desc = shift;

  for my $mapping (@{ $ICONS->{en} }, @{ $ICONS->{de} }) {
    my $match = $mapping->[2];

    if ($desc =~ $match) {
      return $mapping->[0];
    }
  }

  return '(no icon)';
}

sub list_icons {

  my $lang = shift;

  unless (exists $ICONS->{$lang}) {
    die "Language $lang not defined, please choose: "
        . join(', ', sort keys %{ $ICONS }) . "\n";
  }

  my @mappings = @{ $ICONS->{$lang} };
  for my $mapping (sort { $a->[1] cmp $b->[1] } @mappings) {
    printf "%s: %s\n", $mapping->[1], $mapping->[0];
  }
}

sub map_icons {
  local $/ = "\n";

  my @alts = <STDIN>;
  chomp for @alts;

  print join q(, ), map { find_icon($_) } @alts;
}

sub main {
  binmode(STDIN,  ':utf8');
  binmode(STDOUT, ':utf8');
  binmode(STDERR, ':utf8');

  if (@ARGV) {
    list_icons($ARGV[0]);
  } else {
    map_icons;
  }
}

main;
