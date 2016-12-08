# PLEASE HELP ME

## WHAT IS THIS
A small script to output a random plead for help.

## WHY WOULD YOU DO THIS
I needed something to spice up my zsh's ```preexec```

## HOW DO I SET THIS UP
Build the .go file, and copy the built file somewhere.
Something like this in your .profile would suffice if you're using zsh
```
preexec () { /path/to/binary/help; }
```
