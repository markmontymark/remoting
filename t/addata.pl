#!/usr/bin/env perl

use Test::More;
use lib 't/lib-perl';
use GET;
use POST;
use TestUrl;
use DoneTesting;
use JSON::XS;
use Go::FindTypes;

Go::FindTypes::finder( './' );

print JSON::XS::encode_json( $Go::FindTypes::types );
exit;


my $port = shift || '8787';
my $baseurl = "http://localhost:$port/simple-service";

my $tasks = [
	GET->new( 
		url => "$baseurl/items" ),
	POST->new( 
		url => "$baseurl/item",
		query => [
			{Id=>0,AvailableStock => 5},
			{Id=>1,AvailableStock => 6},
			{Id=>2,AvailableStock => 9},
			{Id=>3,AvailableStock => 12},
			{Id=>4,AvailableStock => 2},
			{Id=>5,AvailableStock => 0},
		]),
	POST->new( 
		url => '$baseurl/items',
		query => [
			[
				{Id=>6,AvailableStock => 7},
				{Id=>7,AvailableStock => 56},
				{Id=>8,AvailableStock => 95},
				{Id=>9,AvailableStock => 12},
				{Id=>10,AvailableStock => 25},
				{Id=>11,AvailableStock => 50},
			]
		]),
	TestUrl->new(
		url => "$baseurl/items" ,
		tests => [ sub{
			my($content) = @_;
			my $data = JSON::XS::decode_json( $content );
			print "data $data\n";
			ok $content =~ /Available/sigm;
		}]
	),
	DoneTesting->new()
];

$_->run for @$tasks;
