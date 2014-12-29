
#include "runtime.h"
#include "cgocall.h"

void ·_Cerrno(void*, int32);

void
·_Cfunc_GoString(int8 *p, String s)
{
	s = runtime·gostring((byte*)p);
	FLUSH(&s);
}

void
·_Cfunc_GoStringN(int8 *p, int32 l, String s)
{
	s = runtime·gostringn((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_GoBytes(int8 *p, int32 l, Slice s)
{
	s = runtime·gobytes((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_CString(String s, int8 *p)
{
	p = runtime·cmalloc(s.len+1);
	runtime·memmove((byte*)p, s.str, s.len);
	p[s.len] = 0;
	FLUSH(&p);
}

void
·_Cfunc__CMalloc(uintptr n, int8 *p)
{
	p = runtime·cmalloc(n);
	FLUSH(&p);
}

#pragma cgo_import_static _cgo_4346962e38a9_Cfunc_random
void _cgo_4346962e38a9_Cfunc_random(void*);

void
·_Cfunc_random(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_4346962e38a9_Cfunc_random, &p);
}

#pragma cgo_import_static _cgo_4346962e38a9_Cfunc_srandom
void _cgo_4346962e38a9_Cfunc_srandom(void*);

void
·_Cfunc_srandom(struct{void *y[1];}p)
{
	runtime·cgocall(_cgo_4346962e38a9_Cfunc_srandom, &p);
}

