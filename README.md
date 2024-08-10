oneTimeKeys
===========

## Description
An automated mechanism for generating and authenticating one-time keys.
This package will start and generate a queue of some number of keys.
Then some consuming program can `authenticate()` the token by providing
the same, causing the package to return `true` that the token is valid 
before destroying it ...hence the name 'one-time key'.

## Usage
TBD
