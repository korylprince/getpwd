/*
getpwd is a thin wrapper around getpwnam_r and getpwuid_r.
I wrote this package because the syscall package doesn't have any way to map a uid to a username.
At this point the API shouldn't change, but I won't make any promises yet.
*/
package getpwd
