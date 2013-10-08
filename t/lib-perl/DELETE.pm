use HttpTask;

package DELETE;
use Mo qw'build default builder coerce is required';
extends 'HttpTask';
has method => (default => sub{'DELETE'});

1;
