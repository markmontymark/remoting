
use HttpTask;
use JSON::XS;
use GET;

package TestJSON;

use Mo qw'build default coerce';  #builder coerce is required';
extends 'HttpTask';
has tests =>  (default => sub{[]});
has url =>    (default => sub{sub{@_}});
has content => (default => sub{[]});
has decode => (default => sub{\&JSON::XS::decode_json});

sub run
{
	my $self = shift;
	my $content = $self->content->();
	$_->($self->decode->($content)) for @{ $self->tests};
}

1;
