
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
		url => "$baseurl/users" ),
	POST->new( 
		url => "$baseurl/user",
		query => [
			{Id=>"0",User=> 4},
			{Id=>"1",User=> 3},
			{Id=>"2",User=> 2},
			{Id=>"3",User=> 5},
			{Id=>"4",User=> 1},
		]),

	TestUrl->new(
		url => "$baseurl/users" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 users";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/user/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok $data->{User} == 5 && $data->{Id} == 3, "View a user";
		}]
		),
	DELETE->new( 
		url => "$baseurl/user/3" ),

	TestUrl->new(
		url => "$baseurl/users" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a user, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

