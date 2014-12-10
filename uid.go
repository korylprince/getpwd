package getpwd

/*
#include <stdlib.h>
#include <errno.h>
#include <unistd.h>
#include <pwd.h>

struct passwd * uidtopwd(uid_t uid, char **buf) {

	size_t bufsize = sysconf(_SC_GETPW_R_SIZE_MAX);
	if (bufsize == -1) {
		//should be large enough
		bufsize = 16384;
	}

	*buf = malloc(bufsize);
	if (*buf == NULL) {
		errno = ENOMEM;
		return NULL;
	}

	struct passwd *pwd;
	struct passwd *result;

	pwd = malloc(sizeof(struct passwd));
	if (pwd == NULL) {
		errno = ENOMEM;
		return NULL;
	}

	int ret = getpwuid_r(uid, pwd, *buf, bufsize, &result);
	if (result == NULL) {
		if (ret != 0) {
			errno = ret;
		}
		return NULL;
	}
	return result;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//NewPasswd returns a new Passwd given a uid
//or an error if any occurred
func NewPasswd(uid uint) (*Passwd, error) {
	var buf = new(*C.char)
	pwd, err := C.uidtopwd(C.uid_t(uid), buf)

	defer C.free(unsafe.Pointer(*buf))
	defer C.free(unsafe.Pointer(pwd))

	if err != nil {
		return nil, err
	}
	if pwd == nil {
		return nil, fmt.Errorf("uid does not exist")
	}
	return &Passwd{
		Name:    C.GoString(pwd.pw_name),
		Passwd:  C.GoString(pwd.pw_passwd),
		UID:     uint(pwd.pw_uid),
		GID:     uint(pwd.pw_gid),
		GECOS:   C.GoString(pwd.pw_gecos),
		HomeDir: C.GoString(pwd.pw_dir),
		Shell:   C.GoString(pwd.pw_shell),
	}, nil
}
