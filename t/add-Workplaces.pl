
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
		url => "$baseurl/workplaces" ),
	POST->new( 
		url => "$baseurl/workplace",
		query => [
			{Id=>"0",Workplace=> 4},
			{Id=>"1",Workplace=> 3},
			{Id=>"2",Workplace=> 2},
			{Id=>"3",Workplace=> 5},
			{Id=>"4",Workplace=> 1},
		]),

	TestUrl->new(
		url => "$baseurl/workplaces" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 workplaces";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/workplace/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok $data->{Workplace} == 5 && $data->{Id} == 3, "View a workplace";
		}]
		),
	DELETE->new( 
		url => "$baseurl/workplace/3" ),

	TestUrl->new(
		url => "$baseurl/workplaces" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a workplace, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

