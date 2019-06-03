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
  'en' => {
    'Contains fish'    => "\x{1f41f}",
    'Vegan'            => "\x{24cb}",
    'Mensa vital'      => "\x{1f3cb}\x{fe0f}",
    'Lactose-free'     => "\x{1f6ab}\x{1f95b}",
    'Contains pork'    => "\x{1f437}",
    'Contains alcohol' => "\x{1f37a}",
    'Vegetarian'       => "\x{1f955}",
    'Contains beef'    => "\x{1f42e}",
    'Contains wild'    => "\x{1f98c}",
    'Contains venison' => "\x{1f98c}",
    'Favorite Food'    => "\x{2764}\x{fe0f}",
    'Contains poultry' => "\x{1f414}",
    'Climate plate'    => "\x{1f332}",
    'New meal'         => "\x{1f195}"
  },
  'de' => {
    'Klima Teller'      => "\x{1f332}",
    'Mit Wild'          => "\x{1f98c}",
    "Mit Gefl\x{fc}gel" => "\x{1f414}",
    'Mit Schwein'       => "\x{1f437}",
    'Laktosefrei'       => "\x{1f6ab}\x{1f95b}",
    'Lieblingsessen'    => "\x{2764}\x{fe0f}",
    'Vegetarisch'       => "\x{1f955}",
    'Mit Rind'          => "\x{1f42e}",
    'Mit Fisch'         => "\x{1f41f}",
    'Vegan'             => "\x{24cb}",
    'Mensa Vital'       => "\x{1f3cb}\x{fe0f}",
    'Mit Alkohol'       => "\x{1f37a}",
    'Neues Gericht'     => "\x{1f195}"
  }
};

my %combined_icons = (%{ $icons->{en} }, %{ $icons->{de} });

sub find_icon {
  my ($combined, $alt) = @_;

  for my $desc (keys %{ $combined }) {
    if ($alt =~ /$desc/i) {
      return $combined->{$desc};
    }
  }

  return "[no icon]";
}

binmode(STDIN,  ':utf8');
binmode(STDOUT, ':utf8');
binmode(STDERR, ':utf8');

$/ = "\n";
my @alts = <STDIN>;

chomp for @alts;

my @mapped_icons =
    map { find_icon(\%combined_icons, $_) }
    @alts;

print join ', ', @mapped_icons;
