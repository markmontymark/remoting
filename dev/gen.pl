use strict;
use warnings;

use File::Slurp;


my $api_tmpl = q#
// %%comment%%
list%%collection%%   gorest.EndPoint `method:"GET"    path:"/%%collectionlc%%/"               output:"[]%%object%%"`
view%%object%%    gorest.EndPoint `method:"GET"    path:"/%%objectlc%%/{%%object%%Id:int}"  output:"%%object%%"`
add%%object%%     gorest.EndPoint `method:"POST"   path:"/%%objectlc%%/"                postdata:"%%object%%"`
delete%%object%%  gorest.EndPoint `method:"DELETE" path:"/%%objectlc%%/{%%object%%Id:int}"`
#;

my $test_tmpl = File::Slurp::read_file('dev/add_test.pl');
my $api_impl_tmpl = File::Slurp::read_file('dev/api_impl');
my $collection_tmpl = File::Slurp::read_file('dev/collection.go-tmpl');
my $add_test = File::Slurp::read_file('dev/add_test.pl');

my $simple_service_tmpl = q#

package ui

import (
   "code.google.com/p/gorest"
   //"appengine"
   //"appengine/datastore"
   //"appengine/user"
)

type SimpleService struct{
    //Service level config
    gorest.RestService    `root:"/simple-service/" consumes:"application/json" produces:"application/json"`

	%%api%%
}

%%api_impl%%
#;

my $objects = [ qw/ User Rating Food BusStop Neighborhood Workplace Restaurant MeetingPlace Route / ];

my $api_endpoints = [];
my $api_impls = [];
for (@$objects)
{
	my %tokens = (
		comment => "For the $_ type",
		object => $_,
		objectlc => lc $_,	
		collection => $_ . 's',
		collectionlc => 	lc($_.'s') );

	my $test_script = $test_tmpl;
	my $apitmpl = $api_tmpl;
	my $apiimpltmpl = $api_impl_tmpl;
	my $collection = $collection_tmpl;
	my($k,$v);
	while( ($k,$v) = each %tokens)
	{
		$apitmpl =~ s/%%$k%%/$v/g;
		$apiimpltmpl =~ s/%%$k%%/$v/g;
		$collection =~ s/%%$k%%/$v/g;
		$test_script =~ s/%%$k%%/$v/g;
	}
	push @$api_endpoints, $apitmpl;
	push @$api_impls, $apiimpltmpl;
	File::Slurp::write_file( "ui/$tokens{collection}.go", $collection );
	File::Slurp::write_file( "t/add-$tokens{collection}.pl", $test_script );
	##print $collection;
}

$simple_service_tmpl =~ s/%%api%%/join "\n\n",@$api_endpoints,''/es;
$simple_service_tmpl =~ s/%%api_impl%%/join "\n\n",@$api_impls,''/es;

File::Slurp::write_file('ui/SimpleService.go' ,$simple_service_tmpl);
