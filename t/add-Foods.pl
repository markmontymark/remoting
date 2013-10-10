#!/usr/bin/env perl

use Test::More;
use lib 't/lib-perl';
use GET;
use POST;
use DELETE;
use TestUrl;
use DoneTesting;
use JSON::XS;

##use Go::FindTypes;
##Go::FindTypes::finder( './' );

my $port = shift || '8080';
my $baseurl = "http://localhost:$port/simple-service";

my $tasks = [

	GET->new( 
		url => "$baseurl/foods" ),
	POST->new( 
		url => "$baseurl/food",
		query => [
			{Id=>-1,Name=> "Cupcakes"},
			{Id=>-1,Name=> "Cookies"},
			{Id=>-1,Name=> "Salad"},
			{Id=>-1,Name=> "Croissants"},
			{Id=>-1,Name=> "Bagels"},
		]),

	TestUrl->new(
		url => "$baseurl/foods" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			unless(is scalar @$data,5, "Added 5 foods")
			{
				print "data is ",@$data,"\n";
			}
		}]
		),

	TestUrl->new( 
		url => "$baseurl/food/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is $data->{Name},'Croissants', "View a food, validate Name";
			is $data->{Id},3, "View a food, validate Id";
		}]
		),
	DELETE->new( 
		url => "$baseurl/food/3" ),

	TestUrl->new(
		url => "$baseurl/foods" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a food, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

