#ifndef CAPI_H
#define CAPI_H

#ifdef __cplusplus
extern "C" {
#endif

typedef void QClipboard_;

QClipboard_ *getClipboard();
const char *getText(QClipboard_ *);
void setText(QClipboard_ *, const char *);

#ifdef __cplusplus
}
#endif

#endif // CAPI_H
